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
	r.GET("/field_knowledge", controller.ListAllFieldKnowledge)
	r.GET("/field_knowledge/:id", controller.SearchFieldKnowledgeById)
	r.GET("/field_knowledge/:id/subjects", controller.ListAllSubjects)
	r.POST("/field_knowledge", controller.CreateFieldKnowledge)
	r.PUT("/field_knowledge/:id", controller.EditFieldKnowledge)
	r.DELETE("/field_knowledge", controller.DeleteFieldKnowledge)
}

func subjectRequests(r *gin.Engine) {
	r.GET("/subject", controller.ListAllSubject)
	r.GET("/subject/:id", controller.SearchSubjectById)
	r.POST("/subject", controller.CreateSubject)
	r.PUT("/subject/:id", controller.EditSubject)
	r.DELETE("/subject/:id", controller.DeleteSubject)
}
