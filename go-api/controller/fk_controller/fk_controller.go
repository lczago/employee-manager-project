package fk_controller

import (
	"github.com/gofiber/fiber/v2"
	"go-api/models"
	"go-api/service/fk_service"
	"go-api/utils"
)

var FKService fk_service.FieldKnowledge = new(fk_service.FieldKnowledgeService)

func GetAllKnowledgeFields(c *fiber.Ctx) error {
	KnowledgeFields, err := FKService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&KnowledgeFields)
}

func SearchFieldKnowledgeById(c *fiber.Ctx) error {
	id := c.Params("id")
	fieldKnowledge, err := FKService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func ListAllSubjects(c *fiber.Ctx) error {
	id := c.Params("id")
	subjects, err := FKService.GetAllSubjects(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&subjects)
}

func CreateFieldKnowledge(c *fiber.Ctx) error {
	var pModel models.FieldKnowledgeModel
	if err := c.BodyParser(&pModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err := utils.Validate.Struct(pModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	fieldKnowledge, err := FKService.CreateFieldKnowledge(pModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func EditFieldKnowledge(c *fiber.Ctx) error {
	id := c.Params("id")

	var pModel models.FieldKnowledgeModel
	if err := c.BodyParser(&pModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	fieldKnowledge, err := FKService.EditFieldKnowledge(pModel, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func DeleteFieldKnowledge(c *fiber.Ctx) error {
	id := c.Params("id")
	err := FKService.DeleteFieldKnowledge(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
