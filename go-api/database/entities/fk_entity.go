package entities

type FieldKnowledgeEntity struct {
	Id          uint            `gorm:"primaryKey" json:"_id"`
	Description string          `gorm:"100; unique; not null"`
	Subjects    []SubjectEntity `gorm:"foreignKey:FieldKnowledgeId" json:"-"`
}
