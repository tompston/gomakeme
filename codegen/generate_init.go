package codegen

import (
	"fmt"

	"github.com/tompston/gomakeme/input"
)

// loop over dirs specified in the  "init_project_dirs" to create them
func GenerateInitProjectDirs(project_name string) {
	for i := 0; i < len(init_project_dirs); i++ {
		CreateFolder(fmt.Sprintf("%s/%s", project_name, init_project_dirs[i]))
	}
}

// loop over "init_project_templates" and create the output files when generating
// the init project, we pass down the struct that we defined in the input, as we
// need access to all of the defined modules that are mentioned in the array
func GenerateInitFileBatch(g input.Project) {

	for i := 0; i < len(init_project_templates); i++ {

		project_name := g.ProjectInfo.ProjectName
		template_path := init_project_templates[i]
		template_name, output_path := GenerateTemplateNameAndOutput(template_path, true)
		full_output_path := fmt.Sprintf("%s/%s", project_name, output_path)
		full_template_path := fmt.Sprintf("%s%s", init_project_template_path, template_path)

		ExecuteTemplate(template_name, full_template_path, temp_funcs, full_output_path, g)
		fmt.Println("Created ", output_path)
	}
}

// function that creates the file which imports the modules.
// As we should not run the function that creates the project every time, we create
// a seperate function that will always update the imported modules
func UpdateModuleImporter(g input.Project) {

	project_name := g.ProjectInfo.ProjectName
	template_path := "router/module_importer.go.tpl"
	template_name, output_path := GenerateTemplateNameAndOutput(template_path, true)
	full_output_path := fmt.Sprintf("%s/%s", project_name, output_path)
	full_template_path := fmt.Sprintf("%s%s", init_project_template_path, template_path)

	ExecuteTemplate(template_name, full_template_path, temp_funcs, full_output_path, g)
	fmt.Println("Updated ", output_path)
}

// create the empty init folders and then populate them with the templates that
// you created inside the init_server dir.
func GenerateInitProject(g input.Project) {

	project_name := g.ProjectInfo.ProjectName
	project_path := fmt.Sprintf("./%s", project_name)

	// if the provided project name dir already exists or debug_mode is true, do
	// not update the init dirs.
	if !PathExists(project_path) || debug_mode {
		GenerateInitProjectDirs(project_name)
		GenerateInitFileBatch(g)
	}
}
