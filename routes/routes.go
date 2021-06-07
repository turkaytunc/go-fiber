package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turkaytunc/go-web-fiber/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
