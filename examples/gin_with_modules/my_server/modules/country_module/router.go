package country_module

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_server/settings"
)

var data = ""
var module_name = "Country"
var BASE = fmt.Sprintf(settings.Config("BASE_URL") + "/api" + "/country")

func Routes(app *gin.RouterGroup) {

	module := app.Group("country")

	{
		module.GET("/", GetCountries)
		module.GET("/:id", GetCountry)
		module.POST("/", CreateCountry)
		module.PUT("/", UpdateCountry)
		module.DELETE("/", DeleteCountry)
	}
}

type ExampleStruct struct {
	ExampleId    int    `json:"example_id" validate:"required,min=1,max=20"`
	ExampleTitle string `json:"example_title" validate:"required,min=5,max=50"`
}
