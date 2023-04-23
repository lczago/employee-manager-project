package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func ConnectDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		AuthSource: "",
		Username:   "root",
		Password:   "admin",
	}))
	if err != nil {
		fmt.Println("deu erro ao conectar no mongo")
	}

	db = client
}

func CloseMongo() {
	db.Disconnect(context.TODO())
}

func GetCollection(collectionName string) *mongo.Collection {
	return db.Database("Emgo").Collection(collectionName)
}
