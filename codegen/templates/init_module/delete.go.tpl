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
	res "{{$project_name}}/utils/response"

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

	// query goes here

	return res.ResponseSuccess(c, data, res.DeletedMessage(module_name))
}
