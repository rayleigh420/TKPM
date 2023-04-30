package database

import (
	// "os"b
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	ctx,cancel := context.WithTimeout(context.Background(),50*time.Second)
	dbclient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = dbclient.Connect(ctx); err != nil {
		defer cancel()
		log.Fatal(err.Error())
	}
	defer cancel()
	return dbclient
}

var Client = DBInstance()

func OpenCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	collection := client.Database("go-lib").Collection(collectionName)
	return collection
}