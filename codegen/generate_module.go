package codegen

import (
	"fmt"
	"gomakeme/input"
)

// name of the folder in the project to which the new modules will be added
var ModuleOutputDir = "modules"

// generate the dir inside the project that will hold the files for the module
func GenerateModuleDir(g Project) {
	CreateFolder(ModuleDir(g.ProjectInfo.ProjectName, g.Module.ModuleName))
}

// you can't pass down the Project struct from input package due to the fact that
// the template won't be able to use the ModuleName as a variable inside the template,
// as it's an array. To avoid passing down an array that would break the templates, we assign
// the value of the array to the Project struct that is defined in the current package.
func GenerateModuleFile(g Project, template_path string) {

	p_name := g.ProjectInfo.ProjectName
	m_name := g.Module.ModuleName
	o_path := ModuleDir(p_name, m_name)
	template_name, output_file := GenerateTemplateNameAndOutput(template_path, false)

	full_output_path := fmt.Sprintf("%s%s", o_path, output_file)
	ExecuteTemplate(template_name, template_path, temp_funcs, full_output_path, g)
}

// Generate a new module, if the folder for it does not exist already
func GenerateNewModule(g Project, template_path string) {

	p_name := g.ProjectInfo.ProjectName
	m_name := g.Module.ModuleName
	m_path := ModuleDir(p_name, m_name)
	m_path_relative := fmt.Sprintf("./%s", m_path)

	// if the module mentioned in the config does not exist or debug_mode is true, create it
	if !PathExists(m_path_relative) || debug_mode {

		// create the dir that will hold the files for the module
		GenerateModuleDir(g)

		// create the files inside the dir
		for i := 0; i < len(module_files); i++ {
			GenerateModuleFile(g, module_files[i])
		}

		fmt.Println("* Created module", m_name)
	}
}

// Passing down the Project struct from the input package that is populated with the values from the yaml file.
// For each value that you define in the modules array, create a new struct that will hold seperated values so
// that a template would have access to one module at a time.
// Once you have a struct with seperated values, create it
func GenerateModulesFromConfig(conf *input.Project) {

	// if the Modules array is not empty, generate them
	if len(conf.Modules) != 0 {
		fmt.Println("Generating Modules!")
		for i := 0; i < len(conf.Modules); i++ {

			yaml_data := Project{
				ProjectInfo:   ProjectInfo(conf.ProjectInfo),
				ProjectConfig: ProjectConfig(conf.ProjectConfig),
				Module: Module{
					ModuleName: conf.Modules[i]},
			}

			GenerateNewModule(yaml_data, "codegen/templates/init_module/module_controller.go.tpl")
			// GenerateNewModule(yaml_data, "templates/init_module/module_controller.go.tpl")
		}
	}
}
