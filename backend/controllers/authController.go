package controllers

import (
	"os"
	"time"
	"github.com/dwordpointer/dental-clinic-backend/database"
	"github.com/dwordpointer/dental-clinic-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Register(c *fiber.Ctx) error {
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

    user := models.User{
        Username: data["username"],
        Password: string(password),
    }

    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "could not create user"})
    }

    return c.JSON(fiber.Map{"message": "user created successfully"})
}

func Login(c *fiber.Ctx) error {
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    var user models.User
    database.DB.Where("username = ?", data["username"]).First(&user)

    if user.ID == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect password"})
    }

    claims := jwt.MapClaims{
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString(jwtSecret)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
    }

    return c.JSON(fiber.Map{"token": t})
}
