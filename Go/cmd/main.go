package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting...")

	router := chi.NewRouter()
	router.Get("/api/get", getHandler)
	router.Post("/api/post", postHandler)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Inside GET")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Post req")
}
