{{ $project_name := .ProjectInfo.ProjectName -}}
{{ $modules := 	.Modules  -}}
// this file will be updated on each run. Do not edit

package router

import (
	{{if .Modules -}}
	{{range $x := .Modules}}"{{$project_name}}/modules/{{convertToLowercase $x}}_module"
	{{end}}{{end}}
	"github.com/gofiber/fiber/v2"
)

func ProjectModules(app *fiber.App) {

	{{if .Modules -}}

	// add a prefix for the routes
	api := app.Group("/api")

	// pass down the app + the api prefix
	{{range $x := .Modules}}{{convertToLowercase $x}}_module.Routes(app, api)
	{{end}}{{end}}
}