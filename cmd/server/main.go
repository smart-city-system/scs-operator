package main

import (
	"log"
	"smart-city/config"
	"smart-city/internal/models"
	"smart-city/internal/server"
	"smart-city/pkg/db"
	"smart-city/pkg/logger"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func main() {
	// Load configuration from config file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	//Init logger
	appLogger := logger.NewApiLogger(&cfg)
	appLogger.InitLogger()
	appLogger.Infof("LogLevel: %s, Mode: %s", cfg.Logger.Level, cfg.Server.Mode)

	//Init db
	psqlDb, err := db.NewGormDB(&cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Info("Postgres connected")
	}

	// Auto-migrate models
	err = psqlDb.AutoMigrate(
		&models.User{},
		&models.Premise{},
		&models.Camera{},
		&models.Alert{},
		&models.Incident{},
		&models.DispatchRequest{},
	)
	if err != nil {
		appLogger.Fatalf("Database migration failed: %s", err)
	}
	// Initialize the server
	s := server.NewServer(&cfg, psqlDb, appLogger)

	// Start the server
	if err := s.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
