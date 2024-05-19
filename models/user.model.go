package models

import (
	"log"
	"time"

	"journal/config"
	"journal/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSchema struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"userId,omitempty"`
	EmailId      string             `bson:"email_id" json:"emailId"`
	UserName     string             `bson:"user_name" json:"userName"`
	Password     string             `bson:"password" json:"password"`
	RegisterType int                `bson:"register_type" json:"registerType"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
}

var UserModel *mongo.Collection

func init() {
	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	UserModel = client.Database(config.DATABASE).Collection("user")
}
