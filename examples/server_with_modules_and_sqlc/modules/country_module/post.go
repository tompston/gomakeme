package country_module

import (
	res "server_with_modules_and_sqlc/utils/response"
	"server_with_modules_and_sqlc/utils/validate"

	"github.com/gofiber/fiber/v2"
	"server_with_modules_and_sqlc/settings/database"
)

func CreateCountry(c *fiber.Ctx) error {

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	// query goes here

	return res.ResponseSuccess(c, data, res.CreatedMessage(module_name))
}