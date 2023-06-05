package config

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/corpix/uarand"
)

var ErrNotFound error = errors.New("404")
var Random = rand.New(rand.NewSource(time.Now().UnixNano()))

func PrepareRequest(url string) (*http.Response, error) {
	log.Println(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("user-agent", uarand.GetRandom())
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	log.Println(response.StatusCode)
	return response, err
}