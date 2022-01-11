package task_module

import (
	"fmt"
	"example_with_modules/settings"
	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "Task"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/task")

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/task", GetTasks) 
	// Get By Id
	api.Get("/task/:id", GetTask) 
	// Create
	api.Post("/task", CreateTask) 
	// Update With Id
	api.Put("/task/:id", UpdateTask)
	// Delete With Id
	api.Delete("/task/:id", DeleteTask) 
}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=6,max=50"`
	ExampleTitle string `json:"example_title" validate:"required,min=6,max=50"`
}
