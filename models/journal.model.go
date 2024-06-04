package models

import (
	"log"
	"time"

	"journal/config"
	"journal/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JournalSchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"journalId" form:"journalId"`
	UserID    primitive.ObjectID `bson:"user_id" json:"userId" form:"userId"`
	FolderID  string             `bson:"folder_id" json:"folderId" form:"folderId"`
	Title     string             `bson:"title" json:"title" form:"title"`
	DayRating int                `bson:"day_rating" json:"dayRating" form:"dayRating"`
	Content   string             `bson:"content" json:"Content" form:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt" form:"createdAt"`
}

var JournalModel *mongo.Collection

func init() {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	JournalModel = client.Database(config.DATABASE).Collection("journal")
}
