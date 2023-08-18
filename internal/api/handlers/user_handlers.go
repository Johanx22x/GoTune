package handlers

import (
    "net/http"
    // "encoding/json"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
)

// UserHandler handles all requests for users 
type UserHandler struct {
    userRepository *repositories.UserRepository
}

// NewUserHandler creates a new UserHandler 
func NewUserHandler(userRepository *repositories.UserRepository) *UserHandler {
    return &UserHandler{userRepository}
}

// GetUsers returns all users 
func (handler *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) { 
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// GetUser returns a user by id 
func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// CreateUser creates a new user 
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// FilterUsers returns all users that match the predicate 
func (handler *UserHandler) FilterUsers(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// RemoveUser removes a user from the database 
func (handler *UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// UpdateUser updates a user in the database 
func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
