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
	".gitignore.tpl",
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
	"templates/init_module/router.go.tmpl",
	"templates/init_module/get.go.tmpl",
	"templates/init_module/delete.go.tmpl",
	"templates/init_module/post.go.tmpl",
	"templates/init_module/put.go.tmpl",
}

var sqlc_templates = []Template{
	{
		template_path: "templates/optional/sqlc/sqlc.yaml.tmpl",
		output_dir:    "/db/",
	},
	{
		template_path: "templates/optional/sqlc/README.txt.tmpl",
		output_dir:    "/db/",
	},
	{
		template_path: "templates/optional/sqlc/functions.sql.tmpl",
		output_dir:    "/db/sql/",
	},
	{
		template_path: "templates/optional/sqlc/__module_name__.sql.tmpl",
		output_dir:    "/db/sql/",
	},
}
