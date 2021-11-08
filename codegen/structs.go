package codegen

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
	ProjectName string
	Port        string `yaml:"port"`
	GoVersion   string `yaml:"go_version"`
}

// struct that will be used in the .env template.
type ProjectConfig struct {
	// database variables
	PostgresHost     string
	PostgresUser     string
	PostgresPassw    string
	PostgresDb       string
	PostgresPort     string
	PostgresSSL      string
	PostgresTimezone string
	PostgresData     string
	// add some other variables later, maybe.
}
