package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodf.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectMongo(uri string) *mongo.Client {
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fetal("MongoDB connection error:", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fetal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}