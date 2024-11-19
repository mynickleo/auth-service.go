package app

import (
	"auth-service/internal/interfaces"
	"auth-service/internal/module"
	"auth-service/pkg/sqlcqueries"

	"github.com/gofiber/fiber/v2"
)

type DIContainer struct {
	queries     *sqlcqueries.Queries
	app         *fiber.App
	jwtModule   interfaces.JWTModule
	userModule  interfaces.UserModule
	pointModule interfaces.PointModule
	authModule  interfaces.AuthModule
	readyModule interfaces.ReadyModule
}

func NewDIContainer(q *sqlcqueries.Queries, app *fiber.App) *DIContainer {
	return &DIContainer{
		queries:   q,
		app:       app,
		jwtModule: module.NewJWTModule(),
	}
}

func (di *DIContainer) InitializationModules() error {
	err := di.InitializationReadyModule()
	if err != nil {
		return err
	}

	err = di.InitializationUserModule()
	if err != nil {
		return err
	}

	err = di.InitializationPointModule()
	if err != nil {
		return err
	}

	err = di.InitializationAuthModule()
	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationReadyModule() error {
	di.readyModule = module.NewReadyModule(di.app)
	err := di.readyModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationUserModule() error {
	di.app.Group("/api/users", di.jwtModule.JWTGuard(), di.jwtModule.CheckUserGuard)
	di.userModule = module.NewUserModule(di.queries, di.app)
	err := di.userModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationPointModule() error {
	di.app.Group("/api/points", di.jwtModule.JWTGuard(), di.jwtModule.CheckUserGuard)
	di.pointModule = module.NewPointModule(di.app, di.queries)
	err := di.pointModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}

func (di *DIContainer) InitializationAuthModule() error {
	di.authModule = module.NewAuthModule(di.app, di.userModule.GetRepo(), di.pointModule.GetRepo(), di.jwtModule)
	err := di.authModule.Initialization()

	if err != nil {
		return err
	}

	return nil
}
