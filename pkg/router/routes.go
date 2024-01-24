package router

import (
	"fmt"
	"github.com/arturbaldoramos/go-authentication/pkg/models"
	u "github.com/arturbaldoramos/go-authentication/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {
	router.Get("/user", func(ctx *fiber.Ctx) error {
		uuid := ctx.Query("uuid")
		fmt.Printf("")
		user := models.GetUser(uuid)

		if user != nil {
			resp := u.Message(true, "success")

			userWithoutPassword := map[string]interface{}{
				"id":        user.ID,
				"name":      user.Name,
				"email":     user.Email,
				"CreatedAt": user.CreatedAt,
				"UpdatedAt": user.UpdatedAt,
				"DeletedAt": user.DeletedAt,
			}

			resp["user"] = userWithoutPassword
			return ctx.JSON(resp)
		}
		return ctx.JSON(u.Message(false, "user not found"))
	})
}
