package http

import (
	"encoding/json"
	"genius-lyrics-parser/log"
	"io"
	"net/http"
)

func NewGet(url string) *http.Request {
	request, err := http.NewRequest("GET", url, nil)
	log.CheckError(err)

	return request
}

func Get(client *http.Client, url string) *http.Response {
	request := NewGet(url)
	response := Response(client, request)

	return response
}

func Response(client *http.Client, request *http.Request) *http.Response {
	response, err := client.Do(request)
	log.CheckError(err)

	return response
}

func DecodeJson(body io.ReadCloser, target interface{}) error {
	defer body.Close()

	return json.NewDecoder(body).Decode(target)
}
