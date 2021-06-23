package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/turkaytunc/go-web-fiber/database"
	"github.com/turkaytunc/go-web-fiber/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	routes.Setup(app)

	app.Listen(":4000")
}
