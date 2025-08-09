package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	config "scs-operator/config"
	"scs-operator/internal/models"
	"scs-operator/internal/server"
	"scs-operator/pkg/db"
	kafka_client "scs-operator/pkg/kafka"
	"scs-operator/pkg/logger"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
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
		&models.Premise{},
		&models.Alarm{},
		&models.Incident{},
		&models.IncidentGuidance{},
		&models.IncidentGuidanceStep{},
		&models.GuidanceTemplate{},
		&models.GuidanceStep{},
	)
	if err != nil {
		appLogger.Fatalf("Database migration failed: %s", err)
	}
	// Initialize the server
	s := server.NewServer(&cfg, psqlDb, appLogger)

	// Create a WaitGroup to manage goroutines
	var wg sync.WaitGroup

	// Create a parent context for the Kafka consumer
	consumerCtx, consumerCancel := context.WithCancel(context.Background())

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		if err := s.Run(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatalf("Error starting server: %v", err)
		}
	}()

	// Start Kafka consumer in a separate goroutine
	wg.Add(1) // Increment the WaitGroup counter
	go startKafkaConsumer("alarm.triggered", &cfg, appLogger, consumerCtx, &wg)
	wg.Add(1) // Increment the WaitGroup counter
	go startKafkaConsumer("incident.detected", &cfg, appLogger, consumerCtx, &wg)

	// Block until a signal is received
	<-quit

	appLogger.Info("Shutting down the server and consumer...")

	// Signal consumer to stop processing
	consumerCancel()
	// Wait a moment to allow goroutine to notice context cancellation
	time.Sleep(1 * time.Second) //

	// Create a separate, timeout context for the server shutdown
	serverShutdownCtx, serverShutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer serverShutdownCancel()

	// Shut down the Echo server
	if err := s.Shutdown(serverShutdownCtx); err != nil {
		appLogger.Errorf("Server shutdown failed: %v", err)
	}

	wg.Wait()

	appLogger.Info("Server and consumer stopped.")
}

func startKafkaConsumer(topic string, cfg *config.Config, logger *logger.ApiLogger, ctx context.Context, wg *sync.WaitGroup) {
	// Ensure wg.Done() is called when the function exits
	defer wg.Done()
	// Initialize Kafka consumer
	kafkaCfg := kafka_client.Config{
		Brokers: strings.Split(cfg.Kafka.Brokers, ","),
		Topic:   topic,
	}
	consumerCfg := kafka_client.ConsumerConfig{
		GroupID:        "scs-operator",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: 1000,
		StartOffset:    kafka.FirstOffset,
	}
	consumer := kafka_client.NewConsumer(&kafkaCfg, &consumerCfg)
	defer func() {
		logger.Info("Closing Kafka consumer...")
		if err := consumer.Close(); err != nil {
			logger.Errorf("Failed to close consumer: %v", err)
		}
		logger.Info("Kafka consumer closed.")
	}()
	logger.Info("Kafka consumer initialized")

	// Continuously read messages
	for {
		select {
		case <-ctx.Done():
			logger.Info("Context canceled. Stopping consumer.")
			return
		default:
			msg, err := consumer.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					logger.Info("Consumer stopped due to context cancellation.")
					return
				}
				logger.Errorf("Failed to read message: %v", err)
				continue
			}
			logger.Infof("Received message at offset %d: %s = %s", msg.Offset, string(msg.Key), string(msg.Value))
		}
	}
}
