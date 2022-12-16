package fk_service

import (
	"go-api/database/entities"
	"go-api/database/repository/fk_repository"
	"go-api/models"
)

type FieldKnowledge interface {
	GetAll() (*[]entities.FieldKnowledgeEntity, error)
	GetById(id string) (*entities.FieldKnowledgeEntity, error)
	GetAllSubjects(id string) (*[]entities.SubjectEntity, error)
	CreateFieldKnowledge(pModel models.FieldKnowledgeModel) (*entities.FieldKnowledgeEntity, error)
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
	return service.repo.CreateFieldKnowledge(pModel)
}
