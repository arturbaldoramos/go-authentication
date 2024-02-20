package handler

import (
	"github.com/arturbaldoramos/go-authentication/pkg/models"
	"github.com/arturbaldoramos/go-authentication/pkg/template"
	"github.com/arturbaldoramos/go-authentication/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

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

	//TODO melhorar o jeito que checa se o usuário foi logado, não simplesmente lendo o message
	message := resp["message"].(string)
	if message == "Success login in" {
		ctx.Set("X-Login-Success", "true")
	}
	notificationComponent := template.LoginNotification(message)
	return utils.Render(ctx, notificationComponent)

}

func Logout(ctx *fiber.Ctx) error {

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	return ctx.JSON(utils.Message(true, "Logout successfully"))
}

func LoginPage(ctx *fiber.Ctx) error {
	loginPageComponent := template.LoginPage()
	return utils.Render(ctx, loginPageComponent)
}

func RegisterPage(ctx *fiber.Ctx) error {
	registerPageComponent := template.RegisterPage()
	return utils.Render(ctx, registerPageComponent)
}
