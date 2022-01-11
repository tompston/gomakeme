package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
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

// ------- Structs -------

// Project struct will be passed to the templates and will hold all of the information about the project.
// It consists of two parts currently - ProjectInfo and Module
type Project struct {
	ProjectInfo   ProjectInfo
	Module        Module
	ProjectConfig ProjectConfig
}

// struct that will hold all of the information about a single module that you want to create.
// currently only the name is needed
type Module struct {
	ModuleName string
}

// ProjectName
//  will be used in the go.mod, as the name of the folder in which your project will be
// 	and as the first part of imports for other modules
type ProjectInfo struct {
	ProjectName      string
	Port             string `yaml:"port"`
	GoVersion        string `yaml:"go_version"`
	IncludeDbSnippet string `yaml:"include_db_snippet"`
}

// struct that will be used in the .env template.
type ProjectConfig struct {
	// db variables
	PostgresHost     string
	PostgresUser     string
	PostgresPassw    string
	PostgresDb       string
	PostgresPort     string
	PostgresSSL      string
	PostgresTimezone string
	PostgresData     string
}

// ------- Template functions -------

// functions that will be passed to the templates.
var temp_funcs = template.FuncMap{
	"convertToTitle":     strings.Title,
	"convertToLowercase": strings.ToLower,
	"db_conn_snippet":    DbConnSnippet,
	"validate_url_param": ValidateUrlParam,
	"validate_payload":   ValidatePayload,
}

// --- Funcs that return the code snippet that can be used in the template

func DbConnSnippet() string {
	return `
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db`
}

func ValidateUrlParam() string {
	return `

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.ResponseError(c, nil, response.ParamIsNotIntMessage)
	}
	_ = id`
}

func ValidatePayload() string {
	return `

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return response.ResponseError(c, nil, err.Error())
	}`
}
