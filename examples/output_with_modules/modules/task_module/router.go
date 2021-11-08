package task_module

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/task", GetAllTasks) 

	// Get By Id
	api.Get("/task/:id", GetTask) 

	// Create
	api.Post("/task", CreateTask) 

	// Update With Id
	api.Put("/task/:id", UpdateTask)
	
	// Delete With Id
	api.Delete("/task/:id", DeleteTask) 
}