package main

import (
	"log"

	"github.com/dwordpointer/dental-clinic-backend/controllers"
	"github.com/dwordpointer/dental-clinic-backend/database"
	"github.com/dwordpointer/dental-clinic-backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil{
        log.Fatal("Error loading .env files")
    }

    app := fiber.New()

    app.Use(logger.New())
    database.InitDatabase()

    app.Post("/login", controllers.Login)

    app.Use(middleware.JWTMiddleware)
    app.Get("/patients", controllers.GetPatients)
    app.Post("/patients", controllers.CreatePatient)

    app.Listen(":8080")
}
