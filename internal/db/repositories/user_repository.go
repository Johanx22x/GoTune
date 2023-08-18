package repositories

import (
    "database/sql"
    "github.com/Johanx22x/GoTune/internal/db/models"
)

// UserRepository is a struct that contains the database connection
type UserRepository struct {
    db *sql.DB
}

// NewUserRepository creates a new UserRepository 
func NewUserRepository(db *sql.DB) *UserRepository { 
    return &UserRepository{db}
}

// GetUsers gets all users from the database
func (r *UserRepository) GetUsers() ([]*models.User, error) {
    return nil, nil
}

// GetUser gets a user from the database
func (r *UserRepository) GetUser(id int) (*models.User, error) {
    return nil, nil
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) error {
    return nil
}

// RemoveUser removes a user from the database
func (r *UserRepository) RemoveUser(user *models.User) error {
    return nil
}

// UpdateUser updates a user in the database
func (r *UserRepository) UpdateUser(user *models.User) error {
    return nil
}
