package user_module

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"my_server/settings"
)

var data = ""
var module_name = "User"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/user")

func Routes(app *fiber.App, api fiber.Router) {

	api.Get("/user", GetUsers)
	api.Get("/user/:id", GetUser)
	api.Post("/user", CreateUser)
	api.Put("/user/:id", UpdateUser)
	api.Delete("/user/:id", DeleteUser)

}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=1,max=20"`
	ExampleTitle string `json:"example_title" validate:"required,min=5,max=50"`
}
