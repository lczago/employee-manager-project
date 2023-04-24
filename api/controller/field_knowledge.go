package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-api/models"
	"go-api/service"
)

var fieldKnowledgeService = service.NewFieldKnowledgeService()

func GetAllKnowledgeFields(c *fiber.Ctx) error {
	KnowledgeFields, err := fieldKnowledgeService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&KnowledgeFields)
}

func SearchFieldKnowledgeById(c *fiber.Ctx) error {
	id := c.Params("id")
	fieldKnowledge, err := fieldKnowledgeService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func CreateFieldKnowledge(c *fiber.Ctx) error {
	var model models.FieldKnowledge
	if err := c.BodyParser(&model); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	fieldKnowledge, err := fieldKnowledgeService.CreateFieldKnowledge(model)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func EditFieldKnowledge(c *fiber.Ctx) error {
	id := c.Params("id")

	var pModel models.FieldKnowledge
	if err := c.BodyParser(&pModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	fieldKnowledge, err := fieldKnowledgeService.EditFieldKnowledge(pModel, id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fieldKnowledge)
}

func DeleteFieldKnowledge(c *fiber.Ctx) error {
	id := c.Params("id")
	err := fieldKnowledgeService.DeleteFieldKnowledge(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
