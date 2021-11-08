# **gomakeme**

## Generate boilerplate + endpoints for Fiber REST APIs.

Currently there are two options for the project that you could generate:

1. Minimal setup with no endpoints

2. Minimal setup that use modules

The option you choose is based on the `gomakeme.yml` config file

If there are modules specified in the config file, they will be added to the project.

For example, if your config file has some defined modules, like

```yml
modules: [User, Task]
```

The `user_module` will be created in the `modules` directory.

A single module will hold 5 basic CRUD endpoints, that will be automatically available to the server.

Once the directory for the `user_module` is created, it won't be updated / touched if you run the generator again, so you can edit them.

_Examples of generated outputs can be seen in the `examples` directory_

## How to generate a server

1. Clone the repo

2. Change the `gomakeme.yml` config file

3. Run `go generate`

4. cd into the project and run `go mod tidy`

5. run the server with `go run main.go`

## Notes

- Why Fiber? Because it's one of the [fastest](https://www.techempower.com/benchmarks/) Go frameworks currently + takes inspiration from Express

- The POST, PUT, DELETE endpoints won't work initially due to CSRF middleware in `main.go`.

- Taking a bit of inspiration from Nest.js and to keep the global state management simple, we bundle most of the logic of a single module into seperate packages.

- Only the `router/project_modules.go` file will be updated on the next `go generate` runs. All of the other files won't be touched.

- There are probably bugs.

- Don't plan to update this much in the future.
