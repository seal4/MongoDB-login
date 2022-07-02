package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Compare(username string) string {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://root:nqwk6au1Xs4caQVdhw3rplp5ch2PDCzfFSzj@cluster0.3abna.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("could not connect to MongoDB: ", err)
	}
	log.Println("Connected to MongoDB!")
	coll := client.Database("Users").Collection("Credentials") //innermost layer
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"Username", username}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return username
		}
		panic(err)
	}
	log.Println(result)
	return username
}
