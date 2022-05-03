package song

import (
	"fmt"
	"genius-lyrics-parser/config"
	"genius-lyrics-parser/log"
	"html"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

type Song struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

const (
	ACCESS_TOKEN_QUERY_PARAM_NAME = "access_token"
	TEXT_FORMAT_PARAM_NAME        = "text_format"

	PAGE_QUERY_PARAM_NAME     = "page"
	PER_PAGE_QUERY_PARAM_NAME = "per_page"
)

const (
	TEXT_FORMAT_DOM   = "dom"
	TEXT_FORMAT_HTML  = "html"
	TEXT_FORMAT_PLAIN = "plain"
)

func ParseSongPage(response *http.Response) string {
	doc, readErr := goquery.NewDocumentFromReader(response.Body)
	log.CheckError(readErr)

	htmlText, htmlErr := doc.Find(`div[data-lyrics-container="true"]`).Html()
	log.CheckError(htmlErr)

	return sanitizeLyrics(htmlText)
}

// TODO: Refactor with net/url
func BuildUrl(config *config.Config, id int, textFormat string) string {
	return fmt.Sprintf(
		"%s/%s/%d?%s=%s&%s=%s",
		config.GeniusBaseApiEndpoint,
		config.GeniusSongsApiEndpoint,
		id,
		ACCESS_TOKEN_QUERY_PARAM_NAME,
		config.GeniusUserAuthorizationToken,
		TEXT_FORMAT_PARAM_NAME,
		textFormat,
	)
}

// TODO: Refactor with net/url
func BuildArtistSongsUrl(config *config.Config, id int, page int, perPage int) string {
	return fmt.Sprintf(
		"%s/%s/%d/songs?%s=%s&%s=%d&%s=%d",
		config.GeniusBaseApiEndpoint,
		config.GeniusArtistsApiEndpoint,
		id,
		ACCESS_TOKEN_QUERY_PARAM_NAME,
		config.GeniusUserAuthorizationToken,
		PAGE_QUERY_PARAM_NAME,
		page,
		PER_PAGE_QUERY_PARAM_NAME,
		perPage,
	)
}

func sanitizeLyrics(lyricsHtml string) string {
	text := bluemonday.StripTagsPolicy().Sanitize(strings.ReplaceAll(lyricsHtml, "<br/>", "\n"))

	return html.UnescapeString(text)
}
