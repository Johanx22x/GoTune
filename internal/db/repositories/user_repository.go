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
    users := make([]*models.User, 0)

    rows, err := r.db.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }

    for rows.Next() {
        user := &models.User{}
        err := rows.Scan(&user.ID, &user.Username, &user.Password)
        if err != nil {
            return nil, err
        }

        users = append(users, user)
    }

    return users, nil
}

// GetUser gets a user from the database
func (r *UserRepository) GetUser(id string) (*models.User, error) {
    user := &models.User{}

    err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }

    return user, nil
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) error {
    err := r.db.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password).Scan(&user.ID)
    if err != nil {
        return err
    }

    return nil
}

// RemoveUser removes a user from the database
func (r *UserRepository) RemoveUser(user *models.User) error {
    _, err := r.db.Exec("DELETE FROM users WHERE id = $1", user.ID)
    if err != nil {
        return err
    }

    return nil
}

// UpdateUser updates a user in the database
func (r *UserRepository) UpdateUser(user *models.User) error {
    err := r.db.QueryRow("UPDATE users SET username = $1, password = $2 WHERE id = $3", user.Username, user.Password, user.ID).Scan()
    if err != nil {
        return err
    }

    return nil
}
