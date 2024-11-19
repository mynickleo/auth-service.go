package module

import (
	"auth-service/internal/controller"
	"auth-service/internal/interfaces"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/sqlcqueries"

	"github.com/gofiber/fiber/v2"
)

type PointModule struct {
	app  *fiber.App
	q    *sqlcqueries.Queries
	repo interfaces.UserPointsRepository
}

func NewPointModule(app *fiber.App, q *sqlcqueries.Queries) *PointModule {
	return &PointModule{
		app: app,
		q:   q,
	}
}

func (m *PointModule) Initialization() error {
	m.repo = repository.NewUserPointsRepository(m.q)
	service := service.NewPointService(m.repo)
	controller := controller.NewPointController(service)

	m.app.Put("api/points/add", controller.UpdatePoint)
	m.app.Get("api/points/leaderboard", controller.GetAllPoints)

	return nil
}

func (m *PointModule) GetRepo() interfaces.UserPointsRepository {
	return m.repo
}
