package repository

import (
	"context"
	"novelman/internal/model"
)

type AppRepository interface {
	GetApp(ctx context.Context, id int64) (*model.App, error)
	GetApps(ctx context.Context, id int64) ([]*model.App, error)
	Create(ctx context.Context, app *model.App) error
	Update(ctx context.Context, app *model.App) error
	Delete(ctx context.Context, id int64) error
}

func NewAppRepository(
	repository *Repository,
) AppRepository {
	return &appRepository{
		Repository: repository,
	}
}

type appRepository struct {
	*Repository
}

func (r *appRepository) GetApp(ctx context.Context, id int64) (*model.App, error) {
	var app model.App

	return &app, nil
}

func (r *appRepository) Create(ctx context.Context, app *model.App) error {

	return nil
}

func (r *appRepository) GetApps(ctx context.Context, id int64) ([]*model.App, error)
func (r *appRepository) Update(ctx context.Context, app *model.App) error
func (r *appRepository) Delete(ctx context.Context, id int64) error
