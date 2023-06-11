package config

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/corpix/uarand"
)

var ErrNotFound error = errors.New("404")
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func PrepareRequest(url string) (*http.Response, error) {
	log.Println(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("user-agent", uarand.GetRandom())
	// client, err := generateRandomProxy()
	// if err != nil {
	// 	return nil, err
	// }
	for headerName, headerValue := range request.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	log.Println(response.StatusCode)
	return response, err
}

func TimeoutRequest() {
	time.Sleep(generateRandomTimeout())
}

func generateRandomTimeout() time.Duration {
	t := time.Second * time.Duration(2+random.Intn(14-2))
	log.Println("Waiting for", t)
	return t
}

func generateRandomProxy() (*http.Client, error) {
	// num := Random.Intn(4)
	ur := os.Getenv("proxy" + strconv.Itoa(1))
	proxyUrl, err := url.Parse("http://" + ur)
	fmt.Println(ur)
	if err != nil {
		return http.DefaultClient, err
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}, nil
}
