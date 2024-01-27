package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"

	"github.com/arturbaldoramos/go-authentication/pkg/models"
	u "github.com/arturbaldoramos/go-authentication/pkg/utils"
)

func GetUserById(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	user := models.GetUserByID(uuid)

	if user != nil {
		resp := u.Message(true, "Success")

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
	return ctx.JSON(u.Message(false, "User not found"))
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	return ctx.JSON(user.Create())
}

func GetAllUsers(ctx *fiber.Ctx) error {
	users := models.GetAllUsers()

	return ctx.JSON(users)
}

func DeleteUserById(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")

	return ctx.JSON(models.DeleteUserById(uuid))
}

func Login(ctx *fiber.Ctx) error {
	login := new(models.UserLogin)

	if err := ctx.BodyParser(login); err != nil {
		return err
	}

	resp, token := models.Login(login)

	if token != "" {
		ctx.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: token,
		})
	}

	return ctx.JSON(resp)
}

func Logout(ctx *fiber.Ctx) error {

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	return ctx.JSON(u.Message(true, "Logout successfully"))
}
