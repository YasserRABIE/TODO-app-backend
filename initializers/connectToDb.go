package initializers

import (
	"log"
	"os"

	"github.com/YasserRABIE/authentication-porject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() {
	DB, err := gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the DB")
	}

	CreateTables(DB)
}

func CreateTables(DB *gorm.DB) {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("failed to create user table")
	}

	if err := DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("failed to create task table")
	}
}
