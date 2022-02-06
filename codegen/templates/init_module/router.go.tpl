{{ $project_name 			:= .ProjectInfo.ProjectName  -}}
{{/* controller + sql variables */ -}}
{{ $module_name                  := .Module.ModuleName -}}
{{ $module_name_lwc			:= ConvertToLowercase .Module.ModuleName  -}}
{{ $module_singular_title   := ConvertToTitle $module_name -}}
{{ $module_singular_lwc     := ConvertToLowercase $module_name -}}
{{ $module_plural           := ConvertToPlural $module_name -}}
{{ $module_plural_title     := ConvertToTitle $module_name -}}
{{ $module_plural_lwc	    := ConvertToLowercase $module_plural -}}
{{ $controller_get_all  	:= printf "%s%s" "Get" $module_plural -}}
{{ $controller_get_single  	:= printf "%s%s" "Get" $module_name -}}
{{ $controller_create  		:= printf "%s%s" "Create" $module_name -}}
{{ $controller_update  		:= printf "%s%s" "Update" $module_name -}}
{{ $controller_delete  		:= printf "%s%s" "Delete" $module_name -}}
{{ $table                   := $module_plural_lwc -}}
{{/* snippets  */ -}}
{{ $include_db_conn 		:= printf "%s" .ProjectInfo.IncludeDbSnippet -}}

package {{$module_name_lwc}}_module

import (
	"fmt"
	"{{$project_name}}/settings"
	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "{{$module_name}}"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/{{$module_name_lwc}}")

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/{{$module_name_lwc}}", {{$controller_get_all}}) 
	// Get By Id
	api.Get("/{{$module_name_lwc}}/:id", {{$controller_get_single}}) 
	// Create
	api.Post("/{{$module_name_lwc}}", {{$controller_create}}) 
	// Update With Id
	api.Put("/{{$module_name_lwc}}/:id", {{$controller_update}})
	// Delete With Id
	api.Delete("/{{$module_name_lwc}}/:id", {{$controller_delete}}) 
}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=1,max=20"`
	ExampleTitle string `json:"example_title" validate:"required,min=5,max=50"`
}
