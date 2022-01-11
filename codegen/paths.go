package codegen

var init_project_template_path = "templates/init_server/"

var init_project_dirs = [...]string{
	"/",
	"/settings/database/",
	"/utils/response/",
	"/utils/validate/",
	"/router/",
}

var init_project_templates = [...]string{
	"main.go.tpl",
	".env.tpl",
	"README.md.tpl",
	"go.mod.tpl",
	"Dockerfile.tpl",
	"settings/settings.go.tpl",
	"settings/database/database.go.tpl",
	"utils/response/response.go.tpl",
	"utils/response/messages.go.tpl",
	"utils/validate/validate.go.tpl",
	"router/router.go.tpl",
}

var module_files = [...]string{
	"templates/init_module/router.go.tpl",
	"templates/init_module/get.go.tpl",
	"templates/init_module/delete.go.tpl",
	"templates/init_module/post.go.tpl",
	"templates/init_module/put.go.tpl",
}
