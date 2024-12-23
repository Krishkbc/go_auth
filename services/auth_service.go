package services

import (
	"errors"
	"go-auth/models"
	"go-auth/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	// Check if username already exists
	exists, err := repository.IsUsernameTaken(user.Username)
	if err != nil {
		log.Println("Error checking username:", err) // Log the error
		return errors.New("database error: " + err.Error())
	}
	if exists {
		log.Println("Username already exists:", user.Username) // Log duplicate username
		return errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err) // Log hashing error
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)
	log.Println("Hashed Password:", string(hashedPassword))

	// Save the user to the database
	err = repository.SaveUser(user)
	if err != nil {
		log.Println("Error saving user to database:", err) // Log save error
		return err
	}

	return nil
}
