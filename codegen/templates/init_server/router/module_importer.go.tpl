{{ $project_name := .ProjectInfo.ProjectName -}}
{{ $modules := 	.Modules  -}}
// this file will be updated on each run. Do not edit

package router

{{- if ( eq .ProjectInfo.Framework "fiber") }}
import (
	{{if .Modules -}}
	{{range $x := .Modules}}"{{$project_name}}/modules/{{ConvertToLowercase $x}}_module"
	{{end}}{{end}}
	"github.com/gofiber/fiber/v2"
)

func ProjectModules(app *fiber.App) {

	{{if .Modules -}}

	// add a prefix for the routes
	api := app.Group("/api")

	// pass down the app + the api prefix
	{{range $x := .Modules}}{{ConvertToLowercase $x}}_module.Routes(app, api)
	{{end}}{{end}}
}
{{- end }}

{{- if ( eq .ProjectInfo.Framework "gin") }}
import (
	{{if .Modules -}}
	{{range $x := .Modules}}"{{$project_name}}/modules/{{ConvertToLowercase $x}}_module"
	{{end}}{{end}}
	"github.com/gin-gonic/gin"
)

func ProjectModules(app *gin.Engine) {

	{{if .Modules -}}

	api := app.Group("/api") // prefix for routes
	{
		
		{{range $x := .Modules}}{{ConvertToLowercase $x}}_module.Routes(api)
		{{end}}
		
	}
	{{end}}
}
{{- end }}