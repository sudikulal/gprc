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
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"journalId,omitempty"`
	UserID           primitive.ObjectID `bson:"user_id" json:"userId"`
	FolderID         primitive.ObjectID `bson:"folder_id" json:"folderId"`
	Title            string             `bson:"title" json:"title"`
	EncryptedContent string             `bson:"encrypted_content" json:"encryptedContent"`
	CreatedAt        time.Time          `bson:"created_at" json:"createdAt"`
}

var JournalModel *mongo.Collection

func init() {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	JournalModel = client.Database(config.DATABASE).Collection("journal")
}
