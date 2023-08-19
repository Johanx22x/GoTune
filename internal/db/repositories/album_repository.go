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
    err := r.db.QueryRow(`INSERT INTO album (Title, Artist, ReleaseDate, Genre, ImagePath) VALUES ($1, $2, $3, $4, $5)
    RETURNING id`, album.Title, album.Artist, album.ReleaseDate, album.Genre, album.ImagePath).Scan(&album.ID)
    if err != nil {
        return err
    }

    return nil
}

// GetAlbums returns all albums from the database
func (r *AlbumRepository) GetAlbums() ([]*models.Album, error) {
    rows, err := r.db.Query(`SELECT * FROM album`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    albums := make([]*models.Album, 0)

    for rows.Next() {
        album := &models.Album{}
        err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.ReleaseDate, &album.Genre, &album.ImagePath)
        if err != nil {
            return nil, err
        }
        albums = append(albums, album)
    }

    return albums, nil
}

// GetAlbum returns a album from the database
func (r *AlbumRepository) GetAlbum(id string) (*models.Album, error) {
    album := &models.Album{}
    err := r.db.QueryRow(`SELECT * FROM album WHERE id = $1`,
    id).Scan(&album.ID, &album.Title, &album.Artist, &album.ReleaseDate,
    &album.Genre, &album.ImagePath)
    if err != nil {
        return nil, err
    }
    return album, nil
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
    _, err := r.db.Exec(`DELETE FROM album WHERE id = $1`, album.ID)
    if err != nil {
        return err
    }

    return nil
}

// UpdateAlbum updates a album in the database
func (r *AlbumRepository) UpdateAlbum(album *models.Album) error {
    err := r.db.QueryRow(`UPDATE album SET Title = $1, Artist = $2, ReleaseDate = $3, Genre = $4, ImagePath = $5 WHERE id = $6
    RETURNING id`, album.Title, album.Artist, album.ReleaseDate, album.Genre, album.ImagePath, album.ID).Scan(&album.ID)
    if err != nil {
        return err
    }

    return nil
}
