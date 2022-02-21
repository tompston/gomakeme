// this file will be updated on each run. Do not edit

package router

import (
	"my_server/modules/country_module"
	"my_server/modules/user_module"

	"github.com/gofiber/fiber/v2"
)

func ProjectModules(app *fiber.App) {

	// add a prefix for the routes
	api := app.Group("/api")

	// pass down the app + the api prefix
	user_module.Routes(app, api)
	country_module.Routes(app, api)

}
