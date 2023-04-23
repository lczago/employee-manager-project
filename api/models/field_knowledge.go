package models

type FieldKnowledge struct {
	Description string `validate:"required" bson:"Description" json:"description"`
}
