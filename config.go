package gotune

type Config struct {
    MusicPath string    // The path to the music directory
    PlaylistPath string // The path to the playlist directory
    AlbumPath string    // The path to the album directory
    DatabasePath string // The path to the database directory
    TempPath string     // The path to the temporary directory

    DatabaseFile string // The path to the database file
    LogFile string      // The path to the log file

    Port string         // The port to run the server on
    Debug bool          // Whether to run in debug mode
}

// LoadConfig loads the default configuration
func LoadConfig(debug bool) *Config {
    config := new(Config)

    config.MusicPath = "./data/music"
    config.PlaylistPath = "./data/playlists"
    config.AlbumPath = "./data/albums"
    config.DatabasePath = "./data"
    config.TempPath = "./data/tmp"

    config.DatabaseFile = config.DatabasePath + "/gotune.db"
    config.LogFile = config.TempPath + "/gotune.log"

    config.Port = "8080"
    config.Debug = debug

    return config
}
