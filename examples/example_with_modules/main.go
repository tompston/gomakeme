package main

import (
	"example_with_modules/router"
	"example_with_modules/settings"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/gofiber/helmet/v2"
)

func main() {

	app := fiber.New()    // create the server
	app.Use(logger.New()) // logs the requests
	//app.Use(csrf.New())   // csrf protection 
	//app.Use(helmet.New()) // security headers

	router.Url(app) // use the routes that are defiend by the Url function

	port, _ := strconv.Atoi(settings.Config("GOFIBER_PORT")) // import the .env int
	app.Listen(fmt.Sprintf(":%d", port))					// run the server
}
