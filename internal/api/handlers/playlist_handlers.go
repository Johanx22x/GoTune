package handlers

import (
	"encoding/json"

	"github.com/Johanx22x/GoTune/internal/db/models"
	"github.com/Johanx22x/GoTune/internal/db/repositories"
	"github.com/Johanx22x/GoTune/internal/utils"
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
func (handler *PlaylistHandler) GetPlaylists() (string, error) {
    playlists, err := handler.playlistRepository.GetPlaylists()
    if err != nil {
        return "", err
    }

    playlistsJSON, err := json.Marshal(playlists)
    if err != nil {
        return "", err
    }

    return string(playlistsJSON), nil
}

// GetPlaylist returns a playlist by id 
func (handler *PlaylistHandler) GetPlaylist(id string) (string, error) {
    playlist, err := handler.playlistRepository.GetPlaylist(id)
    if err != nil {
        return "", err
    }

    playlistJSON, err := json.Marshal(playlist)
    if err != nil {
        return "", err
    }

    return string(playlistJSON), nil
}

// CreatePlaylist creates a new playlist 
func (handler *PlaylistHandler) CreatePlaylist(title, description, creatorIDString, createdAtString, imagePath string) (string, error) {
    creatorID, err := utils.StringToInt(creatorIDString)
    if err != nil {
        return "", err
    }

    createdAt, err := utils.StringToInt64(createdAtString)
    if err != nil {
        return "", err
    }

    playlist := &models.Playlist{
        Title: title,
        Description: description,
        CreatorID: creatorID,
        CreatedAt: createdAt,
        ImagePath: imagePath,
    }

    err = handler.playlistRepository.CreatePlaylist(playlist)
    if err != nil {
        return "", err
    }

    playlistJSON, err := json.Marshal(playlist)
    if err != nil {
        return "", err
    }

    return string(playlistJSON), nil
}

// FilterPlaylists returns all playlists that match the predicate
func (handler *PlaylistHandler) FilterPlaylists(predicate func(*models.Playlist) bool) (string, error) {
    playlists, err := handler.playlistRepository.FilterPlaylists(predicate)
    if err != nil {
        return "", err
    }

    playlistsJSON, err := json.Marshal(playlists)
    if err != nil {
        return "", err
    }

    return string(playlistsJSON), nil
}

// RemovePlaylist removes a playlist from the database
func (handler *PlaylistHandler) RemovePlaylist(id string) (string, error) {
    playlist, err := handler.playlistRepository.GetPlaylist(id)
    if err != nil {
        return "", err
    }

    err = handler.playlistRepository.RemovePlaylist(playlist)
    if err != nil {
        return "", err
    }

    playlistJSON, err := json.Marshal(playlist)
    if err != nil {
        return "", err
    }

    return string(playlistJSON), nil
}

// UpdatePlaylist updates a playlist in the database 
func (handler *PlaylistHandler) UpdatePlaylist(id, title, description, creatorIDString, createdAtString, imagePath string) (string, error) {
    playlist, err := handler.playlistRepository.GetPlaylist(id)
    if err != nil {
        return "", err
    }

    creatorID, err := utils.StringToInt(creatorIDString)
    if err != nil {
        return "", err
    }

    createdAt, err := utils.StringToInt64(createdAtString)
    if err != nil {
        return "", err
    }

    playlist.Title = title
    playlist.Description = description
    playlist.CreatorID = creatorID
    playlist.CreatedAt = createdAt
    playlist.ImagePath = imagePath

    err = handler.playlistRepository.UpdatePlaylist(playlist)
    if err != nil {
        return "", err
    }

    playlistJSON, err := json.Marshal(playlist)
    if err != nil {
        return "", err
    }

    return string(playlistJSON), nil
}
