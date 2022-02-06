// this file will be updated on each run. Do not edit

package router

import (
	"server_with_modules_and_sqlc/modules/country_module"
	"server_with_modules_and_sqlc/modules/user_module"

	"github.com/gofiber/fiber/v2"
)

func ProjectModules(app *fiber.App) {

	// add a prefix for the routes
	api := app.Group("/api")

	// pass down the app + the api prefix
	user_module.Routes(app, api)
	country_module.Routes(app, api)

}
