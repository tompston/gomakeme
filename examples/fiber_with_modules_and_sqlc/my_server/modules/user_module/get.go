package user_module

import (
	res "my_server/utils/response"

	"github.com/gofiber/fiber/v2"
	"my_server/settings/database"
	"strconv"
)

func GetUsers(c *fiber.Ctx) error {

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.Response(c, 500, nil, err.Error())
	}
	defer db.Close()
	_ = db

	// query goes here

	return res.Response(c, 200, data, "Found Users!")
}

func GetUser(c *fiber.Ctx) error {

	// validate url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return res.Response(c, 400, nil, res.ParamIsNotIntMessage)
	}
	_ = id

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		return res.Response(c, 500, nil, err.Error())
	}
	defer db.Close()
	_ = db

	// query goes here

	return res.Response(c, 200, data, "Found User!")
}
