package database

import (
	"assignment-2/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file: ", er)
	}
	dbUser, dbPassword, dbPort, dbName, dbHost :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST")
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbPort, dbName)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	db.Debug().AutoMigrate(models.Orders{}, models.Items{})
}

func GetDB() *gorm.DB {
	return db
}
