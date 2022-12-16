package fk_repository

import (
	"go-api/database"
	"go-api/database/entities"
	"go-api/models"
	"go-api/utils"
)

type FieldKnowledgeRepository struct{}

func (repo *FieldKnowledgeRepository) GetAll() (*[]entities.FieldKnowledgeEntity, error) {
	var KnowledgeFields *[]entities.FieldKnowledgeEntity
	err := database.DB.Find(&KnowledgeFields).Error
	return KnowledgeFields, err
}

func (repo *FieldKnowledgeRepository) GetById(id string) (*entities.FieldKnowledgeEntity, error) {
	var entity *entities.FieldKnowledgeEntity
	err := database.DB.Find(&entity, id).Error
	return entity, err
}

func (repo *FieldKnowledgeRepository) GetAllSubjects(id string) (*[]entities.SubjectEntity, error) {
	var entity *entities.FieldKnowledgeEntity
	err := database.DB.Preload("Subjects").First(&entity, id).Error
	return &entity.Subjects, err
}

func (repo *FieldKnowledgeRepository) CreateFieldKnowledge(pModel models.FieldKnowledgeModel) (*entities.FieldKnowledgeEntity, error) {
	entity := utils.ParseObject[models.FieldKnowledgeModel, entities.FieldKnowledgeEntity](&pModel)
	err := database.DB.Create(&entity).Error
	return &entity, err
}
