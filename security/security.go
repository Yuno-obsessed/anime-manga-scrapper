package security

import (
	"crawler/config"
	"log"
	"time"
)

func TimeoutRequest() {
	time.Sleep(generateRandomTimeout())
}

func generateRandomTimeout() time.Duration {
	t := time.Second * time.Duration(2+config.Random.Intn(14-2))
	log.Println("Waiting for", t)
	return t
}
