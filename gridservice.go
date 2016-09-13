package gogridclient

import (
	"log"
	"net/http"
	"time"
)

type GridServiceMessage interface {
	Do(*http.Request) (*http.Response, error)
}

type GridService struct {
	GridClient *GridClient
	TargetUrl  string
}

// Generic Grid service (wraps raw messages)
func NewGridService(apiKey string) *GridService {
	client := NewGridClient(apiKey)
	return &GridService{client, ""}
}

func (gridService GridService) Do(request *http.Request) (*http.Response, error) {
	request.Header.Set("X-Gu-Media-Key", gridService.GridClient.ApiKey)

	start := time.Now()

	response, err := gridService.GridClient.HttpClient.Do(request)

	elapsed := time.Since(start)

	log.Printf("gogridclient: Response took %s", elapsed)

	return response, err
}
