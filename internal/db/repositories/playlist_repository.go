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
    return nil
}

// GetPlaylists returns all playlists from the database
func (r *PlaylistRepository) GetPlaylists() ([]*models.Playlist, error) {
    return nil, nil
}

// GetPlaylist returns a playlist from the database
func (r *PlaylistRepository) GetPlaylist(id int) (*models.Playlist, error) {
    return nil, nil
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
    return nil
}

// UpdatePlaylist updates a playlist in the database
func (r *PlaylistRepository) UpdatePlaylist(playlist *models.Playlist) error {
    return nil
}
