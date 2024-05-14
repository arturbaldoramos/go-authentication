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

func MarketplacePage(ctx *fiber.Ctx) error {
	marketplacePageComponent := template.MarketplacePage()
	return utils.Render(ctx, marketplacePageComponent)
}

func FeaturesPage(ctx *fiber.Ctx) error {
	featuresPageComponent := template.FeaturesPage()
	return utils.Render(ctx, featuresPageComponent)
}

func TeamPage(ctx *fiber.Ctx) error {
	teamPageComponent := template.TeamPage()
	return utils.Render(ctx, teamPageComponent)
}

func ContactPage(ctx *fiber.Ctx) error {
	contactPageComponent := template.ContactPage()
	return utils.Render(ctx, contactPageComponent)
}

func NotFoundPage(ctx *fiber.Ctx) error {
	notFoundComponent := template.NotFoundPage(
		"Something's missing.",
		"Sorry, we can't find that page. You'll find lots to explore on the home page.",
		"404")
	return utils.Render(ctx, notFoundComponent)
}
