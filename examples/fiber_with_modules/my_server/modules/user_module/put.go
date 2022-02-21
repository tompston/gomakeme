package user_module

import (
	res "my_server/utils/response"
	"my_server/utils/validate"

	"github.com/gofiber/fiber/v2"
	"my_server/settings/database"
	"strconv"
)

func UpdateUser(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.Response(c, 400, nil, res.ParamIsNotIntMessage)
	}
	_ = id

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		return res.Response(c, 400, nil, err.Error())
	}

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.Response(c, 500, nil, err.Error())
	}
	defer db.Close()
	_ = db

	// query goes here

	return res.Response(c, 200, data, "Updated User!")
}
