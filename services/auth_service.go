package services

import (
	"errors"
	"go-auth/models"
	"go-auth/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("Username and Password are required")
	}

	bcrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(bcrypt)

	if err := repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}
