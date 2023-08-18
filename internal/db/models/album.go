package models

type Album struct {
    ID          int     // Unique ID for the album
    Title       string
    Artist      string
    ReleaseDate int64   // Unix timestamp
    Genre       string 
    SongIDs     []int   // IDs of the songs in the album
    ImagePath   string 
}
