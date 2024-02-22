package handler

import (
	"github.com/arturbaldoramos/go-authentication/pkg/models"
	"github.com/arturbaldoramos/go-authentication/pkg/template"
	utils "github.com/arturbaldoramos/go-authentication/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUserById(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	user := models.GetUserByID(uuid)

	if user != nil {
		resp := utils.Message(true, "Success")

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
	return ctx.JSON(utils.Message(false, "User not found"))
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	resp := user.Create()
	message := resp["message"].(string)
	if message == "Success" {
		return Login(ctx)
	}
	notificationComponent := template.LoginNotification(message)
	return utils.Render(ctx, notificationComponent)
}

func GetAllUsers(ctx *fiber.Ctx) error {
	users := models.GetAllUsers()

	return ctx.JSON(users)
}

func DeleteUserById(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")

	return ctx.JSON(models.DeleteUserById(uuid))
}
