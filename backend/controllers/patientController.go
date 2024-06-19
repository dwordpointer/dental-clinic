package controllers

import (
	"net/http"

	"github.com/dwordpointer/dental-clinic-backend/database"
	"github.com/dwordpointer/dental-clinic-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetPatients(c *fiber.Ctx) error {
	var patients []models.Patient
	database.DB.Find(&patients)
	return c.JSON(patients)
}

func CreatePatient(c *fiber.Ctx) error {
	var patient models.Patient
	if err := c.BodyParser(&patient); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	database.DB.Create(&patient)
	return c.Status(http.StatusOK).JSON(patient)
}
