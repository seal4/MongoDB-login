package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func upload(username string, password string) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("MongoDB credentials go here").SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("could not connect to MongoDB: ", err)
	}
	log.Println("Connected to MongoDB!")
	Site := client.Database("Users")                  //outermost layer
	LoginCollection := Site.Collection("Credentials") //innermost layer

	result, err := LoginCollection.InsertOne(ctx, bson.D{
		{Key: "Username", Value: username},
		{Key: "Password", Value: hash(password)},
	})
	if err != nil {
		log.Fatal("error uploading data to MongoDB: ", err)
	}
	log.Println("uploaded sucessfully!: ", result)
}
