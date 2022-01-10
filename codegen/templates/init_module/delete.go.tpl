{{ $project_name := .ProjectInfo.ProjectName  -}}
{{ $module_name := 	convertToTitle .Module.ModuleName -}}
{{ $controller_get_all  	:= printf "%s%s%s" "Get" .Module.ModuleName "s" -}}
{{ $controller_get_single  	:= printf "%s%s" "Get" .Module.ModuleName -}}
{{ $controller_create  	:= printf "%s%s" "Create" .Module.ModuleName 	-}}
{{ $controller_update  	:= printf "%s%s" "Update" .Module.ModuleName 	-}}
{{ $controller_delete  	:= printf "%s%s" "Delete" .Module.ModuleName 	-}}
{{ $lc	:= convertToLowercase .Module.ModuleName  -}}
{{ $include_db_conn := printf "%s" .ProjectInfo.IncludeDbSnippet -}}

package {{$lc}}_module

import (
	"{{$project_name}}/utils/response"

	"github.com/gofiber/fiber/v2"
	"strconv" 
	{{- if ( eq $include_db_conn "true") }}
	"{{$project_name}}/settings/database"
	{{- end }}
)

func {{$controller_delete}}(c *fiber.Ctx) error {

	{{- validate_url_param}}

	{{- if ( eq $include_db_conn "true") }}
	{{ db_conn_snippet }}
	{{- end }}

	return response.ResponseSuccess(c, data, response.DeletedMessage(module_name))
}