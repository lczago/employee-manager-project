package models

type SubjectModel struct {
	Description    string              `validate:"required"`
	FieldKnowledge FieldKnowledgeModel `validate:"required" json:"field_knowledge"`
}
