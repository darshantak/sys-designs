package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/create-form", CreateForm).Methods("POST")
	r.HandleFunc("/forms", ListForms).Methods("GET")
	r.HandleFunc("/submit", SubmitForm).Methods("POST")
	r.HandleFunc("/forms/{form_id}", GetFormResponses).Methods("GET")

	// Start server
	log.Println("Server running on http://localhost:7070")
	log.Fatal(http.ListenAndServe(":7070", r))
}
