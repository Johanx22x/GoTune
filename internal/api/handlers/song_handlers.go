package handlers

import (
    "encoding/json"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
    "github.com/Johanx22x/GoTune/internal/db/models"
    "github.com/Johanx22x/GoTune/internal/utils"
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
func (handler *SongHandler) GetSongs() (string, error) {
    songs, err := handler.songRepository.GetSongs()
    if err != nil {
        return "", err
    }

    songsJSON, err := json.Marshal(songs)
    if err != nil {
        return "", err
    }

    return string(songsJSON), nil
}

// GetSong returns a song by id 
func (handler *SongHandler) GetSong(id string) (string, error) {
    song, err := handler.songRepository.GetSong(id)
    if err != nil {
        return "", err
    }

    songJSON, err := json.Marshal(song)
    if err != nil {
        return "", err
    }

    return string(songJSON), nil
}

// CreateSong creates a new song 
func (handler *SongHandler) CreateSong(title, artist, genre, durationString, releaseDateString string) (string, error) {
    // Cast duration to int 
    duration, err := utils.StringToInt(durationString)
    if err != nil {
        return "", err
    }

    // Cast releaseDate to int64 
    releaseDate, err := utils.StringToInt64(releaseDateString)
    if err != nil {
        return "", err
    }

    song := &models.Song{
        Title:       title,
        Artist:      artist,
        Genre:       genre,
        Duration:    duration,
        ReleaseDate: releaseDate,
    }

    err = handler.songRepository.CreateSong(song)
    if err != nil {
        return "", err
    }

    songJSON, err := json.Marshal(song)
    if err != nil {
        return "", err
    }

    return string(songJSON), nil
}

// FilterSongs returns all songs that match the predicate
func (handler *SongHandler) FilterSongs(predicate func(*models.Song) bool) (string, error) {
    songs, err := handler.songRepository.FilterSongs(predicate)
    if err != nil {
        return "", err
    }

    songsJSON, err := json.Marshal(songs)
    if err != nil {
        return "", err
    }

    return string(songsJSON), nil
}

// RemoveSong removes a song from the database
func (handler *SongHandler) RemoveSong(id string) (string, error) {
    song, err := handler.songRepository.GetSong(id)
    if err != nil {
        return "", err
    }

    err = handler.songRepository.RemoveSong(song)
    if err != nil {
        return "", err
    }

    songJSON, err := json.Marshal(song)
    if err != nil {
        return "", err
    }

    return string(songJSON), nil
}

// UpdateSong updates a song in the database
func (handler *SongHandler) UpdateSong(id, title, artist, genre, durationString, releaseDateString string) (string, error) {
    song, err := handler.songRepository.GetSong(id)
    if err != nil {
        return "", err
    }

    // Cast duration to int 
    duration, err := utils.StringToInt(durationString)
    if err != nil {
        return "", err
    }

    // Cast releaseDate to int64 
    releaseDate, err := utils.StringToInt64(releaseDateString)
    if err != nil {
        return "", err
    }

    song.Title = title
    song.Artist = artist
    song.Genre = genre 
    song.Duration = duration 
    song.ReleaseDate = releaseDate 

    err = handler.songRepository.UpdateSong(song)
    if err != nil {
        return "", err
    }

    songJSON, err := json.Marshal(song)
    if err != nil {
        return "", err
    }

    return string(songJSON), nil
}
