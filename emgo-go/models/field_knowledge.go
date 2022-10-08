package models

type FieldKnowledge struct {
	Id          uint      `gorm:"primaryKey" json:"_id"`
	Description string    `gorm:"100; unique; not null" json:"description"`
	Subjects    []Subject `json:"-"`
}
