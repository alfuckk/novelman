package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func NewCasbin(conf *viper.Viper) *casbin.Enforcer {
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	`)
	if err != nil {
		panic(err)
	}

	driver := conf.GetString("data.db.user.driver")
	dsn := conf.GetString("data.db.user.dsn")
	a, _ := gormadapter.NewAdapter(driver, dsn)
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}
	return e
}
