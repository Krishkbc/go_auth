package repository

import (
	"errors"
	"go-auth/models"
)

var mockdb = []models.User{}

func CreateUser(user *models.User) error {
	for _, i := range mockdb {
		if i.Username == user.Username {
			return errors.New("username already exists")
		}
	}
	mockdb = append(mockdb, *user)
	return nil
}

func GetUserByUsername(username string) *models.User {
	return nil
}
