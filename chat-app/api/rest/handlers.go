package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

var users = make(map[string]User)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Request payload", http.StatusBadRequest)
		return
	}
	users[user.Id] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqUser User
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user, exists := users[reqUser.Id]
	if !exists || user.Password != reqUser.Password {
		http.Error(w, "Invalid Username or password", http.StatusBadRequest)
		return
	}
	log.Printf("%s logged in successfully ", reqUser.Username)
	json.NewEncoder(w).Encode(user)
}
