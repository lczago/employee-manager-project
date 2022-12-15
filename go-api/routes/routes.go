package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api/controller"
)

func HandleRequests(api *fiber.App) {
	fieldKnowledgeRequests(api)
	subjectRequests(api)
}

func fieldKnowledgeRequests(api *fiber.App) {
	api.Get("/api/field_knowledge", controller.ListAllFieldKnowledge)
	api.Get("/api/field_knowledge/:id", controller.SearchFieldKnowledgeById)
	api.Get("/api/field_knowledge/:id/subjects", controller.ListAllSubjects)
	api.Post("/api/field_knowledge", controller.CreateFieldKnowledge)
	api.Put("/api/field_knowledge/:id", controller.EditFieldKnowledge)
	api.Delete("/api/field_knowledge", controller.DeleteFieldKnowledge)
}

func subjectRequests(api *fiber.App) {
	api.Get("/api/subjects", controller.ListAllSubject)
	api.Get("/api/subject/:id", controller.SearchSubjectById)
	api.Post("/api/subject", controller.CreateSubject)
	api.Put("/api/subject/:id", controller.EditSubject)
	api.Delete("/api/subject/:id", controller.DeleteSubject)
}
