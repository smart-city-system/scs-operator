package db

import (
	"fmt"
	config "scs-operator/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(cfg *config.Config) (*gorm.DB, error) {
	// Configure GORM with better logging and connection settings
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn(cfg)), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("open db connection: %w", err)
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying sql.DB: %w", err)
	}

	// Configure connection pool settings
	sqlDB.SetMaxOpenConns(25)                 // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)                  // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

func dsn(c *config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable connect_timeout=10",
		c.Database.DbHost, c.Database.DbPort, c.Database.DbUser, c.Database.DbName, c.Database.DbPassword)
}
