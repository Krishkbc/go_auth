package repository

import (
	"context"
	"errors"
	"go-auth/initializers"
	"go-auth/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := initializers.UserCollection.InsertOne(ctx, user)
	if mongo.IsDuplicateKeyError(err) {
		log.Println("Duplicate key error for username:", user.Username) // Log duplicate error
		return errors.New("username already exists")
	}

	if err != nil {
		log.Println("Error inserting user:", err) // Log insert error
		return err
	}

	return nil
}

func IsUsernameTaken(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	count, err := initializers.UserCollection.CountDocuments(ctx, filter)
	return count > 0, err
}
func CreateUser(user *models.User) error {

	// Implementation for creating a user in the repository

	// This is a placeholder implementation

	return nil

}
