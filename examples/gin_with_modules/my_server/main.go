package main

import (
	"my_server/router"
	"my_server/settings"

	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default() // create the server

	router.Endpoints(app) // use endpoints

	port, _ := strconv.Atoi(settings.Config("GOLANG_PORT")) // import the .env int
	app.Run(fmt.Sprintf("localhost:%d", port))              // run the server
}
