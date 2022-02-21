package user_module

import (
	res "my_server/utils/response"

	"github.com/gin-gonic/gin"
	"my_server/settings/database"
	"strconv"
)

func GetUsers(c *gin.Context) {

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		res.Response(c, 500, nil, res.DbConnErrorMessage(err.Error()))
		return
	}
	defer db.Close()
	_ = db

	// query goes here

	res.Response(c, 200, data, "Found Users!")
}

func GetUser(c *gin.Context) {

	// validate url param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Response(c, 400, nil, res.ParamIsNotIntMessage)
		return
	}
	_ = id

	// get db connection
	db, err := database.GetDbConn()
	if err != nil {
		res.Response(c, 500, nil, res.DbConnErrorMessage(err.Error()))
		return
	}
	defer db.Close()
	_ = db

	// query goes here

	res.Response(c, 200, data, "Found User!")
}
