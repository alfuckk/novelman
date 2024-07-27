package repository

import (
	"context"
	"fmt"
	"novelman/internal/model"
	"novelman/pkg/log"
	"novelman/pkg/sid"
	"novelman/pkg/zapgorm2"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ctxTxKey = "TxKey"

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(
	logger *log.Logger,
	db *gorm.DB,
	rdb *redis.Client,
) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

// DB return tx
// If you need to create a Transaction, you must call DB(ctx) and Transaction(ctx,fn)
func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func NewDB(conf *viper.Viper, l *log.Logger) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	logger := zapgorm2.New(l.Logger)
	driver := conf.GetString("data.db.user.driver")
	dsn := conf.GetString("data.db.user.dsn")

	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html
	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger,
		})
	case "postgres":
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		panic("unknown db driver")
	}
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&model.Admin{}, &model.Permission{}, &model.Role{}, &model.RolePermission{}, &model.AdminRole{}, &model.App{}, &model.User{}, &model.UserApp{})
	if err := InitAdmin(l, db); err != nil {
		l.Info("InitAdmin", zap.String("", ""))
	}

	return db
}

func InitAdmin(l *log.Logger, db *gorm.DB) error {
	// Check if roles exist
	var totalRoles, totalAdmins int64

	if err := db.Model(&model.Role{}).Count(&totalRoles).Error; err != nil {
		return err
	}

	if totalRoles == 0 {
		// Create initial role
		roles := []*model.Role{
			{RoleName: "super admin"},
		}
		for _, role := range roles {
			if err := db.Create(role).Error; err != nil {
				return err
			}
		}

		l.Info("Created initial role", zap.String("role", "super admin"))

		// Check if admin user exists
		if err := db.Model(&model.Admin{}).Where("email = ?", "admin@admin.com").Count(&totalAdmins).Error; err != nil {
			return err
		}

		if totalAdmins == 0 {
			password, err := sid.NewSid().GenString()
			if err != nil {
				return err
			}

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

			if err != nil {
				return err
			}
			// Create admin user
			admin := model.Admin{
				Email:    "admin@admin.com",
				Password: string(hashedPassword),
				Roles:    roles,
			}
			if err := db.Preload("Role").Create(&admin).Error; err != nil {
				return err
			}

			l.Debug("Created admin user", zap.String("email", admin.Email), zap.String("password", password))
		}
	}

	return nil
}

func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}
