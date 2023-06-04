package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CharacterInfo struct {
	Id            string   `json:"id"`
	ImageUrl      string   `json:"image_url"`
	Description   string   `json:"description"`
	MainName      string   `json:"main_name"`
	OfficialName  string   `json:"official_name"`
	DateOfBirth   string   `json:"date_of_birth"`
	Age           string   `json:"age"`
	Gender        string   `json:"gender"`
	Height        string   `json:"height"`
	Weight        string   `json:"weight"`
	Abilities     []string `json:"abilities"`
	AbilitiesDesc []string `json:"abilities_desc"`
	AgeRange      string   `json:"age_range"`
	AgeRangeDesc  string   `json:"age_range_desc"`
	Entity        []string `json:"entity"`
	EntityDesc    []string `json:"entity_desc"`
	Role          []string `json:"role"`
	RoleDesc      []string `json:"role_desc"`
	CreatorIds    []string `json:"creator_id"`
	AnimeIds      []string `json:"anime_id"`
	AnimeEpsOccur []string `json:"anime_eps_occur"`
}

func (d CharacterInfo) Print() {
	fmt.Printf(
		"\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n",
		d.ImageUrl, d.Description, d.MainName, d.OfficialName, d.DateOfBirth,
		d.Age, d.Gender, d.Height, d.Weight, d.Abilities, d.AbilitiesDesc, d.AgeRange, d.AgeRangeDesc,
		d.Entity, d.EntityDesc, d.Role, d.RoleDesc, d.CreatorIds, d.AnimeIds, d.AnimeEpsOccur,
	)
}

type CharacterInfoList []CharacterInfo

// Write all structs to one as an array
func (d CharacterInfoList) WriteToFile() error {
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

	log.Println("Successfully wrote", len(d), "characters to file.")
	return nil
}

// Write each struct to it's own id-based json
func (d CharacterInfoList) WriteToFiles() error {
	for _, character := range d {
		file, err := os.Create("jsons/character" + character.Id + ".json")
		if err != nil {
			return err
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")

		err = encoder.Encode(character)
		if err != nil {
			return err
		}
		log.Println("Successfully wrote", character.Id, "to file.")
	}
	return nil
}
