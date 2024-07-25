package repository

import (
	"context"
	"errors"
	"novelman/internal/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAdmin(ctx context.Context, id int64) (*model.Admin, error)
	Create(ctx context.Context, admin *model.Admin) error
	GetByEmail(ctx context.Context, email string) (*model.Admin, error)
}

func NewAdminRepository(
	repository *Repository,
) AdminRepository {
	return &adminRepository{
		Repository: repository,
	}
}

type adminRepository struct {
	*Repository
}

func (r *adminRepository) GetAdmin(ctx context.Context, id int64) (*model.Admin, error) {
	var admin model.Admin

	return &admin, nil
}

func (r *adminRepository) Create(ctx context.Context, admin *model.Admin) error {
	if err := r.DB(ctx).Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) GetByEmail(ctx context.Context, email string) (*model.Admin, error) {
	var admin model.Admin
	if err := r.DB(ctx).Where("email = ?", email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}
