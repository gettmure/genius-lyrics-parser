# Description
[Genius](https://genius.com) lyrics parser. It was created to make dataset for neural network, which writes its own lyrics based on parsed lyrics.

Script performs parsing for every song per artist. It is hardcoded (for now) in `main.go`:

```GO
var artist_ids = []int{
    1267272, // morgenshtern
    988966,  // face
    1138683, // pasha technik
    1154082, // 1.kla$
    1312547, // black economy
}

```

# Tools
- Docker version 20.10.14
- Docker-compose version 1.29.2
- Go 1.18
- Makefile

# Usage
Envinronment variables are stored in `.env.local` file. 

Configure `GENIUS_USER_AUTHORIZATION_TOKEN` value from your [Genius API](https://docs.genius.com//) application and proxy `HTTP_PROXY` (if needed). See example in `.env`:
```bash
GENIUS_BASE_API_ENDPOINT=https://api.genius.com
GENIUS_SONGS_API_ENDPOINT=songs
GENIUS_ARTISTS_API_ENDPOINT=artists
GENIUS_USER_AUTHORIZATION_TOKEN=token_from_genius_api_application
HTTP_PROXY=http://user:password@ip:port

```

Run one of the following commands to parse lyrics:
- `make parse`
- `docker-compose up`
- `go run main.go`
- `go build && ./genius-lyrics-parser`

Lyrics will be stored in `./parsed_lyrics` folder.