package country_module

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"my_server/settings"
)

var data = ""
var module_name = "Country"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/country")

func Routes(app *fiber.App, api fiber.Router) {

	api.Get("/country", GetCountries)
	api.Get("/country/:id", GetCountry)
	api.Post("/country", CreateCountry)
	api.Put("/country/:id", UpdateCountry)
	api.Delete("/country/:id", DeleteCountry)

}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=1,max=20"`
	ExampleTitle string `json:"example_title" validate:"required,min=5,max=50"`
}
