package models

import (
	"log"
	"time"

	"journal/config"
	"journal/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FolderSchema struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"folderId" form:"folderId"`
	UserID     primitive.ObjectID `bson:"user_id" json:"userId" form:"userId"`
	FolderName string             `bson:"folder_name" json:"folderName"  form:"folderName"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt" form:"createdAt"`
}

var FolderModel *mongo.Collection

func init() {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	FolderModel = client.Database(config.DATABASE).Collection("folder")
}
