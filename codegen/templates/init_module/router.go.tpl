{{ $project_name := .ProjectInfo.ProjectName  -}}
{{ $module_name := 	convertToTitle .Module.ModuleName -}}
{{ $controller_get_all  	:= printf "%s%s%s%s" "Get" "All" .Module.ModuleName "s" -}}
{{ $controller_get_single  	:= printf "%s%s" "Get" .Module.ModuleName -}}
{{ $controller_create  	:= printf "%s%s" "Create" .Module.ModuleName 	-}}
{{ $controller_update  	:= printf "%s%s" "Update" .Module.ModuleName 	-}}
{{ $controller_delete  	:= printf "%s%s" "Delete" .Module.ModuleName 	-}}
{{ $lc	:= convertToLowercase .Module.ModuleName  -}}

package {{$lc}}_module

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/{{$lc}}", {{$controller_get_all}}) 

	// Get By Id
	api.Get("/{{$lc}}/:id", {{$controller_get_single}}) 

	// Create
	api.Post("/{{$lc}}", {{$controller_create}}) 

	// Update With Id
	api.Put("/{{$lc}}/:id", {{$controller_update}})
	
	// Delete With Id
	api.Delete("/{{$lc}}/:id", {{$controller_delete}}) 
}