package user_module

import (
	"fmt"
	"server_with_modules_and_sqlc/settings"
	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "User"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/user")

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/user", GetUsers) 
	// Get By Id
	api.Get("/user/:id", GetUser) 
	// Create
	api.Post("/user", CreateUser) 
	// Update With Id
	api.Put("/user/:id", UpdateUser)
	// Delete With Id
	api.Delete("/user/:id", DeleteUser) 
}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=1,max=20"`
	ExampleTitle string `json:"example_title" validate:"required,min=5,max=50"`
}
