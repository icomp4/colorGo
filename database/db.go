package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zipqt/goScrape/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Start() error {
	// Loading env vars
	err := godotenv.Load()
	if err != nil {
		return err
	}
	// Secret DB url
	dbURL := os.Getenv("DB_URL")
	dsn := dbURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// Updating global DB var, similar to static in java
	DB = db

	migrate := DB.AutoMigrate(&models.Page{})
	if migrate != nil {
		log.Println(migrate)
	}
	return nil
}
