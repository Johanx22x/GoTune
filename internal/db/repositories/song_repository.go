package repositories

import (
	"database/sql"

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

    rows, err := r.db.Query("SELECT * FROM song")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        song := new(models.Song)
        err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Duration, &song.ReleaseDate, &song.Genre)
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

// GetSong returns a song from the database by id 
func (r *SongRepository) GetSong(id string) (*models.Song, error) {
    song := new(models.Song)

    err := r.db.QueryRow("SELECT * FROM song WHERE id = $1", id).Scan(&song.ID, &song.Title, &song.Artist, &song.Duration, &song.ReleaseDate, &song.Genre)
    if err != nil {
        return nil, err
    }

    return song, nil
}

// CreateSong creates a new song in the database
func (r *SongRepository) CreateSong(song *models.Song) error {
    err := r.db.QueryRow(`INSERT INTO song (Title, Artist, Duration,
    ReleaseDate, Genre) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
    song.Title, song.Artist, song.Duration, song.ReleaseDate, song.Genre).Scan(&song.ID)
    if err != nil {
        return err
    }

    return nil
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
    _, err := r.db.Exec("DELETE FROM song WHERE id = $1", song.ID)
    if err != nil {
        return err
    }

    return nil
}

// UpdateSong updates a song in the database
func (r *SongRepository) UpdateSong(song *models.Song) error {
    err := r.db.QueryRow(`UPDATE song SET Title = $1, Artist = $2, Duration =
    $3, ReleaseDate = $4, Genre = $5 WHERE id = $6`, song.Title, song.Artist,
    song.Duration, song.ReleaseDate, song.Genre, song.ID).Scan()
    if err != nil {
        return err
    }

    return nil
}
