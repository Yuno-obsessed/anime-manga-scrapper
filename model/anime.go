package model

import (
	"encoding/json"
	"log"
	"os"
)

type AnimeInfo struct {
	Id                  string        `json:"id"`
	ImageUrl            string        `json:"image_url"`
	Title               string        `json:"title"`
	OfficialTitleEng    string        `json:"off_title_eng"`
	OfficialTitleJa     string        `json:"off_title_ja"`
	Description         string        `json:"description"`
	Type                string        `json:"type"`
	Year                string        `json:"year"`
	Tags                []string      `json:"tags"`
	TagsDesc            []string      `json:"tags_desc"`
	StaffNCredit        []StaffCredit `json:"staff_credit"` // a lot of stuff here lol
	MainCharacters      []string      `json:"main_characters"`
	SecondaryCharacters []string      `json:"secondary_characters"`
	AppearsCharacters   []string      `json:"appears_characters"`
	Episodes            []Episode     `json:"episodes"`
}

type StaffCredit struct {
	Credit Credit
	Staff  []Staff
}

type Staff struct {
	Id   string
	Name string
}

type Credit struct {
	Id   string
	Name string
}

type Episode struct {
	Id       string
	Num      string
	Title    string
	Duration string
	AirDate  string
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
