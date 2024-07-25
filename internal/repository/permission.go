package repository

import (
	"context"
	"novelman/internal/model"
)

type PermissionRepository interface {
	GetPermission(ctx context.Context, id int64) (*model.Permission, error)
	Create(ctx context.Context, permission *model.Permission) error
}

func NewPermissionRepository(
	repository *Repository,
) PermissionRepository {
	return &permissionRepository{
		Repository: repository,
	}
}

type permissionRepository struct {
	*Repository
}

func (r *permissionRepository) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	var permission model.Permission

	return &permission, nil
}
func (r *permissionRepository) Create(ctx context.Context, permission *model.Permission) error {
	// var role model.Role

	return nil
}
