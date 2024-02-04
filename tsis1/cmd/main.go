package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tsis1/internal/handler"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/players", handler.GetPlayers).Methods("GET")
	router.HandleFunc("/players/{name}", handler.GetPlayerByName).Methods("GET")
	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}g
