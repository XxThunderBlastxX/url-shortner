package routes

import (
	"github.com/gofiber/fiber/v2"
	"url-shortner/controller"
)

func Routes(c *fiber.App) {
	//Post req URl
	c.Post("/", controller.ShortenUrl)

	//Get req and redirect to orginal
	c.Get("/:urlId", controller.ResolveUrl)
}
