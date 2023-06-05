package model

import (
	"encoding/json"
	"log"
	"os"
)

type CharacterInfo struct {
	Id            string
	ImageUrl      string
	Description   string
	MainName      string
	OfficialName  string
	DateOfBirth   string
	Age           string
	Gender        string
	Height        string
	Weight        string
	Abilities     []string
	AbilitiesDesc []string
	AgeRange      string
	AgeRangeDesc  string
	Entities      []string
	EntityDesc    []string
	Roles         []string
	RolesDesc     []string
	CreatorIds    []string
	AnimeIds      []string
	AnimeEpsOccur []string
}

type CharacterInfoList []*CharacterInfo

// Write each struct to it's own file
func (d *CharacterInfo) WriteToFile() error {
	dirPath := "./jsons/character"

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Directory created: %s", dirPath)

	file, err := os.Create("jsons/character/character" + d.Id + ".json")
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

	log.Println("Successfully wrote character", d.Id, "to file.")
	return nil
}

// Write a slice of structs to one json
func (d *CharacterInfoList) WriteToFile() error {
	file, err := os.Create("jsons/characters.json")
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
	log.Println("Successfully wrote", len(*d), "characters to file", file.Name())
	return nil
}
