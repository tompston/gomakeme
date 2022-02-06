package country_module

import (
	"fmt"
	"server_with_modules/settings"
	"github.com/gofiber/fiber/v2"
)

var data = ""
var module_name = "Country"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/country")

func Routes(app *fiber.App, api fiber.Router){

	// Get All 
	api.Get("/country", GetCountries) 
	// Get By Id
	api.Get("/country/:id", GetCountry) 
	// Create
	api.Post("/country", CreateCountry) 
	// Update With Id
	api.Put("/country/:id", UpdateCountry)
	// Delete With Id
	api.Delete("/country/:id", DeleteCountry) 
}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=1,max=20"`
	ExampleTitle string `json:"example_title" validate:"required,min=5,max=50"`
}
