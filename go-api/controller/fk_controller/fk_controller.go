package fk_controller

import (
	"github.com/gofiber/fiber/v2"
	"go-api/database"
	"go-api/models"
	"go-api/service/fk_service"
	"go-api/utils"
)

var FKService fk_service.FieldKnowledge = new(fk_service.FieldKnowledgeService)

func GetAllKnowledgeFields(c *fiber.Ctx) error {
	KnowledgeFields, err := FKService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&err)
	}
	return c.Status(fiber.StatusOK).JSON(&KnowledgeFields)
}

func SearchFieldKnowledgeById(c *fiber.Ctx) error {
	id := c.Params("id")
	fieldKnowledge, err := FKService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func ListAllSubjects(c *fiber.Ctx) error {
	id := c.Params("id")
	subjects, err := FKService.GetAllSubjects(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&subjects)
}

func CreateFieldKnowledge(c *fiber.Ctx) error {
	var pModel models.FieldKnowledgeModel
	if err := c.BodyParser(&pModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	err := utils.Validate.Struct(pModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	fieldKnowledge, err := FKService.CreateFieldKnowledge(pModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func EditFieldKnowledge(c *fiber.Ctx) error {
	var fieldKnowledge models.FieldKnowledgeModel
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
	var fieldKnowledge models.FieldKnowledgeModel
	id := c.Params("id")
	if err := database.DB.Delete(&fieldKnowledge, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}
