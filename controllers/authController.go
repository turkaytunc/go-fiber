package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	user := models.User{FirstName: data["firstName"], LastName: data["lastName"], Email: data["email"], Password: string(hashedPassword)}

	database.DB.Create(&user)
	return c.Status(200).JSON(&user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user models.User

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{"message": "user not found"})
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if passErr != nil {
		return c.Status(400).JSON(fiber.Map{"message": "password is incorrect"})
	}

	cl := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{Issuer: strconv.Itoa(int(user.Id)), ExpiresAt: int64(time.Hour)})

	token, tokenErr := cl.SignedString([]byte("secret"))

	if err != nil {
		return tokenErr
	}

	c.Cookie(&fiber.Cookie{Name: "jwt", Value: token, Expires: time.Now().Add(time.Hour)})

	return c.JSON(&user)
}
