package quark

import (
	"net/http"

	"github.com/dghubble/sling"
)

type ApiService struct {
	sling *sling.Sling
}

func NewApiService(httpClient *http.Client, apiUrl string) *ApiService {
	return &ApiService{
		sling: sling.New().Client(httpClient).Base(apiUrl),
	}
}
