package module

import (
	"auth-service/internal/controller"
	"auth-service/internal/interfaces"
	"auth-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthModule struct {
	app            *fiber.App
	userRepo       interfaces.UserRepository
	userPointsRepo interfaces.UserPointsRepository
	jwtModule      interfaces.JWTModule
}

func NewAuthModule(
	app *fiber.App,
	userRepo interfaces.UserRepository,
	userPointsRepo interfaces.UserPointsRepository,
	jwtModule interfaces.JWTModule,
) *AuthModule {
	return &AuthModule{
		app:            app,
		userRepo:       userRepo,
		userPointsRepo: userPointsRepo,
		jwtModule:      jwtModule,
	}
}

func (m *AuthModule) Initialization() error {
	service := service.NewAuthService(m.userPointsRepo, m.userRepo, m.jwtModule)
	controller := controller.NewAuthController(service)

	m.app.Post("/api/auth/register", controller.Register)
	m.app.Post("/api/auth/login", controller.Login)

	return nil
}
