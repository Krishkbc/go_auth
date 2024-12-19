package handlers

import (
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}
