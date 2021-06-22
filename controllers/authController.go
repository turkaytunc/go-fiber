package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turkaytunc/go-web-fiber/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{FirstName: "mahmut"}
	return c.JSON(user)
}
