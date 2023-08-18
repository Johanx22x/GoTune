package handlers

import (
    "net/http"
    // "encoding/json"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
)

// PlaylistHandler handles all requests for playlists
type PlaylistHandler struct {
    playlistRepository *repositories.PlaylistRepository
}

// NewPlaylistHandler creates a new PlaylistHandler 
func NewPlaylistHandler(playlistRepository *repositories.PlaylistRepository) *PlaylistHandler {
    return &PlaylistHandler{playlistRepository}
}

// GetPlaylists returns all playlists 
func (handler *PlaylistHandler) GetPlaylists(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// GetPlaylist returns a playlist by id 
func (handler *PlaylistHandler) GetPlaylist(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// CreatePlaylist creates a new playlist 
func (handler *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// FilterPlaylists returns all playlists that match the predicate
func (handler *PlaylistHandler) FilterPlaylists(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// RemovePlaylist removes a playlist from the database
func (handler *PlaylistHandler) RemovePlaylist(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// UpdatePlaylist updates a playlist in the database 
func (handler *PlaylistHandler) UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
