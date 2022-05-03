package client

import (
	"genius-lyrics-parser/log"
	baseLog "log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const REQUEST_MINUTES_TIMEOUT time.Duration = 5

func New() *http.Client {
	timeout := time.Minute * REQUEST_MINUTES_TIMEOUT
	proxyEnvURL := os.Getenv("HTTP_PROXY")

	if len(proxyEnvURL) == 0 {
		baseLog.Println("no http proxy provided, using default client")

		return &http.Client{
			Timeout: timeout,
		}
	}

	proxyURL, err := url.Parse(os.Getenv("HTTP_PROXY"))
	log.CheckError(err)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}
