package user_module

import (
	res "my_server/utils/response"
	"my_server/utils/validate"

	"github.com/gin-gonic/gin"
	"my_server/settings/database"
)

func CreateUser(c *gin.Context) {

	// validate the sent json object
	payload := new(ExampleStruct) // define which struct you want to get
	if err := validate.ValidatePayload(c, payload); err != nil {
		res.Response(c, 400, nil, err.Error())
		return
	}

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		res.Response(c, 500, nil, res.DbConnErrorMessage(err.Error()))
		return
	}
	defer db.Close()
	_ = db

	// query goes here

	res.Response(c, 200, data, "Created User!")
}
