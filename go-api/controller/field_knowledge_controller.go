package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-api/database"
	"go-api/models"
)

func ListAllFieldKnowledge(c *fiber.Ctx) error {
	var fieldKnowledge []models.FieldKnowledge
	database.DB.Find(&fieldKnowledge)
	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func SearchFieldKnowledgeById(c *fiber.Ctx) error {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params("id")
	if err := database.DB.First(&fieldKnowledge, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func ListAllSubjects(c *fiber.Ctx) error {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params("id")
	if err := database.DB.Preload("Subjects").First(&fieldKnowledge, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge.Subjects)
}

func CreateFieldKnowledge(c *fiber.Ctx) error {
	var fieldKnowledge models.FieldKnowledge
	if err := c.BodyParser(&fieldKnowledge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := database.DB.Create(&fieldKnowledge).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func EditFieldKnowledge(c *fiber.Ctx) error {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params("id")
	database.DB.First(&fieldKnowledge, id)
	if err := c.BodyParser(&fieldKnowledge); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := database.DB.Model(&fieldKnowledge).UpdateColumns(fieldKnowledge).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func DeleteFieldKnowledge(c *fiber.Ctx) error {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params("id")
	if err := database.DB.Delete(&fieldKnowledge, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}
