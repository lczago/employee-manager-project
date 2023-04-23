package entities

import (
	"go-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FieldKnowledge struct {
	Id                    primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	models.FieldKnowledge `bson:",inline"`
}
