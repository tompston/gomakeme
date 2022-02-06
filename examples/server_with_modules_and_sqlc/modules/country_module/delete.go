package country_module


import (
	res "server_with_modules_and_sqlc/utils/response"

	"github.com/gofiber/fiber/v2"
	"strconv"
	"server_with_modules_and_sqlc/settings/database"
)

func DeleteCountry(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.ResponseError(c, nil, res.ParamIsNotIntMessage)
	}
	_ = id
	
	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.ResponseError(c, nil, err.Error())
	}
	defer db.Close()
	_ = db

	// query goes here

	return res.ResponseSuccess(c, data, res.DeletedMessage(module_name))
}
