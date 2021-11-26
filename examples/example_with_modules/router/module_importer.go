// this file will be updated on each run. Do not edit

package router

import (
	"example_with_modules/modules/user_module"
	"example_with_modules/modules/task_module"
	
	"github.com/gofiber/fiber/v2"
)

func ProjectModules(app *fiber.App) {

	// add a prefix for the routes
	api := app.Group("/api")

	// pass down the app + the api prefix
	user_module.Routes(app, api)
	task_module.Routes(app, api)
	
}