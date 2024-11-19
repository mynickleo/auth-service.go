package main

import (
	"auth-service/internal/app"
	"log"
)

func main() {
	err := app.InitializationApp()
	if err != nil {
		log.Fatal(err)
	}
}
