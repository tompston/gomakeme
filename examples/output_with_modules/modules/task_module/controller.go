package task_module

import (
	"output_with_modules/utils/response"
	"github.com/gofiber/fiber/v2"
)

var data = ""

func GetAllTasks(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "Tasks found!")
}

func GetTask(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "Task found!")
}

func CreateTask(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "Task created succesfully!")
}

func UpdateTask(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "Task updated succesfully!")
}

func DeleteTask(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, data, "Task deleted succesfully!")
}
