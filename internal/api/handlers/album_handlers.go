package handlers

import (
    "net/http"
    // "encoding/json"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
)

// AlbumHandler handles all requests for albums
type AlbumHandler struct {
    albumRepository *repositories.AlbumRepository
}

// NewAlbumHandler creates a new AlbumHandler 
func NewAlbumHandler(albumRepository *repositories.AlbumRepository) *AlbumHandler {
    return &AlbumHandler{albumRepository}
}

// GetAlbums returns all albums 
func (handler *AlbumHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// GetAlbum returns a album by id 
func (handler *AlbumHandler) GetAlbum(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// CreateAlbum creates a new album
func (handler *AlbumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// FilterAlbums returns all albums that match the predicate 
func (handler *AlbumHandler) FilterAlbums(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// RemoveAlbum removes a album from the database 
func (handler *AlbumHandler) RemoveAlbum(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}

// UpdateAlbum updates a album in the database 
func (handler *AlbumHandler) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
