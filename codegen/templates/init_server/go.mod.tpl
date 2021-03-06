{{ $project_name := .ProjectInfo.ProjectName -}}
{{ $go_version := .ProjectInfo.GoVersion -}}

module {{$project_name}}

go {{ $go_version }}






{{- if ( eq .ProjectInfo.Framework "fiber") }}
require (
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/gofiber/fiber/v2 v2.18.0
	github.com/gofiber/helmet/v2 v2.2.1
	github.com/joho/godotenv v1.3.0
	github.com/klauspost/compress v1.13.5 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.2
	golang.org/x/sys v0.0.0-20210906170528-6f6e22806c34 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)
{{- end }}





{{- if ( eq .ProjectInfo.Framework "gin") }}
require (
	github.com/gin-gonic/gin v1.7.7
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.10.2
	gopkg.in/go-playground/validator.v9 v9.31.0
)
{{- end }}