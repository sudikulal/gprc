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
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"todoId,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"userId" form:"userId"`
	Title     string             `bson:"title" json:"title" form:"title"`
	Status 	  int				 `bson:"status" json:"status" form:"status"`
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
