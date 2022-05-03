package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GeniusBaseApiEndpoint        string
	GeniusSongsApiEndpoint       string
	GeniusArtistsApiEndpoint     string
	GeniusUserAuthorizationToken string
}

const (
	GENIUS_BASE_API_ENDPOINT_PARAM_NAME        = "GENIUS_BASE_API_ENDPOINT"
	GENIUS_SONGS_API_ENDPOINT_PARAM_NAME       = "GENIUS_SONGS_API_ENDPOINT"
	GENIUS_ARTISTS_API_ENDPOINT_PARAM_NAME     = "GENIUS_ARTISTS_API_ENDPOINT"
	GENIUS_USER_AUTHORIZATION_TOKEN_PARAM_NAME = "GENIUS_USER_AUTHORIZATION_TOKEN"
)

func LoadConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)

	if err != nil {
		return nil, err
	}

	token := os.Getenv(GENIUS_USER_AUTHORIZATION_TOKEN_PARAM_NAME)

	if len(token) == 0 {
		return nil, errors.New("empty genius authorization token")
	}

	return &Config{
		GeniusBaseApiEndpoint:        os.Getenv(GENIUS_BASE_API_ENDPOINT_PARAM_NAME),
		GeniusSongsApiEndpoint:       os.Getenv(GENIUS_SONGS_API_ENDPOINT_PARAM_NAME),
		GeniusArtistsApiEndpoint:     os.Getenv(GENIUS_ARTISTS_API_ENDPOINT_PARAM_NAME),
		GeniusUserAuthorizationToken: token,
	}, nil

}
