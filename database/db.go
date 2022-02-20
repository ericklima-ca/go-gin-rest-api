package database

import (
	"log"
	"os"

	"github.com/ericklima-ca/go-gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

const DATABASE_DEFAULT = "postgresql://root:root@localhost:5432/root?sslmode=disable"

// "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

func Connect() {
	var connString string = os.Getenv("DATABASE_URL")

	if connString == "" {
		connString = DATABASE_DEFAULT
	}

	DB, err = gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Panic(err.Error())
	}
	DB.AutoMigrate(
		&models.Order{},
		&models.Document{},
	)
}
