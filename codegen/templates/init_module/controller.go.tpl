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
	"{{$project_name}}/utils/response"
	"github.com/gofiber/fiber/v2"
)

var data = ""

func {{$controller_get_all}}(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "{{$module_name}}s found!")
}

func {{$controller_get_single}}(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "{{$module_name}} found!")
}

func {{$controller_create}}(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "{{$module_name}} created succesfully!")
}

func {{$controller_update}}(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "{{$module_name}} updated succesfully!")
}

func {{$controller_delete}}(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "{{$module_name}} deleted succesfully!")
}
