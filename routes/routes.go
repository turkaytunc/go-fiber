package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turkaytunc/go-web-fiber/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
}
