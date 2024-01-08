package db

import (
	"fmt"
	"github.com/arturbaldoramos/go-authentication/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {

	//load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	errMigrate := db.AutoMigrate(&models.User{})
	if errMigrate != nil {
		fmt.Println("Error executing migration")
		panic(errMigrate)
	}
	fmt.Println("Successfully  connected...", db)

	return db
}
