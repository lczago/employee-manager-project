package controller

import (
	"emgo-go/database"
	"emgo-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAllFieldKnowledge(c *gin.Context) {
	var fieldKnowledge []models.FieldKnowledge
	database.DB.Find(&fieldKnowledge)
	c.JSON(200, fieldKnowledge)
}

func SearchFieldKnowledgeById(c *gin.Context) {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params.ByName("id")
	if err := database.DB.First(&fieldKnowledge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, fieldKnowledge)
}

func ListAllSubjects(c *gin.Context) {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params.ByName("id")
	if err := database.DB.Preload("Subjects.FieldKnowledge").First(&fieldKnowledge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, fieldKnowledge.Subjects)
}

func CreateFieldKnowledge(c *gin.Context) {
	var fieldKnowledge models.FieldKnowledge
	if err := c.ShouldBindJSON(&fieldKnowledge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := database.DB.Create(&fieldKnowledge).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, fieldKnowledge)
}

func EditFieldKnowledge(c *gin.Context) {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params.ByName("id")
	database.DB.First(&fieldKnowledge, id)
	if err := c.ShouldBindJSON(&fieldKnowledge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := database.DB.Model(&fieldKnowledge).UpdateColumns(fieldKnowledge).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, fieldKnowledge)
}

func DeleteFieldKnowledge(c *gin.Context) {
	var fieldKnowledge models.FieldKnowledge
	id := c.Params.ByName("id")
	if err := database.DB.Delete(&fieldKnowledge, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"SUCCESS": "Field knowledge deleted successfully!"})
}
