package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
)

// SongHandler handles all requests for songs
type SongHandler struct {
    songRepository *repositories.SongRepository
}

// NewSongHandler creates a new SongHandler 
func NewSongHandler(songRepository *repositories.SongRepository) *SongHandler {
    return &SongHandler{songRepository}
}

// GetSongs returns all songs 
func (handler *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {
    songs, err := handler.songRepository.GetSongs()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    songsJSON, err := json.Marshal(songs)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return 
    }

    // Write songs to response 
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(songsJSON)
}

// GetSong returns a song by id 
func (handler *SongHandler) GetSong(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// CreateSong creates a new song 
func (handler *SongHandler) CreateSong(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// FilterSongs returns all songs that match the predicate
func (handler *SongHandler) FilterSongs(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// RemoveSong removes a song from the database
func (handler *SongHandler) RemoveSong(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// UpdateSong updates a song in the database
func (handler *SongHandler) UpdateSong(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
