package controller

import "github.com/gofiber/fiber/v2"

func ShortenUrl(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Working fine"})
}
