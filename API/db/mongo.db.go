package db

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"journal/config"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once

var clientInstanceError error

type Collection string

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(config.MONGO_URI)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		if err = client.Ping(ctx, nil); err != nil {
			log.Fatal(err)
		}

		clientInstance = client

		clientInstanceError = err
		log.Println("Connected to MongoDB!")
	})

	return clientInstance, clientInstanceError
}

// func init() {
// 	GetMongoClient()
// }
