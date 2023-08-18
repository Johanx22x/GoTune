package repositories

import (
    "database/sql"
    "github.com/Johanx22x/GoTune/internal/db/models"
    "github.com/Johanx22x/GoTune/internal/utils"
)

// AlbumRepository is a struct that contains the database connection
type AlbumRepository struct {
    db *sql.DB
}

// NewAlbumRepository creates a new NewAlbumRepository
func NewAlbumRepository(db *sql.DB) *AlbumRepository {
    return &AlbumRepository{db}
}

// CreateAlbum creates a new album in the database
func (r *AlbumRepository) CreateAlbum(album *models.Album) error {
    return nil
}

// GetAlbums returns all albums from the database
func (r *AlbumRepository) GetAlbums() ([]*models.Album, error) {
    return nil, nil
}

// GetAlbum returns a album from the database
func (r *AlbumRepository) GetAlbum(id int) (*models.Album, error) {
    return nil, nil
}

// FilterAlbums returns all albums that match the predicate
func (r *AlbumRepository) FilterAlbums(predicate func(*models.Album) bool) ([]*models.Album, error) {
    albums, err := r.GetAlbums()
    if err != nil {
        return nil, err
    }
    return utils.Filter(albums, predicate), nil
}

// RemoveAlbum removes a album from the database
func (r *AlbumRepository) RemoveAlbum(album *models.Album) error {
    return nil
}

// UpdateAlbum updates a album in the database
func (r *AlbumRepository) UpdateAlbum(album *models.Album) error {
    return nil
}
