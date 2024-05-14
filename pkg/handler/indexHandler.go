package handler

import (
	"github.com/arturbaldoramos/go-authentication/pkg/template"
	"github.com/arturbaldoramos/go-authentication/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func MainPage(ctx *fiber.Ctx) error {
	mainPageComponent := template.MainPage()
	return utils.Render(ctx, mainPageComponent)
}

func DashboardPage(ctx *fiber.Ctx) error {
	dashboardPageComponent := template.DashboardPage()
	return utils.Render(ctx, dashboardPageComponent)
}

func NotFoundPage(ctx *fiber.Ctx) error {
	notFoundComponent := template.NotFoundPage(
		"Something's missing.",
		"Sorry, we can't find that page. You'll find lots to explore on the home page.",
		"404")
	return utils.Render(ctx, notFoundComponent)
}
