package services

import (
	"errors"
	"go-auth/models"
)

func RegisterUser(user *models.User) error {
	if user.username == "" || user.password == "" {
		return errors.New("Username and Password are required")
	}

}
