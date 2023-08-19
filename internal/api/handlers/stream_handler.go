package handlers

import (
	"io"
	"os"

	"github.com/Johanx22x/GoTune"
)

type StreamHandler struct {}

func NewStreamHandler() *StreamHandler {
    return &StreamHandler{}
}

// StreamSongHandler streams a song from the server to the client
func (sh *StreamHandler) StreamSong(songID string) (string, error) {
    // Get songs path
    config := gotune.LoadConfig(true)
    musicPath := config.MusicPath

    // Get song file
	filePath := musicPath + "/" + songID + ".mp3"

    // Open file
	file, err := os.Open(filePath)
	if err != nil {
        return "", err
	}
	defer file.Close()

    // Prepare response
    content, err := io.ReadAll(file)
    if err != nil {
        return "", err
    }

    return string(content), nil
}
