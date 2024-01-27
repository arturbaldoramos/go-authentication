package router

import (
	"github.com/arturbaldoramos/go-authentication/pkg/handler"
	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {
	router.Get("/user/:uuid", handler.GetUserById)
	router.Delete("/user/:uuid", handler.DeleteUserById)
	router.Get("/user", handler.GetAllUsers)
	router.Post("/user", handler.CreateUser)
}
