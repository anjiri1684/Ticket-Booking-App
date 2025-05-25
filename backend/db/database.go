package db

import (
	"fmt"

	"github.com/anjiri1684/ticket-booking-project-v1/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s port=5432",
		config.DBHost,
		config.DBUser,
		config.DBName,
		config.DBPassword,
		config.DBSSLmode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	// Bust Neon plan cache
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Unable to get raw database object: %v", err)
	}
	sqlDB.SetMaxIdleConns(0) // 🔥 This avoids using cached statements in Neon

	log.Info("Connected to the database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to run table migrations: %v", err)
	}

	return db
}
