package service

import (
	"context"
	v1 "novelman/api/v1"
	"novelman/internal/model"
	"novelman/internal/repository"
)

type AppService interface {
	GetApp(ctx context.Context, param *v1.GetAppRequest) (*model.App, error)
	GetApps(ctx context.Context, param *v1.GetAppsRequest) ([]*model.App, error)
	DeleteApp(ctx context.Context, param *v1.DeleteAppRequest) (*model.App, error)
	EditApp(ctx context.Context, param *v1.EditAppRequest) (*model.App, error)
	DisableApp(ctx context.Context, param *v1.DeleteAppRequest) (*model.App, error)
	CreateApp(ctx context.Context, param *v1.CreateAppRequest) error
}

func NewAppService(
	service *Service,
	appRepository repository.AppRepository,
) AppService {
	return &appService{
		Service:       service,
		appRepository: appRepository,
	}
}

type appService struct {
	*Service
	appRepository repository.AppRepository
}

func (s *appService) GetApp(ctx context.Context, param *v1.GetAppRequest) (*model.App, error) {
	return s.appRepository.GetApp(ctx, int64(param.AppId))
}

func (s *appService) CreateApp(ctx context.Context, app *v1.CreateAppRequest) error {
	return s.appRepository.Create(ctx, &model.App{
		AppName: app.AppName,
	})
}

func (s *appService) GetApps(ctx context.Context, param *v1.GetAppsRequest) ([]*model.App, error)
func (s *appService) DeleteApp(ctx context.Context, param *v1.DeleteAppRequest) (*model.App, error)
func (s *appService) EditApp(ctx context.Context, param *v1.EditAppRequest) (*model.App, error)
func (s *appService) DisableApp(ctx context.Context, param *v1.DeleteAppRequest) (*model.App, error)
