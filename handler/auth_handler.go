package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang2/model"
	"golang2/service"
)


// Handler untuk login
func LoginHandler(c *fiber.Ctx) error {
	var input model.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format JSON tidak valid",
		})
	}

	success, msg := service.Login(input)

	if success {
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    input.Username,
		})
	}

	return c.Status(401).JSON(fiber.Map{
		"message": msg,
	})
}


// Handler untuk register
func RegisterHandler(c *fiber.Ctx) error {
	var input model.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format JSON tidak valid",
		})
	}

	success, msg := service.Register(input)

	if success {
		return c.JSON(fiber.Map{
			"message": msg,
			"user":    input.Username,
		})
	}

	return c.Status(409).JSON(fiber.Map{
		"message": msg,
	})
}
