# GoTune API Documentation

The GoTune API allows you to interact with the GoTune music application using TCP requests. The API provides endpoints to manage songs, users, playlists, and albums.

## Sending TCP Requests

To send requests to the GoTune API, you'll need to establish a TCP connection and send properly formatted messages. Each message should be space-separated words representing the request type and parameters.

### Request Format

The general format of a request message is:

```bash
<REQUEST_TYPE> <RESOURCE> <PARAMETERS>
```

- `<REQUEST_TYPE>`: The type of the request, such as GET, POST, PUT, DELETE, or STREAM.
- `<RESOURCE>`: The resource you want to interact with, like songs, users, playlists, or albums.
- `<PARAMETERS>`: Any parameters required for the request, such as IDs or additional data.

## Supported Endpoints

### GET Requests

- Get Songs

```bash
Get Songs
GET songs <SONG_ID>
```

- Get Users

```bash
GET users
GET users <USER_ID>
```

- Get Playlists

```bash
GET playlists
GET playlists <PLAYLIST_ID>
```

- Get Albums

```bash
GET albums
GET albums <ALBUM_ID>
```

### POST Requests

- Create Song

```bash
POST songs <TITLE> <ARTIST> <DURATION> <RELEASE_DATE> <GENRE>
```

- Create User

```bash
POST users <USERNAME> <PASSWORD>
```

- Create Playlist

```bash
POST playlists <TITLE> <DESCRIPTION> <CREATOR_ID> <CREATED_AT> <SONG_IDS>
```

- Create Album

```bash
POST albums <TITLE> <ARTIST> <RELEASE_DATE> <GENRE> <SONG_IDS> <IMAGE_PATH>
```

### PUT Requests

- Update Song

```bash
PUT songs <SONG_ID> <NEW_TITLE> <NEW_ARTIST> <NEW_DURATION> <NEW_RELEASE_DATE> <NEW_GENRE>
```

- Update User

```bash
PUT users <USER_ID> <NEW_USERNAME> <NEW_PASSWORD>
```

- Update Playlist

```bash
PUT playlists <PLAYLIST_ID> <NEW_TITLE> <NEW_DESCRIPTION> <NEW_CREATOR_ID> <NEW_CREATED_AT> <NEW_SONG_IDS>
```

- Update Album

```bash
PUT albums <ALBUM_ID> <NEW_TITLE> <NEW_ARTIST> <NEW_RELEASE_DATE> <NEW_GENRE> <NEW_SONG_IDS> <NEW_IMAGE_PATH>
```

### DELETE Requests

- Delete Songs

```bash
DELETE songs <SONG_ID>
```

- Delete Users

```bash
DELETE users <USER_ID>
```

- Delete Playlist

```bash
DELETE playlists <PLAYLIST_ID>
```

- Delete Album

```bash
    DELETE albums <ALBUM_ID>
    ```

### STREAM Requests

- Stream Song

```bash
STREAM <SONG_ID>
```
