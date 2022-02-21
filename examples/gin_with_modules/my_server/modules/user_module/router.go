package user_module

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_server/settings"
)

var data = ""
var module_name = "User"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/user")

func Routes(app *gin.RouterGroup) {

	module := app.Group("user")

	{
		module.GET("/", GetUsers)
		module.GET("/:id", GetUser)
		module.POST("/", CreateUser)
		module.PUT("/", UpdateUser)
		module.DELETE("/", DeleteUser)
	}
}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=1,max=20"`
	ExampleTitle string `json:"example_title" validate:"required,min=5,max=50"`
}
