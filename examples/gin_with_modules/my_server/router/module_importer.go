// this file will be updated on each run. Do not edit

package router

import (
	"my_server/modules/country_module"
	"my_server/modules/user_module"

	"github.com/gin-gonic/gin"
)

func ProjectModules(app *gin.Engine) {

	api := app.Group("/api") // prefix for routes
	{

		user_module.Routes(api)
		country_module.Routes(api)

	}

}
