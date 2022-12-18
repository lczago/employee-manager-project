package fk_repository

import (
	"go-api/database"
	"go-api/database/entities"
	"gorm.io/gorm/clause"
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

func (repo *FieldKnowledgeRepository) CreateFieldKnowledge(entity entities.FieldKnowledgeEntity) (*entities.FieldKnowledgeEntity, error) {
	err := database.DB.Create(&entity).Error
	return &entity, err
}

func (repo *FieldKnowledgeRepository) EditFieldKnowledge(entity entities.FieldKnowledgeEntity, id string) (*entities.FieldKnowledgeEntity, error) {
	err := database.DB.Model(&entity).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&entity).Error
	return &entity, err
}

func (repo *FieldKnowledgeRepository) DeleteFieldKnowledge(id string) error {
	var fieldKnowledge *entities.FieldKnowledgeEntity
	err := database.DB.Delete(fieldKnowledge, id).Error
	return err
}
