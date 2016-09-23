package gogridclient

import (
	"github.com/guardian/argogo"
	"log"
	"net/http"
)

type MediaApi struct {
	gridService *GridService
}

type Query struct {
	q string
}

type Image struct {
	id string
}

type ImageResponse struct {
	argogo.ArgoResponse
	Images []Image
}

func (imageResponse ImageResponse) UnmarshalArgoData(*argogo.ArgoResponse) {
	imageResponse.Images = []Image{Image{}}
}

type ImageQueryResponse struct {
	argogo.ArgoResponse
	Offset         int
	Length         int
	Total          int
	ImageResponses []ImageResponse
}

func (imageQueryResponse ImageQueryResponse) UnmarshalArgoData(*argogo.ArgoResponse) {
	imageQueryResponse.ImageResponses = []ImageResponse{ImageResponse{}}
}

func (mediaApi MediaApi) ImageQuery(query Query) ([]Image, error) {
	log.Println("gogridclient: Querying images", query.q)

	path := "/images" + "?" + query.q
	url := mediaApi.gridService.TargetUrl + path

	req, err := http.NewRequest("GET", url, nil)

	//if err != nil {
	//	return nil, err
	//}

	//usageService.gridService.Do(req)

	return []Image{Image{}}, nil
}
