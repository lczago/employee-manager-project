package models

type SubjectModel struct {
	Description    string `validate:"required" json:"description"`
	FieldKnowledge string `validate:"required" json:"field_knowledge"`
}
