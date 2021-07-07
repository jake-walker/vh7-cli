package quark

import (
	"fmt"
)

func (s *ApiService) GetInfo(link string) (*Response, error) {
	urlResponse := new(Response)
	path := fmt.Sprintf("/info/%s", link)
	_, err := s.sling.New().Get(path).ReceiveSuccess(urlResponse)

	return urlResponse, err
}

func (s *ApiService) GetLanguages() (*[]Language, error) {
	languages := new([]Language)
	_, err := s.sling.New().Get("/languages").ReceiveSuccess(languages)

	return languages, err
}
