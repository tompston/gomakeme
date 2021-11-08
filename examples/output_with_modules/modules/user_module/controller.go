package user_module

import (
	"output_with_modules/utils/response"
	"github.com/gofiber/fiber/v2"
)

var data = ""

func GetAllUsers(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "Users found!")
}

func GetUser(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "User found!")
}

func CreateUser(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "User created succesfully!")
}

func UpdateUser(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "User updated succesfully!")
}

func DeleteUser(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "User deleted succesfully!")
}
