package router

import (
	res "my_server/utils/response"

	"github.com/gin-gonic/gin"
)

func Endpoints(app *gin.Engine) {

	app.GET("/", func(c *gin.Context) {
		res.Response(c, 200, "", "Hello There!")
	})

	ProjectModules(app)
}
