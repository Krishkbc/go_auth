package initializers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var UserCollection *mongo.Collection

func ConnectDatabase() {
	uri := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB successfully")

	// Set global variables
	MongoClient = client
	UserCollection = client.Database("auth_app").Collection("users")

}

func CreateUserIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"username": 1}, // Create index on the username field
		Options: options.Index().SetUnique(true),
	}

	_, err := UserCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("Error creating index:", err)
	}
	log.Println("Created index for username")
}
