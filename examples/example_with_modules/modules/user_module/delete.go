package user_module

import (
	"example_with_modules/utils/response"

	"github.com/gofiber/fiber/v2"
	"strconv"
	"example_with_modules/settings/database"
)

func DeleteUser(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.ResponseError(c, nil, response.ParamIsNotIntMessage)
	}
	_ = id
	
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return response.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	return response.ResponseSuccess(c, data, response.DeletedMessage(module_name))
}