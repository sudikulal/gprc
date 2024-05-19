package models

import (
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "journal/db"
    "journal/config"
)

type FolderSchema struct {
    ID               primitive.ObjectID `bson:"_id,omitempty" json:"folderId,omitempty"`
    UserID           primitive.ObjectID `bson:"user_id" json:"userId"`
    CreatedAt        time.Time          `bson:"created_at" json:"createdAt"`
}

var FolderModel *mongo.Collection

func init() {
    client, err := db.GetMongoClient()
    if err != nil {
        log.Fatal(err)
    }
    FolderModel = client.Database(config.DATABASE).Collection("folder")
}
