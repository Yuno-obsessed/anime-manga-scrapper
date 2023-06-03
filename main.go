package main

import (
	"crawler/model"
	"log"
)

// func main() {
// 	_ = crawlers.CrawlCharacter(1)
// }

func main() {
	c1 := model.CharacterInfo{
		ImageUrl:    "url",
		Description: "re",
	}
	c2 := model.CharacterInfo{
		Weight: "45",
		Height: "50",
	}
	list := model.CharacterInfoLists{c1, c2}
	err := list.WriteToFile()
	if err != nil {
		log.Println(err)
	}
}
