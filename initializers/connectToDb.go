package initializers

import (
	"log"
	"os"

	"github.com/YasserRABIE/authentication-porject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the DB")
	}

	DB = db

	CreateTables(db)
}

func CreateTables(DB *gorm.DB) {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("failed to create user table")
	}

	if err := DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("failed to create task table")
	}
}
