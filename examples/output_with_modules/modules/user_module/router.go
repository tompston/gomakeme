package user_module

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/user", GetAllUsers) 

	// Get By Id
	api.Get("/user/:id", GetUser) 

	// Create
	api.Post("/user", CreateUser) 

	// Update With Id
	api.Put("/user/:id", UpdateUser)
	
	// Delete With Id
	api.Delete("/user/:id", DeleteUser) 
}