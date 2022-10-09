package routes

import (
	"emgo-go/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	fieldKnowledgeRequests(r)
	subjectRequests(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err.Error())
		return
	}
}

func fieldKnowledgeRequests(r *gin.Engine) {
	r.GET("/api/field_knowledge", controller.ListAllFieldKnowledge)
	r.GET("/api/field_knowledge/:id", controller.SearchFieldKnowledgeById)
	r.GET("/api/field_knowledge/:id/subjects", controller.ListAllSubjects)
	r.POST("/api/field_knowledge", controller.CreateFieldKnowledge)
	r.PUT("/api/field_knowledge/:id", controller.EditFieldKnowledge)
	r.DELETE("/api/field_knowledge", controller.DeleteFieldKnowledge)
}

func subjectRequests(r *gin.Engine) {
	r.GET("/api/subjects", controller.ListAllSubject)
	r.GET("/api/subject/:id", controller.SearchSubjectById)
	r.POST("/api/subject", controller.CreateSubject)
	r.PUT("/api/subject/:id", controller.EditSubject)
	r.DELETE("/api/subject/:id", controller.DeleteSubject)
}
