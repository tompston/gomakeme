package user_module

import (
	"example_with_modules/utils/response"
	"example_with_modules/utils/validate"

	"github.com/gofiber/fiber/v2"
	"example_with_modules/settings/database"
)

func CreateUser(c *fiber.Ctx) error {

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	return response.ResponseSuccess(c, data, response.CreatedMessage(module_name))
}