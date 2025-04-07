// db/database.go
package db

import (
	"fmt"
	"log"

	"github.com/Omotolani98/k8ly/services/reqly/config"
	"github.com/Omotolani98/k8ly/services/reqly/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.LoggedRequest{})
}
