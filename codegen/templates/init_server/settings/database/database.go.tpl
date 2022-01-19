{{ $project_name := .ProjectInfo.ProjectName -}}

package database

import (
	"database/sql"
	"fmt"
	"os"

	"{{$project_name}}/settings"

	_ "github.com/lib/pq"
)

// create the dsn string from variables that are specified in the .env file
func DsnString() string {
	// load all the env variables
	settings.LoadEnvFile()

	HOST := os.Getenv("HOST")
	USER := os.Getenv("POSTGRES_USER")
	PASSW := os.Getenv("POSTGRES_PASSWORD")
	DBNAME := os.Getenv("POSTGRES_DB")
	PORT := os.Getenv("HOST_PORT")
	SSLMODE := os.Getenv("SSLMODE")
	TZ := os.Getenv("TZ")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", HOST, USER, PASSW, DBNAME, PORT, SSLMODE, TZ)

	return dsn
}

// get a db connection with the inbuilt package
func GetDbConn() (*sql.DB, error) {

	dsn := DsnString()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, err
}
