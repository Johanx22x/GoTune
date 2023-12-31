package api

import (
    "database/sql"
    "strings"
    "errors"

    "github.com/Johanx22x/GoTune/internal/api/handlers"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
)

// ProcessRequest processes a TCP request
func ProcessRequest(message string, db *sql.DB) (string, error) {
    words := strings.Split(message, " ")

    switch words[0] {
    case "GET":
        return processGetRequest(words, db)
    case "POST":
        return processPostRequest(words, db)
    case "PUT":
        return processPutRequest(words, db)
    case "DELETE":
        return processDeleteRequest(words, db)
    case "STREAM":
        return processStreamRequest(words, db)
    case "HANDSHAKE":
        return "OK", nil
    default:
        return "", errors.New("invalid request")
    }
}

// processGetRequest processes a GET request
func processGetRequest(words []string, db *sql.DB) (string, error) {
    switch words[1] {
    case "songs":
        songsHandler := handlers.NewSongHandler(repositories.NewSongRepository(db))
        if len(words) == 2 {
            return songsHandler.GetSongs()
        } else {
            return songsHandler.GetSong(words[2])
        }
    case "playlists":
        playlistsHandler := handlers.NewPlaylistHandler(repositories.NewPlaylistRepository(db))
        if len(words) == 2 { 
            return playlistsHandler.GetPlaylists()
        } else { 
            return playlistsHandler.GetPlaylist(words[2])
        }
    default:
        return "", errors.New("invalid request")
    }
}

// processPostRequest processes a POST request
func processPostRequest(words []string, db *sql.DB) (string, error) {
    switch words[1] {
    case "songs":
        songsHandler := handlers.NewSongHandler(repositories.NewSongRepository(db))
        return songsHandler.CreateSong(words[2], words[3], words[4], words[5], words[6])
    case "playlists":
        playlistsHandler := handlers.NewPlaylistHandler(repositories.NewPlaylistRepository(db))
        return playlistsHandler.CreatePlaylist(words[2], words[3], words[4], words[5], words[6])
    default:
        return "", errors.New("invalid request")
    }
}

// processPutRequest processes a PUT request
func processPutRequest(words []string, db *sql.DB) (string, error) {
    switch words[1] {
    case "songs":
        songsHandler := handlers.NewSongHandler(repositories.NewSongRepository(db))
        return songsHandler.UpdateSong(words[2], words[3], words[4], words[5], words[6], words[7])
    case "playlists":
        playlistsHandler := handlers.NewPlaylistHandler(repositories.NewPlaylistRepository(db))
        return playlistsHandler.UpdatePlaylist(words[2], words[3], words[4], words[5], words[6], words[7])
    default:
        return "", errors.New("invalid request")
    }
}

// processDeleteRequest processes a DELETE request
func processDeleteRequest(words []string, db *sql.DB) (string, error) {
    switch words[1] {
    case "songs":
        songsHandler := handlers.NewSongHandler(repositories.NewSongRepository(db))
        return songsHandler.RemoveSong(words[2])
    case "playlists":
        playlistsHandler := handlers.NewPlaylistHandler(repositories.NewPlaylistRepository(db))
        return playlistsHandler.RemovePlaylist(words[2])
    default:
        return "", errors.New("invalid request")
    }
}

// processStreamRequest processes a STREAM request
func processStreamRequest(words []string, db *sql.DB) (string, error) {
    // Create stream handler
    streamHandler := handlers.NewStreamHandler()

    // Get the song
    response, err := streamHandler.StreamSong(words[1])
    if err != nil {
        return "", err
    }

    return response, nil
}
