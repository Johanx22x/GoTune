package handlers

import (
    "encoding/json"

    "github.com/Johanx22x/GoTune/internal/db/repositories"
    "github.com/Johanx22x/GoTune/internal/db/models"
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
func (handler *UserHandler) GetUsers() (string, error) {
    users, err := handler.userRepository.GetUsers()
    if err != nil {
        return "", err
    }

    usersJSON, err := json.Marshal(users)
    if err != nil {
        return "", err
    }

    return string(usersJSON), nil
}

// GetUser returns a user by id 
func (handler *UserHandler) GetUser(id string) (string, error) {
    user, err := handler.userRepository.GetUser(id)
    if err != nil {
        return "", err
    }

    userJSON, err := json.Marshal(user)
    if err != nil {
        return "", err
    }

    return string(userJSON), nil
}

// CreateUser creates a new user 
func (handler *UserHandler) CreateUser(username, password string) (string, error) {
    user := &models.User{Username: username, Password: password}

    err := handler.userRepository.CreateUser(user)
    if err != nil {
        return "", err
    }

    userJSON, err := json.Marshal(user)
    if err != nil {
        return "", err
    }

    return string(userJSON), nil
}

// RemoveUser removes a user from the database 
func (handler *UserHandler) RemoveUser(id string) (string, error) {
    user, err := handler.userRepository.GetUser(id)
    if err != nil {
        return "", err
    }

    err = handler.userRepository.RemoveUser(user)
    if err != nil {
        return "", err
    }

    userJSON, err := json.Marshal(user)
    if err != nil {
        return "", err
    }

    return string(userJSON), nil
}

// UpdateUser updates a user in the database 
func (handler *UserHandler) UpdateUser(id, username, password string) (string, error) {
    user, err := handler.userRepository.GetUser(id)
    if err != nil {
        return "", err
    }

    user.Username = username
    user.Password = password

    err = handler.userRepository.UpdateUser(user)
    if err != nil {
        return "", err
    }

    userJSON, err := json.Marshal(user)
    if err != nil {
        return "", err
    }

    return string(userJSON), nil
}
