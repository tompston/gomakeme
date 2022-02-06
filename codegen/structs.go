package codegen

// Project struct will be passed to the templates and will hold all of the information about the project.
// It consists of two parts currently - ProjectInfo and Module
type Project struct {
	ProjectInfo   ProjectInfo
	Module        Module
	ProjectConfig ProjectConfig
}

// name of the module
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
	SQLC             bool   `yaml:"sqlc"`
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

type Template struct {
	template_path string
	output_dir    string // output dir into which the template will be created
}
