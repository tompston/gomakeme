package main

import (
	"output_with_modules/router"
	"output_with_modules/settings"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()    // create the server
	app.Use(logger.New()) // logs the requests
	app.Use(csrf.New())   // csrf protection, (note that POST, PUT, DELETE wont work initially with this)
	app.Use(helmet.New()) // security headers

	router.Url(app) // use the routes that are defiend by the Url function

	port, _ := strconv.Atoi(settings.Config("GOLANG_PORT")) // import the .env int
	app.Listen(fmt.Sprintf("%s%d", ":", port))              // run the server
}
