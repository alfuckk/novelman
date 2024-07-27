//go:build wireinject
// +build wireinject

package wire

import (
	"novelman/internal/handler"
	"novelman/internal/repository"
	"novelman/internal/server"
	"novelman/internal/service"
	"novelman/pkg/app"
	"novelman/pkg/jwt"
	"novelman/pkg/log"
	"novelman/pkg/server/http"
	"novelman/pkg/sid"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewAdminRepository,
	repository.NewAppRepository,
	repository.NewPermissionRepository,
	repository.NewRoleRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewAdminService,
	service.NewAppService,
	service.NewRoleService,
	service.NewPermissionService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewAdminHandler,
	handler.NewAppHandler,
	handler.NewPermissionHandler,
	handler.NewRoleHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("admin"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
