package input

type Project struct {
	ProjectInfo   ProjectInfo `yaml:"project_info"`
	Modules       []string    `yaml:"modules"`
	ProjectConfig ProjectConfig
}

type ProjectInfo struct {
	ProjectName      string `yaml:"name"`
	Port             string `yaml:"port"`
	GoVersion        string `yaml:"go_version"`
	Framework        string `yaml:"framework"`
	IncludeDbSnippet string `yaml:"include_db_snippet"`
	SQLC             bool   `yaml:"sqlc"`
}

// pre-defined and changed in the .env file so that
// there are no secrets stored in the gomakeme.yml file
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
}

var placeholder_config = ProjectConfig{
	PostgresHost:     "localhost",
	PostgresUser:     "change_me",
	PostgresPassw:    "change_me",
	PostgresDb:       "change_me",
	PostgresPort:     "5432",
	PostgresSSL:      "disable",
	PostgresTimezone: "Europe/Helsinki",
	PostgresData:     "/var/lib/postgresql/data",
}
