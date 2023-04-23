package service

import (
	"go-api/database/entities"
	"go-api/database/repository"
	"go-api/models"
	"go-api/utils"
)

type FieldKnowledgeService struct {
	repo repository.IFieldKnowledge
}

func NewFieldKnowledgeService() *FieldKnowledgeService {
	return &FieldKnowledgeService{
		repo: repository.NewFieldKnowledgeRepository(),
	}
}

func (service *FieldKnowledgeService) GetAll() (*[]entities.FieldKnowledge, error) {
	return service.repo.GetAll()
}

func (service *FieldKnowledgeService) GetById(id string) (*entities.FieldKnowledge, error) {
	return service.repo.GetById(id)
}

func (service *FieldKnowledgeService) CreateFieldKnowledge(model models.FieldKnowledge) (*entities.FieldKnowledge, error) {
	err := utils.Validate.Struct(model)
	if err != nil {
		return nil, err
	}

	entity := entities.FieldKnowledge{
		FieldKnowledge: model,
	}
	return service.repo.CreateFieldKnowledge(entity)
}

func (service *FieldKnowledgeService) EditFieldKnowledge(model models.FieldKnowledge, id string) (*entities.FieldKnowledge, error) {
	err := utils.Validate.Struct(model)
	if err != nil {
		return nil, err
	}
	return service.repo.EditFieldKnowledge(model, id)
}

func (service *FieldKnowledgeService) DeleteFieldKnowledge(id string) error {
	return service.repo.DeleteFieldKnowledge(id)
}
