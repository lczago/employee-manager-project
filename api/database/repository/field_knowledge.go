package repository

import (
	"context"
	"errors"
	"go-api/database"
	"go-api/database/entities"
	"go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IFieldKnowledge interface {
	GetAll() (*[]entities.FieldKnowledge, error)
	GetById(id string) (*entities.FieldKnowledge, error)
	CreateFieldKnowledge(entity entities.FieldKnowledge) (*entities.FieldKnowledge, error)
	EditFieldKnowledge(model models.FieldKnowledge, id string) (*entities.FieldKnowledge, error)
	DeleteFieldKnowledge(id string) error
}

type FieldKnowledgeRepository struct {
	CollectionName string
}

func NewFieldKnowledgeRepository() IFieldKnowledge {
	return &FieldKnowledgeRepository{
		CollectionName: "FieldKnowledge",
	}
}

func (repo *FieldKnowledgeRepository) GetAll() (*[]entities.FieldKnowledge, error) {
	coll := database.GetCollection(repo.CollectionName)

	cursor, errFind := coll.Find(context.TODO(), bson.M{})
	if errFind != nil {
		return nil, errFind
	}

	var KnowledgeFields []entities.FieldKnowledge
	if errDecode := cursor.All(context.TODO(), &KnowledgeFields); errDecode != nil {
		return nil, errDecode
	}

	return &KnowledgeFields, nil
}

func (repo *FieldKnowledgeRepository) GetById(id string) (*entities.FieldKnowledge, error) {
	coll := database.GetCollection(repo.CollectionName)

	filter := bson.M{
		"_id": id,
	}

	var entity *entities.FieldKnowledge
	if err := coll.FindOne(context.TODO(), filter).Decode(&entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (repo *FieldKnowledgeRepository) CreateFieldKnowledge(entity entities.FieldKnowledge) (*entities.FieldKnowledge, error) {
	coll := database.GetCollection(repo.CollectionName)

	filter := bson.M{"Description": entity.Description}
	opts := options.Count().SetCollation(&options.Collation{Locale: "pt", Strength: 1})
	count, err := coll.CountDocuments(context.TODO(), filter, opts)
	if err != nil || count > 0 {
		return nil, errors.New("field Knowledge already exists")
	}

	result, err := coll.InsertOne(context.TODO(), entity)
	if err != nil {
		return nil, err
	}

	entity.Id = result.InsertedID.(primitive.ObjectID)

	return &entity, nil
}

func (repo *FieldKnowledgeRepository) EditFieldKnowledge(model models.FieldKnowledge, id string) (*entities.FieldKnowledge, error) {
	coll := database.GetCollection(repo.CollectionName)

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	update := bson.M{"$set": model}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	opts.SetCollation(&options.Collation{Locale: "pt", Strength: 1})

	var entity entities.FieldKnowledge
	if err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (repo *FieldKnowledgeRepository) DeleteFieldKnowledge(id string) error {
	coll := database.GetCollection(repo.CollectionName)

	filter := bson.M{"_id": id}

	if _, err := coll.DeleteOne(context.TODO(), filter); err != nil {
		return err
	}

	return nil
}
