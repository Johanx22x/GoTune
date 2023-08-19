package repositories

import (
    "database/sql"
    "github.com/Johanx22x/GoTune/internal/db/models"
    "github.com/Johanx22x/GoTune/internal/utils"
)

// PlaylistRepository is a struct that contains the database connection
type PlaylistRepository struct {
    db *sql.DB
}

// NewPlaylistRepository creates a new PlaylistRepository 
func NewPlaylistRepository(db *sql.DB) *PlaylistRepository {
    return &PlaylistRepository{db}
}

// CreatePlaylist creates a new playlist in the database
func (r *PlaylistRepository) CreatePlaylist(playlist *models.Playlist) error {
    err := r.db.QueryRow(`INSERT INTO playlist (Title, Description, CreatorID, CreatedAt, ImagePath) VALUES ($1, $2, $3, $4, $5)
    RETURNING id`, playlist.Title, playlist.Description, playlist.CreatorID, playlist.CreatedAt, playlist.ImagePath).Scan(&playlist.ID)
    if err != nil {
        return err
    }

    return nil
}

// GetPlaylists returns all playlists from the database
func (r *PlaylistRepository) GetPlaylists() ([]*models.Playlist, error) {
    rows, err := r.db.Query(`SELECT * FROM playlist`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    playlists := make([]*models.Playlist, 0)

    for rows.Next() {
        playlist := &models.Playlist{}
        err := rows.Scan(&playlist.ID, &playlist.Title, &playlist.Description, &playlist.CreatorID, &playlist.CreatedAt, &playlist.ImagePath)
        if err != nil {
            return nil, err
        }
        playlists = append(playlists, playlist)
    }

    return playlists, nil
}

// GetPlaylist returns a playlist from the database
func (r *PlaylistRepository) GetPlaylist(id string) (*models.Playlist, error) {
    playlist := &models.Playlist{}
    err := r.db.QueryRow(`SELECT * FROM playlist WHERE id = $1`, id).Scan(&playlist.ID, &playlist.Title, &playlist.Description, &playlist.CreatorID, &playlist.CreatedAt, &playlist.ImagePath)
    if err != nil {
        return nil, err
    }
    return playlist, nil
}

// FilterPlaylists returns all playlists that match the predicate
func (r *PlaylistRepository) FilterPlaylists(predicate func(*models.Playlist) bool) ([]*models.Playlist, error) {
    playlists, err := r.GetPlaylists()
    if err != nil {
        return nil, err
    }
    return utils.Filter(playlists, predicate), nil
}

// RemovePlaylist removes a playlist from the database
func (r *PlaylistRepository) RemovePlaylist(playlist *models.Playlist) error {
    _, err := r.db.Exec(`DELETE FROM playlist WHERE id = $1`, playlist.ID)
    if err != nil {
        return err
    }

    return nil
}

// UpdatePlaylist updates a playlist in the database
func (r *PlaylistRepository) UpdatePlaylist(playlist *models.Playlist) error {
    err := r.db.QueryRow(`UPDATE playlist SET Title = $1, Description = $2,
    CreatorID = $3, CreatedAt = $4, ImagePath = $5 WHERE id = $6 RETURNING id`,
    playlist.Title, playlist.Description, playlist.CreatorID, playlist.CreatedAt, playlist.ImagePath, playlist.ID).Scan(&playlist.ID)
    if err != nil {
        return err
    }

    return nil
}
