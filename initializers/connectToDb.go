package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() {
	_, err := gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the DB")
	}
}
