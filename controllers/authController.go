package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turkaytunc/go-web-fiber/database"
	"github.com/turkaytunc/go-web-fiber/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{FirstName: data["firstName"], LastName: data["lastName"], Email: data["email"], Password: hashedPassword}

	database.Connect().Create(&user)
	return c.Status(200).JSON(&user)
}
