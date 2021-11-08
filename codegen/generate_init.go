package codegen

import (
	"fmt"
	"gomakeme/input"
)

var init_project_template_path = "codegen/templates/init_server/"

var init_project_dirs = [...]string{
	"/",
	"/settings/database/",
	"/utils/response/",
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
	"router/router.go.tpl",
}

// loop over dirs specified in the  "init_project_dirs" to create them
func GenerateInitProjectDirs(project_name string) {
	for i := 0; i < len(init_project_dirs); i++ {
		CreateFolder(fmt.Sprintf("%s/%s", project_name, init_project_dirs[i]))
	}
}

// loop over "init_project_templates" and create the output files
// when generatin the init project, we pass down the struct that we defined in the
// input, as we need access to all of the defined modules that are mentioned in the array
func GenerateInitFileBatch(global_project_data input.Project) {

	for i := 0; i < len(init_project_templates); i++ {
		proj_name := global_project_data.ProjectInfo.ProjectName
		temp_path := init_project_templates[i]
		temp_name, output_path := GenerateTemplateNameAndOutput(temp_path, true)
		full_output_path := fmt.Sprintf("%s/%s", proj_name, output_path)
		full_template_path := fmt.Sprintf("%s%s", init_project_template_path, temp_path)
		ExecuteTemplate(temp_name, full_template_path, temp_funcs, full_output_path, global_project_data)
		fmt.Println("Created ", output_path)
	}
}

// function that creates the file which imports the modules.
// As we should not run the function that creates the project every time, we create a seperate func that
// will always update the imported modules
func UpdateModuleImporter(g input.Project) {

	proj_name := g.ProjectInfo.ProjectName
	temp_path := "router/project_modules.go.tpl"
	temp_name, output_path := GenerateTemplateNameAndOutput(temp_path, true)
	full_output_path := fmt.Sprintf("%s/%s", proj_name, output_path)
	full_template_path := fmt.Sprintf("%s%s", init_project_template_path, temp_path)
	ExecuteTemplate(temp_name, full_template_path, temp_funcs, full_output_path, g)
	fmt.Println("Updated ", output_path)
}

// this function will create the empty init folders and then populate them with the templates that
// you created inside the init_server dir.
// Just like in modules, add a check that will validate if the project exists to run this func.
func GenerateInitProject(global_project_data input.Project) {

	p_name := global_project_data.ProjectInfo.ProjectName
	p_path := fmt.Sprintf("./%s", p_name)

	// if the provided project name dir already exists, do not update
	// the init dirs.
	if !PathExists(p_path) {
		GenerateInitProjectDirs(p_name)
		GenerateInitFileBatch(global_project_data)
	} else {
		fmt.Println("Project already exists!")
	}
}
