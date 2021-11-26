package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ConvertToLowecase(input string) string {
	return strings.ToLower(input)
}

func ConvertToLowercasePlural(input string) string {
	return fmt.Sprintf("%s%s", strings.ToLower(input), "s")
}

// if you pass a nested path of folders that do not exist, this function will also create those folders
func CreateFolder(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// creates the directory that will hold all of the files
// 	"task" -> "task_module"
func ModuleDirName(module_name string) string {
	return strings.ToLower(fmt.Sprintf("%s%s", module_name, "_module"))
}

// get the path to the folder which will hold the output of the generated code for the module
// 	"todo" -> "modules/todo_module"
func ModuleDirPath(module_name string) string {
	return fmt.Sprintf("%s/%s", "modules", ModuleDirName(module_name))
}

// create the folder which will hold the output
//   passing in "todo" will create "modules/todo_module" folder
func CreateModuleDir(module_name string) {
	os.Mkdir(ModuleDirPath(module_name), 0755)
}

// return the full path to the file to which you will write the generated output
// 	"todo", "_endpoint.txt"	--> "modules/todo_module/todo_endpoint.txt"
func ModuleOutputFile(module_name string, file_ends_with string) string {
	output_dir := ModuleDirPath(module_name)
	output_filename := fmt.Sprintf("%s%s", ConvertToLowecase(module_name), file_ends_with)
	output := fmt.Sprintf("%s/%s", output_dir, output_filename)
	return output
}

// return the path to the "modules" folder
// 		("my_project") --> "my_project/modules"
func ProjectModuleDirPath(project_name string) string {
	return fmt.Sprintf("%s/%s", ConvertToLowecase(project_name), ModuleOutputDir)
}

// return the full path to the folder that holds all of the files associated with the a single module of the project
// 		("my_Project", "Task") --> my_project/modules/task_module/
func ModuleDir(project_name string, module_name string) string {
	return fmt.Sprintf("%s/%s/", ProjectModuleDirPath(project_name), ModuleDirName(module_name))
}

// pass in the relative path to the dir you want to check.
// if the dir exists, functin will return true boolean
func PathExists(path string) bool {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// fmt.Println(path, "does not exist")
		return !os.IsNotExist(err)
	}

	return true
}
