package user_module

import (
	res "server_with_modules/utils/response"
	"server_with_modules/utils/validate"

	"github.com/gofiber/fiber/v2"
	"strconv"
	"server_with_modules/settings/database"
)

func UpdateUser(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}
	_ = id

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

	return res.ResponseSuccess(c, data, res.UpdatedMessage(module_name))
}