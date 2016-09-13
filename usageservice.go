package gogridclient

import (
	"log"
	"net/http"
)

type UsageService struct {
	gridService *GridService
}

func (usageService UsageService) Reindex(id string) (*http.Response, error) {
	log.Println("gogridclient: Attempting reindex for", id)

	path := "/usages/digital/content/" + id + "/reindex"
	url := usageService.gridService.TargetUrl + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return usageService.gridService.Do(req)
}

func NewUsageService(targetUrl string, apiKey string) *UsageService {
	gridService := NewGridService(apiKey)
	gridService.TargetUrl = targetUrl

	return &UsageService{gridService}
}
