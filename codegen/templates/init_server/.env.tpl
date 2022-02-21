{{ $pg_host := 	.ProjectConfig.PostgresHost -}}
{{ $pg_user := 	.ProjectConfig.PostgresUser -}}
{{ $pg_passw := .ProjectConfig.PostgresPassw -}}
{{ $pg_db := 	.ProjectConfig.PostgresDb -}}
{{ $pg_port := 	.ProjectConfig.PostgresPort -}}
{{ $pg_ssl := 	.ProjectConfig.PostgresSSL -}}
{{ $pg_tz := 	.ProjectConfig.PostgresTimezone -}}
{{ $pg_data := 	.ProjectConfig.PostgresData -}}
{{ $port := 	.ProjectInfo.Port -}}

# --- POSTGRES DATABASE VARAIBLES 
# Don't change the name of the key, as it's imported in the settings + is the default
# variable that is recognized by the Postgres docker image
HOST={{$pg_host}}
POSTGRES_USER={{$pg_user}}
POSTGRES_PASSWORD={{$pg_passw}}
POSTGRES_DB={{$pg_db}}
HOST_PORT={{$pg_port}}
SSLMODE={{$pg_ssl}}
TZ={{$pg_tz}}
PGDATA={{$pg_data}}

# --- GOLANG VARIABLES
GOLANG_PORT={{$port}}
# Used in main.go as the port + in docker-compose
BASE_URL=http://localhost:{{$port}}
# Used when you want to link to other pages of the API (like pagitation). Needs to match the GOLANG_PORT.
FRONTEND_URL=http://localhost:3000
# Variable that can be used in main.go for CORS, so that you could make api requests from the frontend, when running locally.
