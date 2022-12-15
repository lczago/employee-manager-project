package models

import "gorm.io/gorm"

type Subject struct {
	Id               uint           `gorm:"primaryKey" json:"_id"`
	Description      string         `gorm:"size:100; unique; not null"`
	FieldKnowledgeId uint           `json:"-"`
	FieldKnowledge   FieldKnowledge `json:"field_knowledge"`
}

func (subject *Subject) AfterFind(db *gorm.DB) error {
	return db.First(&subject.FieldKnowledge, subject.FieldKnowledgeId).Error
}
