package http


import (
	// "encoding/json"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"
	// "github.com/mateusgms/golang-blog/models"
	// "github.com/mateusgms/golang-blog/services"
)

// CreateUserHandler is a handler for creating a user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

// ListUserHandler is a handler for listing all users
func ListUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// GetUserHandler is a handler for getting a user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// UpdateUserHandler is a handler for updating a user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// DeleteUserHandler is a handler for deleting a user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
