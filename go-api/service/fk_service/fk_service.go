package fk_service

import (
	"go-api/database/entities"
	"go-api/database/repository/fk_repository"
	"go-api/models"
	"go-api/utils"
)

type FieldKnowledge interface {
	GetAll() (*[]entities.FieldKnowledgeEntity, error)
	GetById(id string) (*entities.FieldKnowledgeEntity, error)
	GetAllSubjects(id string) (*[]entities.SubjectEntity, error)
	CreateFieldKnowledge(pModel models.FieldKnowledgeModel) (*entities.FieldKnowledgeEntity, error)
	EditFieldKnowledge(pModel models.FieldKnowledgeModel, id string) (*entities.FieldKnowledgeEntity, error)
	DeleteFieldKnowledge(id string) error
}

type FieldKnowledgeService struct {
	repo *fk_repository.FieldKnowledgeRepository
}

func (service *FieldKnowledgeService) GetAll() (*[]entities.FieldKnowledgeEntity, error) {
	return service.repo.GetAll()
}

func (service *FieldKnowledgeService) GetById(id string) (*entities.FieldKnowledgeEntity, error) {
	return service.repo.GetById(id)
}

func (service *FieldKnowledgeService) GetAllSubjects(id string) (*[]entities.SubjectEntity, error) {
	return service.repo.GetAllSubjects(id)
}

func (service *FieldKnowledgeService) CreateFieldKnowledge(pModel models.FieldKnowledgeModel) (*entities.FieldKnowledgeEntity, error) {
	entity := utils.ParseObject[models.FieldKnowledgeModel, entities.FieldKnowledgeEntity](&pModel)
	return service.repo.CreateFieldKnowledge(entity)
}

func (service *FieldKnowledgeService) EditFieldKnowledge(pModel models.FieldKnowledgeModel, id string) (*entities.FieldKnowledgeEntity, error) {
	entity := utils.ParseObject[models.FieldKnowledgeModel, *entities.FieldKnowledgeEntity](&pModel)
	return service.repo.EditFieldKnowledge(*entity, id)
}

func (service *FieldKnowledgeService) DeleteFieldKnowledge(id string) error {
	return service.repo.DeleteFieldKnowledge(id)
}
