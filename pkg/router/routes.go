package router

import (
	"github.com/arturbaldoramos/go-authentication/pkg/handler"
	"github.com/arturbaldoramos/go-authentication/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {
	router.Get("/", handler.MainPage)
	router.Get("/user", middleware.AuthMiddleware, handler.GetAllUsers)
	router.Get("/user/:uuid", middleware.AuthMiddleware, handler.GetUserById)
	router.Delete("/user/:uuid", middleware.AuthMiddleware, handler.DeleteUserById)
	router.Post("/user", handler.CreateUser)

	router.Get("/login", handler.LoginPage)
	router.Get("/register", handler.RegisterPage)

	router.Post("/login", handler.Login)
	router.Post("/logout", handler.Logout)
	router.Static("/css", "/pkg/static/output.css")
}
