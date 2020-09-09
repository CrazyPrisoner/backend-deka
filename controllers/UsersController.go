package controllers

import (
	"deka_backend/models"

	"github.com/gofiber/fiber"
)

func CreateUserController(c *fiber.Ctx) {
	c.Send(models.CreateUser())
}

func GetUserByIDController(c *fiber.Ctx) {
	c.Send(models.GetUserByID())
}

func GetAllUsersController(c *fiber.Ctx) {
	c.Send(models.GetAllUsers())
}

func DeleteUserController(c *fiber.Ctx) {
	c.Send(models.DeleteUser())
}
