package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
)

func ConvertToLowercase(input string) string {
	return strings.ToLower(input)
}

func ConvertToPlural(input string) string {
	return pluralize.NewClient().Plural(input)
}

func ConvertToTitle(input string) string {
	return strings.Title(input)
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

// return the path to the "modules" folder
// 		("my_project") --> "my_project/modules"
func ProjectModuleDirPath(project_name string) string {
	return fmt.Sprintf("%s/%s", ConvertToLowercase(project_name), ModuleOutputDir)
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

// --- Funcs that return the code snippet that can be used in the template

func DbConnSnippet() string {
	return `
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db`
}

func ValidateUrlParam() string {
	return `

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}
	_ = id`
}

func ValidatePayload() string {
	return `

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}`
}

// ------- Template functions -------

// functions that will be passed to the templates
var temp_funcs = template.FuncMap{

	// -- funcs that convert input
	// note that the key string should be the same as the name
	// of the function to keep things simple and shared
	"ConvertToTitle":     ConvertToTitle,
	"ConvertToLowercase": ConvertToLowercase,
	"ConvertToPlural":    ConvertToPlural,

	// -- funcs that return a predefined string
	"db_conn_snippet":    DbConnSnippet,
	"validate_url_param": ValidateUrlParam,
	"validate_payload":   ValidatePayload,
}
