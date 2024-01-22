package router

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func Initialize() {
	router := fiber.New()

	initializeRoutes(router)
	log.Fatal(router.Listen(os.Getenv("API_PORT")))
}
