package main

import (
	"fmt"
	"go-auth/handlers"
	"go-auth/initializers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariable()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	// r.HandleFunc("/Login", handlers.Login).Methods("POST")

	fmt.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
