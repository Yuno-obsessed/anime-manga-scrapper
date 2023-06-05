package model

import (
	"encoding/json"
	"log"
	"os"
)

type AnimeInfo struct {
	Id       string `json:"id"`
	ImageUrl string `json:"image_url"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Year     string `json:"year"`
}

type AnimeInfoList []*AnimeInfo

// Write each struct to it's own file
func (d *AnimeInfo) WriteToFile() error {
	dirPath := "./jsons/anime"

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Directory created: %s", dirPath)
	file, err := os.Create("jsons/anime/anime" + d.Id + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(d)
	if err != nil {
		return err
	}

	log.Println("Successfully wrote anime", d.Id, "to file.")
	return nil
}

// Write a slice of structs to one json
func (d *AnimeInfoList) WriteToFile() error {
	file, err := os.Create("jsons/anime.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(d)
	if err != nil {
		return err
	}
	log.Println("Successfully wrote", len(*d), "anime to file", file.Name())

	return nil
}
