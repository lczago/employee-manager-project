package dto

import (
	"emgo-go/models"
)

type NewSubjectForm struct {
	Description    string `json:"description" validate:"required"`
	FieldKnowledge uint   `json:"field_knowledge" validate:"required"`
}

func (subjectForm NewSubjectForm) MapForm() models.Subject {
	var subject models.Subject
	subject.Description = subjectForm.Description
	subject.FieldKnowledgeId = subjectForm.FieldKnowledge
	return subject
}
