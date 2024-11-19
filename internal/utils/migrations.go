package utils

import (
	"auth-service/config"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func RunMigrations() error {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DB,
	)

	cmd := exec.Command("migrate", "-path", "./internal/database/postgres/migrations", "-database", databaseURL, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Migration command failed: %v", err)
		return errors.New("failed to apply migrations")
	}

	log.Println("Migrations applied successfully")
	return nil
}
