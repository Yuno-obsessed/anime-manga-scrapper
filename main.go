package main

import (
	"crawler/app"
	"fmt"
)

func main() {
	err := app.CrawlCharacters()
	if err != nil {
		fmt.Println(err)
	}
}
