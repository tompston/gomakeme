# **gomakeme**

## Generate boilerplate + endpoints for Fiber REST APIs.

> Never spend 6 minutes doing something by hand when you can spend 1 week to automate it

Currently there are two options for the project that you could generate:

1. Minimal setup with no endpoints

2. Minimal setup that use modules

The option you choose is based on the `gomakeme.yml` config file. If there are modules specified in the config file, they will be added to the project.

For example, if your config file has some defined modules, like

```yml
modules: [User, Task]
```

The `user_module` will be created in the `modules` directory.

### Module

A single module will hold 5 basic CRUD endpoints, that will be automatically available to the server. The generated controllers are held in seperate files, where the name of the file indicates what type of http request is associated with it. The controllers can also hold common functions, such as

- validation of url params, if there are any
- validation and custom error messages for the sent payloads
- database connection ( if set to true in the config file )

Once the directory for the `user_module` is created, it won't be updated if you run the `go generate` again, so you can edit them.

## How to generate a new REST API

1. Clone the repo

2. Change the `gomakeme.yml` config file

3. Run `go generate`

4. cd into the project and run

```bash
go mod tidy
go mod download
go run main.go
# + change the .env vars
```

## Notes

- Why Fiber? Because it's one of the [fastest](https://www.techempower.com/benchmarks/) Go frameworks currently + takes inspiration from Express

- The project does not include any ORMs currently, as there are too many options to choose from.

- Taking a bit of inspiration from Nest.js, we bundle all of the logic of a single module ( router + controllers ) into one package.

- Only the `router/project_modules.go` file will be updated on the next `go generate` runs. All of the other files won't be touched

- There are probably bugs

<!--

# wsl

export PATH=$PATH:/usr/local/go/bin
GOOS=linux go build -o main .

export PATH=$PATH:/usr/local/go/bin
GOOS=linux go build -o ./gomakeme

export PATH=$PATH:/usr/local/go/bin
GOOS=linux go build -o ./gomakeme_linux


export PATH=$PATH:/usr/local/go/bin
GOOS=linux GOARCH=amd64 go build -o ./gomakeme_linux_amd64

-->
