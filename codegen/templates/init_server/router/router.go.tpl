{{ $project_name := .ProjectInfo.ProjectName -}}

package router





{{- if ( eq .ProjectInfo.Framework "fiber") }}
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
{{- end }}






{{- if ( eq .ProjectInfo.Framework "gin") }}

import (
	res "{{$project_name}}/utils/response"

	"github.com/gin-gonic/gin"
)

func Endpoints(app *gin.Engine) {

	app.GET("/", func(c *gin.Context) {
		res.Response(c, 200, "", "Hello There!")
	})

	ProjectModules(app)
}
{{- end }}
