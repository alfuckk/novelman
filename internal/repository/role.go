package repository

import (
	"context"
	"novelman/internal/model"
)

type RoleRepository interface {
	GetRole(ctx context.Context, id int64) (*model.Role, error)
	Create(ctx context.Context, role *model.Role) error
}

func NewRoleRepository(
	repository *Repository,
) RoleRepository {
	return &roleRepository{
		Repository: repository,
	}
}

type roleRepository struct {
	*Repository
}

func (r *roleRepository) GetRole(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role

	return &role, nil
}

func (r *roleRepository) Create(ctx context.Context, role *model.Role) error {
	// var role model.Role

	return nil
}
