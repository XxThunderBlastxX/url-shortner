package routes

import (
	"url-shortner/controller"

	"github.com/gofiber/fiber/v2"
)

func Routes(c *fiber.App) {
	//Post req URl
	c.Post("/", controller.ShortenUrl)

	//Get req and redirect to orginal
	c.Get("/:urlId", controller.ResolveUrl)
}
