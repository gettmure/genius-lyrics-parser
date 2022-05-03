package genius

import (
	"fmt"
	"genius-lyrics-parser/config"
	"genius-lyrics-parser/file"
	"genius-lyrics-parser/http"
	"genius-lyrics-parser/song"
	baseHttp "net/http"
	"path/filepath"
	"strings"
)

type GeniusResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		Song     *song.Song   `json:"song"`
		Songs    []*song.Song `json:"songs"`
		NextPage *int         `json:"next_page"`
	} `json:"response"`
}

func FetchSong(config *config.Config, client *baseHttp.Client, id int) *song.Song {
	songUrl := song.BuildUrl(config, id, song.TEXT_FORMAT_PLAIN)
	response := http.Get(client, songUrl)

	return parseGeniusResponse(*response).Response.Song
}

func FetchArtistSongs(config *config.Config, client *baseHttp.Client, id int) {
	for page := 1; page != -1; {
		artistSongsUrl := song.BuildArtistSongsUrl(config, id, page, 50)
		response := http.Get(client, artistSongsUrl)
		geniusResponse := parseGeniusResponse(*response)

		for _, song := range geniusResponse.Response.Songs {
			text := fetchSongLyrics(client, song.Url)
			filename := fmt.Sprintf("%s.txt", filepath.Join("parsed_lyrics", strings.TrimLeft(song.Path, "/")))

			file.Write(filename, text)
		}

		if geniusResponse.Response.NextPage == nil {
			break
		}

		page = *geniusResponse.Response.NextPage
	}
}

func fetchSongLyrics(client *baseHttp.Client, url string) string {
	response := http.Get(client, url)

	return song.ParseSongPage(response)
}

func parseGeniusResponse(response baseHttp.Response) GeniusResponse {
	geniusResponse := &GeniusResponse{}
	http.DecodeJson(response.Body, geniusResponse)

	return *geniusResponse
}
