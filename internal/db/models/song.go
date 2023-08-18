package models

type Song struct {
    ID          int     // Unique ID
    Title       string 
    Artist      string
    Duration    int     // Seconds
    ReleaseDate int64   // Unix timestamp
    Genre       string 
    Location    string  // Path to file
}
