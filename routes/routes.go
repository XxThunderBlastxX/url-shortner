package routes

import (
	"github.com/gofiber/fiber/v2"
	"url-shortner/controller"
)

func Routes(c *fiber.App) {
	c.Post("/", controller.ShortenUrl)
}
