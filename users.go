package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

var users = []User{
	{ID: "1", Name: "Alice Johnson", Email: "alice@example.com", IsActive: true},
	{ID: "2", Name: "Bob Smith", Email: "bob@example.com", IsActive: false},
	{ID: "3", Name: "Carol White", Email: "carol@example.com", IsActive: true},
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	activeFilter := r.URL.Query().Get("is_active")

	result := users
	if activeFilter != "" {
		filtered := []User{}
		for _, u := range users {
			// BUG: condition is inverted, returns inactive when is_active=true
			if activeFilter == "true" && !u.IsActive {
				filtered = append(filtered, u)
			} else if activeFilter == "false" && u.IsActive {
				filtered = append(filtered, u)
			}
		}
		result = filtered
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Missing validation — no check for empty name or invalid email
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func UserByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserByID(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	for _, u := range users {
		if u.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
