package repository

import (
    "context"
	"novelman/internal/model"
)

type AdminRepository interface {
	GetAdmin(ctx context.Context, id int64) (*model.Admin, error)
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
