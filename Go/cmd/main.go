package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting...")

	router := chi.NewRouter()
	router.Get("/api/get", getHandler)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Inside GET")
}
