package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://nitisarath:p5U424lf6bmlFCBw@cluster0.m01vueg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// ------------------ NEW CODE TO QUERY THE DATABASE ------------------
	// Select the "User_1" database and the "User" collection
	database := client.Database("User_1")
	collection := database.Collection("User")

	// Define a filter to find a specific document.
	// This example finds a document where the "name" field is "Alice".
	filter := bson.M{"name": "Name"}

	// Create a variable to hold the query result
	var result bson.M

	// Find a single document that matches the filter
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document was found with the specified filter.")
			return
		}
		log.Fatal(err)
	}

	fmt.Println("Found a single document:")
	fmt.Println(result)
}
