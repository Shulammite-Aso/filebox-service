package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	Collection *mongo.Collection
}

func Init(url string) Handler {

	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Connecting to MongoDB...")
	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("Error Starting MongoDB Client: %v", err)
	}

	collection := client.Database("fileboxDB").Collection("file")

	return Handler{collection}
}
