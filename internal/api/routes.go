package api

import (
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
    "github.com/Johanx22x/GoTune/internal/api/handlers"
    "github.com/Johanx22x/GoTune/internal/db/repositories"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(router *mux.Router, db *sql.DB) *mux.Router {
    // Create songs handler
    songRepository := repositories.NewSongRepository(db)
    songHandler := handlers.NewSongHandler(songRepository)

    // Register songs routes
    router.HandleFunc("/songs", songHandler.GetSongs).Methods(http.MethodGet)
    router.HandleFunc("/songs/{id}", songHandler.GetSong).Methods(http.MethodGet)
    router.HandleFunc("/songs", songHandler.CreateSong).Methods(http.MethodPost)
    router.HandleFunc("/songs/{id}", songHandler.UpdateSong).Methods(http.MethodPut)
    router.HandleFunc("/songs/{id}", songHandler.RemoveSong).Methods(http.MethodDelete)

    // Create users handler 
    userRepository := repositories.NewUserRepository(db)
    userHandler := handlers.NewUserHandler(userRepository)

    // Register users routes
    router.HandleFunc("/users", userHandler.GetUsers).Methods(http.MethodGet)
    router.HandleFunc("/users/{id}", userHandler.GetUser).Methods(http.MethodGet)
    router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
    router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods(http.MethodPut)
    router.HandleFunc("/users/{id}", userHandler.RemoveUser).Methods(http.MethodDelete)

    // Create playlists handler 
    playlistRepository := repositories.NewPlaylistRepository(db)
    playlistHandler := handlers.NewPlaylistHandler(playlistRepository)

    // Register playlists routes 
    router.HandleFunc("/playlists", playlistHandler.GetPlaylists).Methods(http.MethodGet)
    router.HandleFunc("/playlists/{id}", playlistHandler.GetPlaylist).Methods(http.MethodGet)
    router.HandleFunc("/playlists", playlistHandler.CreatePlaylist).Methods(http.MethodPost)
    router.HandleFunc("/playlists/{id}", playlistHandler.UpdatePlaylist).Methods(http.MethodPut)
    router.HandleFunc("/playlists/{id}", playlistHandler.RemovePlaylist).Methods(http.MethodDelete)

    // Create albums handler 
    albumRepository := repositories.NewAlbumRepository(db)
    albumHandler := handlers.NewAlbumHandler(albumRepository)

    // Register albums routes 
    router.HandleFunc("/albums", albumHandler.GetAlbums).Methods(http.MethodGet)
    router.HandleFunc("/albums/{id}", albumHandler.GetAlbum).Methods(http.MethodGet)
    router.HandleFunc("/albums", albumHandler.CreateAlbum).Methods(http.MethodPost)
    router.HandleFunc("/albums/{id}", albumHandler.UpdateAlbum).Methods(http.MethodPut)
    router.HandleFunc("/albums/{id}", albumHandler.RemoveAlbum).Methods(http.MethodDelete)

    // Create stream handler 
    streamHandler := handlers.NewStreamHandler()

    // Register stream routes 
    router.HandleFunc("/stream/{id}", streamHandler.StreamSong).Methods(http.MethodGet)

    return router
}
