package config

import (
	"errors"
	"log"
	"net/http"
)

var ErrNotFound error = errors.New("404")

func PrepareRequest(url string) (*http.Response, error) {
	log.Println(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.57")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	log.Println(response.StatusCode)
	return response, err
}
