package main

import (
	"itsware/configs"
	"itsware/db"
	"itsware/internal/controllers"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Error reading settings: %s", err)
	}

	db.InitDB()

	controllers.RunRoutes()

}
