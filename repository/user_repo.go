package repository

import (
	"database/sql"
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

func SaveUser(db *sql.DB, user *models.User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := db.Exec(query, user.Username, user.Password)
	return err
}
