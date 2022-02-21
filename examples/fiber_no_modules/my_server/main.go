package main

import (
	"fmt"
	"my_server/router"
	"my_server/settings"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func main() {

	app := fiber.New()    // create the server
	app.Use(logger.New()) // logs the requests
	app.Use(helmet.New()) // security headers

	router.Url(app) // use the routes that are defiend by the Url function

	port, _ := strconv.Atoi(settings.Config("GOLANG_PORT")) // import the .env int
	app.Listen(fmt.Sprintf(":%d", port))                    // run the server
}
