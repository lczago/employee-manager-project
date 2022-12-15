package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-api/database"
	"go-api/dto"
	"go-api/models"
)

func ListAllSubject(c *fiber.Ctx) error {
	var subject []models.Subject
	database.DB.Find(&subject)
	return c.Status(fiber.StatusOK).JSON(&subject)
}

func SearchSubjectById(c *fiber.Ctx) error {
	var subject models.Subject
	id := c.Params("id")
	if err := database.DB.First(&subject, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(&subject)
}

func CreateSubject(c *fiber.Ctx) error {
	var subjectForm dto.NewSubjectForm
	if err := c.BodyParser(&subjectForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	var subject = subjectForm.MapForm()
	if err := database.DB.Create(&subject).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(&subjectForm)
}

func EditSubject(c *fiber.Ctx) error {
	var subjectForm dto.NewSubjectForm
	var subject models.Subject
	id := c.Params("id")
	database.DB.First(&subject, id)
	if err := c.BodyParser(&subjectForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	var subjectToUpdate = subjectForm.MapForm()
	if err := database.DB.Model(&subject).UpdateColumns(subjectToUpdate).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(&subjectForm)
}

func DeleteSubject(c *fiber.Ctx) error {
	var subject models.Subject
	id := c.Params("id")
	if err := database.DB.Delete(&subject, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(&subject)
}
