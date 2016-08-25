package gogridclient

import (
	"net/http"
)

type GridClient struct {
	ApiKey     string
	HttpClient *http.Client
}

func NewGridClient(apiKey string) *GridClient {
	client := GridClient{
		ApiKey:     apiKey,
		HttpClient: &http.Client{},
	}

	return &client
}
