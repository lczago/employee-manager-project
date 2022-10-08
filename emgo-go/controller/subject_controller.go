package controller

import (
	"emgo-go/database"
	"emgo-go/dto"
	"emgo-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAllSubject(c *gin.Context) {
	var subject []models.Subject
	database.DB.Find(&subject)
	c.JSON(200, subject)
}

func SearchSubjectById(c *gin.Context) {
	var subject models.Subject
	id := c.Params.ByName("id")
	if err := database.DB.First(&subject, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, subject)
}

func CreateSubject(c *gin.Context) {
	var subjectForm dto.NewSubjectForm
	if err := c.ShouldBindJSON(&subjectForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var subject = subjectForm.MapForm()
	if err := database.DB.Create(&subject).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, subjectForm)
}

func EditSubject(c *gin.Context) {
	var subjectForm dto.NewSubjectForm
	var subject models.Subject
	id := c.Params.ByName("id")
	database.DB.First(&subject, id)
	if err := c.ShouldBindJSON(&subjectForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var subjectToUpdate = subjectForm.MapForm()
	if err := database.DB.Model(&subject).UpdateColumns(subjectToUpdate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, subjectForm)
}

func DeleteSubject(c *gin.Context) {
	var subject models.Subject
	id := c.Params.ByName("id")
	if err := database.DB.Delete(&subject, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"SUCCESS": "Field knowledge deleted successfully!"})
}
