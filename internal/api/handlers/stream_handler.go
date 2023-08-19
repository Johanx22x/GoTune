package handlers

import (
    "io"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "github.com/Johanx22x/GoTune"
)

type StreamHandler struct {}

func NewStreamHandler() *StreamHandler {
    return &StreamHandler{}
}

// StreamSongHandler streams a song from the server to the client
func (sh *StreamHandler) StreamSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songID := vars["id"]

    config := gotune.LoadConfig(true)
    musicPath := config.MusicPath

	filePath := musicPath + songID + ".mp3"

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "audio/mpeg")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error streaming audio", http.StatusInternalServerError)
		return
	}
}
