package main

import (
	"log"
	"os"
	"url-shortner/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	//Load env variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//Initialised the app
	app := fiber.New()

	//Added CORS
	app.Use(cors.New())

	//App Routes
	routes.Routes(app)

	//Application listening to port
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
