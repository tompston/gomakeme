package codegen

import (
	"fmt"
	"strings"

	"github.com/tompston/gomakeme/input"
)

// main func that generates stuff based on config file
func GenerateOptional(g *input.Project) {

	// if sqlc : true
	if g.ProjectInfo.SQLC {
		AddSQLC(g)
	}
}

// this func will be wrapped in conditional statement
func AddSQLC(g *input.Project) {

	project_name := g.ProjectInfo.ProjectName
	sqlc_config_path := fmt.Sprintf("./%s%s", project_name, "/db/sqlc.yaml") // returns ./change_my_name/db/sqlc.yaml

	// if
	//  - sqlc.yaml file does not exist (meaning has not been generated before)
	//  - debug mode = true
	// create init dirs
	if !PathExists(sqlc_config_path) || debug_mode {
		CreateFolder(fmt.Sprintf("%s/%s", project_name, "/db/sqlc/"))
		CreateFolder(fmt.Sprintf("%s/%s", project_name, "/db/sql/"))
	}

	// if sqlc.yaml file exists
	for i := 0; i < len(sqlc_templates); i++ {

		output_dir := sqlc_templates[i].output_dir
		templ_path := sqlc_templates[i].template_path

		// if there are modules provided in the array
		if len(g.Modules) != 0 {
			output_path := fmt.Sprintf("%s%s", project_name, output_dir) // point the output path
			template_name, output_file := GenerateTemplateNameAndOutput(templ_path, false)
			full_output_path := fmt.Sprintf("%s%s", output_path, output_file)

			// if the current template is the one which should be repeated based on
			// how many values are provided in the modules array
			if strings.Contains(templ_path, "__module_name__") {

				// for each module in modules array
				for x := 0; x < len(g.Modules); x++ {

					module := g.Modules[x]
					data := ConvertInputStructToCodegenStruct(g, module) // convert to correct struct
					full_output_path := fmt.Sprintf("%s%s", output_path, output_file)
					module_name := data.Module.ModuleName
					module_name_lwc_plural := ConvertToLowercase(ConvertToPlural(module_name))
					full_output_path = strings.Replace(full_output_path, "__module_name__", module_name_lwc_plural, -1)

					// if the sql file for the module does not exist or debug mode, create the template
					if !PathExists(full_output_path) || debug_mode {
						ExecuteTemplate(template_name, templ_path, temp_funcs, full_output_path, data)
						fmt.Printf("Created  %s.sql!\n", module_name_lwc_plural)
					}
				}
			}

			// if file does not exist already, generate it
			if !PathExists(full_output_path) || debug_mode {
				// skip the template that includes the replacable string in the name
				if !strings.Contains(templ_path, "__module_name__") {
					ExecuteTemplate(template_name, templ_path, temp_funcs, full_output_path, g)
					fmt.Println("Created ", full_output_path) // log what is created
				}
			}
		}
	}
}
