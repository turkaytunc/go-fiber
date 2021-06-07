package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turkaytunc/go-web-fiber/database"
	"github.com/turkaytunc/go-web-fiber/routes"
)

func main() {

	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":4000")
}
