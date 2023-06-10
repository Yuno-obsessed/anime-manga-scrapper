package main

import (
	"crawler/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	err = app.CrawlAnime()
	if err != nil {
		log.Println(err)
	}
}
