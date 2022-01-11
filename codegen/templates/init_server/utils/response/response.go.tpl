package response

import (
	"github.com/gofiber/fiber/v2"
)

// idea / code taken from here https://github.com/fachryansyah/fotongo/blob/master/utils/response.go

func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": message,
		"data":    data,
	})
}

func ResponseNotFound(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(404).JSON(fiber.Map{
		"status":  404,
		"message": message,
	})
}

func ResponseError(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(500).JSON(fiber.Map{
		"status":  500,
		"message": message,
		"data":    data,
	})
}

func ResponseUnauthenticated(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(403).JSON(fiber.Map{
		"status":  403,
		"message": message,
		"data":    data,
	})
}

func ResponseValidationError(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(304).JSON(fiber.Map{
		"status":  304,
		"message": message,
		"data":    data,
	})
}
