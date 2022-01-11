package router

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Home")
}

func Url(app *fiber.App) {
	// home_view
	app.Get("/", Home)

	// function that will import the generated modules, if there are any
	ProjectModules(app)
}
