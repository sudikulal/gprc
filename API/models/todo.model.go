package models

import (
	"log"
	"time"

	"journal/config"
	"journal/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoSchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"journalId,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"userId"`
	Title     string             `bson:"title" json:"title"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

var TodoModel *mongo.Collection

func init() {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	TodoModel = client.Database(config.DATABASE).Collection("todo")
}
