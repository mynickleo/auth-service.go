package app

import (
	"auth-service/config"
	"auth-service/internal/database/postgres"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializationApp() error {
	err := config.InitConfig()
	if err != nil {
		return err
	}

	database := postgres.NewDataBase()
	err = database.InitializationDB()
	if err != nil {
		return err
	}

	app := fiber.New()
	app.Use(logger.New())

	diContainer := NewDIContainer(database.GetQueries(), app)
	diContainer.InitializationModules()

	log.Fatal(app.Listen(":" + config.AppConfig.Port))

	return nil
}
