package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"url-shortner/routes"
)

func main() {
	//Initialised the app
	app := fiber.New()

	//Added CORS
	app.Use(cors.New())

	//App Routes
	routes.Routes(app)

	//Application listening to port
	port := ":3200"
	log.Fatal(app.Listen(port))
}
