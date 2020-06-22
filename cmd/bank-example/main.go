package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"bank-example/internal/app"
	"bank-example/internal/controllers"
	"bank-example/internal/models"
)

// entry point into bank
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users", controllers.UpdateUser).Methods("POST")
	router.HandleFunc("/api/v1/users", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users/balance", controllers.GetBalance).Methods("GET")
	router.HandleFunc("/api/v1/users/authorize", controllers.AuthorizeTransaction).Methods("POST")
	router.HandleFunc("/health", controllers.Health).Methods("GET")
	router.Use(app.Authentication)

	models.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Print(err)
	}
}
