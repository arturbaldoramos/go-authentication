package main

import (
	"fmt"
	"github.com/arturbaldoramos/go-authentication/pkg/models"
	"github.com/arturbaldoramos/go-authentication/pkg/routes"
	"github.com/joho/godotenv"
	"log"
	"os"

	database "github.com/arturbaldoramos/go-authentication/pkg/db"
)

func main() {

	//load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//setup db
	database.ConnectDB()
	fmt.Printf("[+] Success connecting to Database on port:%s\n", os.Getenv("DB_PORT"))

	//run db migration
	errMigrate := database.DB.AutoMigrate(&models.User{})
	if errMigrate != nil {
		fmt.Println("Error executing migration")
		panic(errMigrate)
	}
	fmt.Println("[+] Success running migrations")

	//initialize router
	fmt.Printf("[+] Success running API on port%s\n", os.Getenv("API_PORT"))
	router.Initialize()
}
