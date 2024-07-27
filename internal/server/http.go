package server

import (
	apiV1 "novelman/api/v1"
	"novelman/docs"
	"novelman/internal/handler"
	"novelman/internal/middleware"
	"novelman/pkg/jwt"
	"novelman/pkg/log"
	"novelman/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	appHandler *handler.AppHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			adminAuthRouter := noAuthRouter.Group("/admin")
			{
				adminAuthRouter.POST("/login", adminHandler.Login)
			}
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		adminRouter := v1.Group("/admin")
		{
			appRouter := adminRouter.Group("/app")
			{
				appRouter.POST("/create", appHandler.CreateApp)
				appRouter.POST("/delete", appHandler.CreateApp)
				appRouter.POST("/edit", appHandler.CreateApp)
				appRouter.POST("/get", appHandler.CreateApp)
				appRouter.POST("/list", appHandler.CreateApp)
			}
			userRouter := adminRouter.Group("/user")
			{
				userRouter.POST("/create", userHandler.GetProfile)
				userRouter.POST("/delete", userHandler.GetProfile)
				userRouter.POST("/ban", userHandler.GetProfile)
				userRouter.POST("/edit", userHandler.GetProfile)
				userRouter.POST("/get", userHandler.GetProfile)
				userRouter.POST("/list", userHandler.GetProfile)
			}
			// adminRouter.POST("/")
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
