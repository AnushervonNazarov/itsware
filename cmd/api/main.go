package main

import (
	"itsware/configs"
	"itsware/db"
	"itsware/logger"
	"itsware/router"
	"itsware/server"
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

	if err := logger.Init(); err != nil {
		log.Fatalf("Logger initialization error: %s", err)
	}

	db.InitDB()

	mainServer := new(server.Server)
	if err := mainServer.Run(configs.AppSettings.AppParams.PortRun, router.RunRoutes()); err != nil {
		log.Fatalf("Error starting HTTP server: %s", err)
	}
}
