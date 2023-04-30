package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading.env file")
	}
	ctx,cancel := context.WithTimeout(context.Background(),50*time.Second)
	uri := os.Getenv("MONGO_URI")
	// fmt.Println(uri)
	// fmt.Println("mongodb://127.0.0.1:27017")
	// "mongodb://127.0.0.1:27017"
	dbclient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	
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