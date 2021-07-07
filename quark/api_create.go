package quark

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func (s *ApiService) CreateShorten(url string) (*Response, error) {
	data := &Url{
		Url: url,
	}

	urlResponse := new(Response)
	_, err := s.sling.New().Post("/shorten").BodyJSON(data).ReceiveSuccess(urlResponse)

	return urlResponse, err
}

func (s *ApiService) CreatePaste(code string, language string) (*Response, error) {
	data := &Paste{
		Language: language,
		Code:     code,
	}

	pasteResponse := new(Response)
	_, err := s.sling.New().Post("/paste").BodyJSON(data).ReceiveSuccess(pasteResponse)

	return pasteResponse, err
}

func (s *ApiService) CreateUpload(file *os.File) (*Response, error) {
	// Build the multipart body
	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	uploadResponse := new(Response)
	_, err := s.sling.New().Post("/upload").Body(data).Set("Content-Type", writer.FormDataContentType()).ReceiveSuccess(uploadResponse)

	return uploadResponse, err
}
