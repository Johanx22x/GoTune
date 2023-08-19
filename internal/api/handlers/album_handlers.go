package handlers

import (
    "encoding/json"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
    "github.com/Johanx22x/GoTune/internal/db/models"
    "github.com/Johanx22x/GoTune/internal/utils"
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
func (handler *AlbumHandler) GetAlbums() (string, error) {
    albums, err := handler.albumRepository.GetAlbums()
    if err != nil {
        return "", err
    }

    albumsJSON, err := json.Marshal(albums)
    if err != nil {
        return "", err
    }

    return string(albumsJSON), nil
}

// GetAlbum returns a album by id 
func (handler *AlbumHandler) GetAlbum(id string) (string, error) {
    album, err := handler.albumRepository.GetAlbum(id)
    if err != nil {
        return "", err
    }

    albumJSON, err := json.Marshal(album)
    if err != nil {
        return "", err
    }

    return string(albumJSON), nil
}

// CreateAlbum creates a new album
func (handler *AlbumHandler) CreateAlbum(title, artist, releaseDateString, genre, imagePath string) (string, error) {
    releaseDate, err := utils.StringToInt64(releaseDateString)
    if err != nil {
        return "", err
    }

    album := models.Album{
        Title: title,
        Artist: artist,
        ReleaseDate: releaseDate,
        Genre: genre,
        ImagePath: imagePath,
    }

    err = handler.albumRepository.CreateAlbum(&album)
    if err != nil {
        return "", err
    }

    albumJSON, err := json.Marshal(album)
    if err != nil {
        return "", err
    }

    return string(albumJSON), nil
}

// FilterAlbums returns all albums that match the predicate 
func (handler *AlbumHandler) FilterAlbums(predicate func(album *models.Album) bool) (string, error) {
    albums, err := handler.albumRepository.FilterAlbums(predicate)
    if err != nil {
        return "", err
    }

    albumsJSON, err := json.Marshal(albums)
    if err != nil {
        return "", err
    }

    return string(albumsJSON), nil
}

// RemoveAlbum removes a album from the database 
func (handler *AlbumHandler) RemoveAlbum(id string) (string, error) {
    album, err := handler.albumRepository.GetAlbum(id)
    if err != nil {
        return "", err
    }

    err = handler.albumRepository.RemoveAlbum(album)
    if err != nil {
        return "", err
    }

    albumJSON, err := json.Marshal(album)
    if err != nil {
        return "", err
    }

    return string(albumJSON), nil
}

// UpdateAlbum updates a album in the database 
func (handler *AlbumHandler) UpdateAlbum(id, title, artist, releaseDateString, genre, imagePath string) (string, error) {
    album, err := handler.albumRepository.GetAlbum(id)
    if err != nil {
        return "", err
    }

    album.Title = title
    album.Artist = artist
    album.Genre = genre
    album.ImagePath = imagePath

    releaseDate, err := utils.StringToInt64(releaseDateString)
    if err != nil {
        return "", err
    }

    album.ReleaseDate = releaseDate

    err = handler.albumRepository.UpdateAlbum(album)
    if err != nil {
        return "", err
    }

    albumJSON, err := json.Marshal(album)
    if err != nil {
        return "", err
    }

    return string(albumJSON), nil
}
