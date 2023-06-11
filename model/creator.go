package model

import (
	"encoding/json"
	"log"
	"os"
)

type CreatorInfo struct {
	Id           string   `json:"id"`
	ImageUrl     string   `json:"image_url"`
	MainName     string   `json:"main_name"`
	OfficialName string   `json:"official_name"`
	Description  string   `json:"description"`
	DateOfBirth  string   `json:"date_of_birth"`
	Gender       string   `json:"gender"`
	Jobs         []string `json:"jobs"`
	JobsDesc     []string `json:"jobs_desc"`
	Nationality  string   `json:"nationality"`
	AnimeIds     []string `json:"anime_ids"`
}

type CreatorInfoList []*CreatorInfo

// Write each struct to it's own file
func (d *CreatorInfo) WriteToFile() error {
	dirPath := "./jsons/creator"

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Directory created: %s", dirPath)
	file, err := os.Create("jsons/creator/creator" + d.Id + ".json")
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

	log.Println("Successfully wrote creator", d.Id, "to file.")
	return nil
}

// Write a slice of structs to one json
func (d *CreatorInfoList) WriteToFile() error {
	file, err := os.Create("jsons/creators.json")
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
	log.Println("Successfully wrote", len(*d), "creators to file", file.Name())

	return nil
}
