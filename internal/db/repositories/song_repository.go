package repositories

import (
	"database/sql"
	"errors"

	"github.com/Johanx22x/GoTune/internal/db/models"
	"github.com/Johanx22x/GoTune/internal/utils"
)

// SongRepository is a struct that contains the database connection
type SongRepository struct {
    db *sql.DB 
}

// NewSongRepository creates a new SongRepository
func NewSongRepository(db *sql.DB) *SongRepository {
    return &SongRepository{db}
}

// GetSongs returns all songs from the database
func (r *SongRepository) GetSongs() ([]*models.Song, error) {
    songs := make([]*models.Song, 0)

    rows, err := r.db.Query("SELECT * FROM songs")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        song := new(models.Song)
        err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Genre, &song.Duration, &song.ReleaseDate)
        if err != nil {
            return nil, err
        }
        songs = append(songs, song)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return songs, nil
}

// CreateSong creates a new song in the database
func (r *SongRepository) CreateSong() error {
    return errors.New("not implemented")
}

// FilterSongs returns all songs that match the predicate
func (r *SongRepository) FilterSongs(predicate func(*models.Song) bool) ([]*models.Song, error) {
    songs, err := r.GetSongs()
    if err != nil {
        return nil, err
    }
    return utils.Filter(songs, predicate), nil
}

// RemoveSong removes a song from the database
func (r *SongRepository) RemoveSong(song *models.Song) error {
    return errors.New("not implemented")
}

// UpdateSong updates a song in the database
func (r *SongRepository) UpdateSong(song *models.Song) error {
    return errors.New("not implemented")
}
