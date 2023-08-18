package models

type Playlist struct {
    ID          int 
    Title       string 
    Description string 
    CreatorID   int
    CreatedAt   int64
    SongIDs     []int
    ImagePath   string
}
