package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/arturbaldoramos/go-authentication/pkg/models"
	u "github.com/arturbaldoramos/go-authentication/pkg/utils"
)

func GetUserById(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	user := models.GetUserByID(uuid)

	if user != nil {
		resp := u.Message(true, "success")

		userWithoutPassword := map[string]interface{}{
			"id":        user.ID,
			"name":      user.Name,
			"email":     user.Email,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
		}

		resp["user"] = userWithoutPassword
		return ctx.JSON(resp)
	}
	return ctx.JSON(u.Message(false, "user not found"))
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	createdUser := user.Create()
	return ctx.JSON(createdUser)

}
