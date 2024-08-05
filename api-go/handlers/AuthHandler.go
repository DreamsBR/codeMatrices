package handlers

import (
	"go-api/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthHandler(c *fiber.Ctx) error {
	token, err := utils.GenerateJWT()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot generate JWT"})
	}
	return c.JSON(fiber.Map{"token": token})
}
