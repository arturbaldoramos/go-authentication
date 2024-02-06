package router

import (
	"github.com/arturbaldoramos/go-authentication/pkg/handler"
	"github.com/arturbaldoramos/go-authentication/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {
	router.Get("/user/:uuid", middleware.AuthMiddleware, handler.GetUserById)
	router.Delete("/user/:uuid", middleware.AuthMiddleware, handler.DeleteUserById)
	router.Get("/user", middleware.AuthMiddleware, handler.GetAllUsers)
	router.Post("/user", middleware.AuthMiddleware, handler.CreateUser)

	router.Post("/login", handler.Login)
	router.Post("/logout", handler.Logout)
	router.Static("/css", "/pkg/static/output.css")
}
