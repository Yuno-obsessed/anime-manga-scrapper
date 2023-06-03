package main

import (
	"crawler/app"
	"fmt"
)

func main() {
	charactersResult, err := app.CrawlCharacters()
	if err != nil {
		fmt.Println(err)
	}
	err = charactersResult.WriteToFiles()
	if err != nil {
		fmt.Println(err)
	}
}
