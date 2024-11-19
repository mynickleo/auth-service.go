package module

import (
	"auth-service/internal/controller"
	"auth-service/internal/interfaces"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/sqlcqueries"

	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	q    *sqlcqueries.Queries
	app  *fiber.App
	repo interfaces.UserRepository
}

func NewUserModule(q *sqlcqueries.Queries, app *fiber.App) *UserModule {
	return &UserModule{
		q:   q,
		app: app,
	}
}

func (m *UserModule) Initialization() error {
	m.repo = repository.NewUserRepository(m.q)
	service := service.NewUserService(m.repo)
	controller := controller.NewUserController(service)

	m.app.Post("/api/users", controller.CreateUser)
	m.app.Get("/api/users", controller.GetUsers)
	m.app.Get("/api/users/stats", controller.GetCurrentUser)
	m.app.Put("/api/users/:id", controller.UpdateUser)
	m.app.Delete("/api/users/:id", controller.DeleteUser)

	return nil
}

func (m *UserModule) GetRepo() interfaces.UserRepository {
	return m.repo
}
