package main

import (
	"genius-lyrics-parser/client"
	"genius-lyrics-parser/config"
	"genius-lyrics-parser/genius"
	"genius-lyrics-parser/log"
	baseLog "log"
)

const ENV_PATH = ".env.local"

func main() {
	baseLog.Println("starting, load config...")

	config, err := config.LoadConfig(ENV_PATH)
	log.CheckError(err)

	client := client.New()

	var artist_ids = []int{
		1267272, // morgenshtern
		988966,  // face
		1138683, // pasha technik
		1154082, // 1.kla$
		1312547, // black economy
	}

	baseLog.Println("start parsing...")

	for _, artistId := range artist_ids {
		genius.FetchArtistSongs(config, client, artistId)
	}

	baseLog.Println("done, check parsed_lyrics folder in project root")
}
