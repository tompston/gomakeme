# gomakeme

Generate boilerplate + endpoints for Fiber REST APIs.

#### Table of contents
* [Install](#install)
* [Generate the REST API](#generate-the-rest-api)
* [Explanation](#explanation)
* [Config file](#config-file)
* [Notes](#notes)

## Install

```bash
# go version >= 1.17
go install github.com/tompston/gomakeme@latest

# go version < 1.17
go get github.com/tompston/gomakeme

# or clone the repo
git clone https://github.com/tompston/gomakeme.git
```

## Generate the REST API

1. Create the `gomakeme.yml` config file, copy example settings from this repo and update values
2. run `gomakeme` in the directory which holds the yaml config
3. cd into the created project and run

```bash
go mod tidy
go mod download
go run main.go
# + change the .env vars
```

## Explanation

It seems like one of the boring / repetitive parts of writing basic REST APIs is the boilerpate for the endpoints.

Want to add a new table to the database and create endpoints for it? Before starting to do so, you need to write the new routes and controllers that deal with the common CRUD opetations.

Want to add another table? Repat the same process again, but change only the name to the new table.

So this stuff is boring and manual. That's why you can automate it.

<h3 align="center">
    Once generated, only the "router/project_modules.go" file will be updated on the next `gomakeme` runs. All of the other files won't be touched.
</h3>

## Config file

**Commented out lines + modules lines are optional**  
 **All of the other lines are mandatory!**

Currently there are three main options for the project that you could generate:

1. Minimal server with no endpoints
2. Server that uses modules
3. Server that uses modules + sqlc as an ORM

_\* examples of all three options can be found inside examples dir_

The option you choose is based on the `gomakeme.yml` config file. If there are modules specified in the config file, they will be added to the project.

For example, if your config file has some defined modules, like

```yml
modules: [User, Task]
```

The `user_module` will be created in the `modules` directory.

#### sqlc: true

If true, then sqlc config files and sql files for modules will also be generated inside `/db` dir. The sql files that would be used by sqlc would include

- table with the name of the module + 3 frequently used columns (id, created_at, updated_at)
- 5 CRUD queries with placeholder values

#### include_db_snippet: true

If true, includes a snippet inside the controllers that can create a connection to the database

### Module

A single module will hold 5 basic CRUD endpoints, that will be automatically available to the server. The generated controllers are held in seperate files, where the name of the file indicates what type of http request is associated with it. The controllers can also hold common functions, such as

- validation of url params, if there are any
- validation and custom error messages for the sent payloads
- database connection ( if set to true in the config file )

Once the directory for the `user_module` is created, it won't be updated if you run the `gomakeme` again, so you can edit them.

<a name="Notes"/>

## Notes

- Why Fiber? Because it's one of the [fastest](https://www.techempower.com/benchmarks/) Go frameworks currently + takes inspiration from Express
- Taking a bit of inspiration from Nest.js, we bundle all of the logic of a single module ( router + controllers ) into one package.
- There are probably bugs

<!--

# wsl
export PATH=$PATH:/usr/local/go/bin

GOOS=linux go build -o main .
GOOS=linux go build -o ./gomakeme
GOOS=linux go build -o ./gomakeme_linux
GOOS=linux GOARCH=amd64 go build -o ./gomakeme_linux_amd64
GOOS=windows go build -o ./bin/gomakeme_win

-- publishing
https://go.dev/doc/modules/publishing

go mod tidy
git add .
git commit -m "changes for v0.0.3 - shortened the response package import name in the controllers to be more elegant + added comment that indicates where the ORM / SQL queries go"
git tag v0.0.3
git push origin v0.0.3
GOPROXY=proxy.golang.org go list -m github.com/tompston/gomakeme@v0.0.3


-- testing script
go run main.go
cd change_my_name
go mod tidy
go mod download
code .
go run main.go

-->