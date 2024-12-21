package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/models"
	"go-auth/services"
	"net/http"
)

type Response struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.RegisterUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := Response{
		Message:  "User created successfully",
		Username: user.Username,
	}

	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}
