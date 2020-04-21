package server

import (
	"context"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
	"go-user-auth-api/application"
	"go-user-auth-api/application/handler"
	_ "go-user-auth-api/docs"
	"go-user-auth-api/domain"
	"go-user-auth-api/infrastructure/commandquerybus"
	"go-user-auth-api/infrastructure/configuration"
	"go-user-auth-api/infrastructure/errors"
	"go-user-auth-api/infrastructure/repository"
	"go-user-auth-api/infrastructure/web"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer() {
	//Server
	e := echo.New()

	//Init
	config := configuration.NewConfigurationManager()
	databaseConfig := config.GetDatabaseConfig()
	msSqlConnection := repository.NewMSSqlConnection(databaseConfig)

	//CommandQueryBus
	commandQueryBus := commandquerybus.New()

	//Error
	e.HTTPErrorHandler = errors.CustomEchoHTTPErrorHandler

	//Cors
	e.Use(middleware2.CORS())

	//Log
	e.Logger.SetLevel(log.OFF)
	//TODO will fix later bad code....
	e.Use(middleware2.GzipWithConfig(middleware2.GzipConfig{
		Level: 5,
		Skipper: func(e echo.Context) bool {
			if e.Path() == "/swagger/*" {
				return true
			}
			return false
		},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use()

	//UserController
	userRepository := repository.NewUserRepository(*msSqlConnection)
	userServiceImp := application.NewUserServiceImp(userRepository)
	userController := web.NewUserController(userServiceImp)
	userController.Register(e)

	//PermissionController
	permissionRepository := repository.NewPermissionRepository(*msSqlConnection)
	permissionServiceImp := application.NewPermissionServiceImp(permissionRepository)
	permissionController := web.NewPermissionController(permissionServiceImp)
	permissionController.Register(e)

	//RoleController
	roleRepository := repository.NewRoleRepository(*msSqlConnection)
	roleServiceImp := application.NewRoleServiceImp(roleRepository)

	queryHandler := handler.NewRoleQueryHandler(roleServiceImp)
	commandQueryBus.RegisterQueryHandler(handler.GetRoleByIdQuery{}, queryHandler.GetRoleByIdQueryHandler)
	commandQueryBus.RegisterQueryHandler(handler.GetAllRolesQuery{}, queryHandler.GetAllRolesQueryQueryHandler)

	commandHandler := handler.NewRoleCommandHandler(roleServiceImp)
	commandQueryBus.RegisterCommandHandler(domain.CreateRoleCommand{}, commandHandler.CreateRoleCommandHandler)
	commandQueryBus.RegisterCommandHandler(domain.DeleteRoleCommand{}, commandHandler.DeleteRoleCommandHandler)

	roleController := web.NewRoleController(*commandQueryBus)
	roleController.Register(e)

	//ApplicationController
	applicationRepository := repository.NewApplicationRepository(*msSqlConnection)
	applicationServiceImp := application.NewApplicationServiceImp(applicationRepository)
	applicationController := web.NewApplicationController(applicationServiceImp)
	applicationController.Register(e)

	//UserRoleController
	userRoleRepository := repository.NewUserRoleRepository(*msSqlConnection)
	userRoleServiceImp := application.NewUserRoleServiceImp(userRepository, roleRepository, userRoleRepository)
	userRoleController := web.NewUserRoleController(userRoleServiceImp)
	userRoleController.Register(e)

	//RolePermissionController
	rolePermissionRepository := repository.NewRolePermissionRepository(*msSqlConnection)
	rolePermissionServiceImp := application.NewRolePermissionServiceImp(roleRepository, permissionRepository, rolePermissionRepository, applicationRepository)
	rolePermissionController := web.NewRolePermissionController(rolePermissionServiceImp)
	rolePermissionController.Register(e)

	//UserPermissionController
	userPermissionRepository := repository.NewUserPermissionRepository(*msSqlConnection)
	userPermissionServiceImp := application.NewUserPermissionServiceImp(userPermissionRepository)
	userPermissionController := web.NewUserPermissionController(userPermissionServiceImp)
	userPermissionController.Register(e)

	//HealthCheck
	registerHealthCheck(e)

	//Swagger
	registerSwaggerRedirect(e)

	// Start server
	go func() {
		if err := e.Start(config.GetServerConfig().Port); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func registerHealthCheck(e *echo.Echo) {
	e.GET("/healthcheck", health)
}

func registerSwaggerRedirect(e *echo.Echo) {
	e.GET("/", swaggerRedirect)
}

func swaggerRedirect(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/swagger/index.html")
}

func health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
