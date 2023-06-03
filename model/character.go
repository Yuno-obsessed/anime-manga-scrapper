package model

import (
	"encoding/json"
	"fmt"
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
	AgeRange      string   `json:"age_range"`
	AgeRangeDesc  string   `json:"age_range_desc"`
	Entity        []string `json:"entity"`
	EntityDesc    []string `json:"entity_desc"`
	Role          []string `json:"role"`
	RoleDesc      []string `json:"role_desc"`
	CreatorIds    []string `json:"creator_id"`
	AnimeIds      []string `json:"anime_id"`
	AnimeEpsOccur []string `json:"cnime_eps_occur"`
}

func (d CharacterInfo) Print() {
	fmt.Printf(
		"\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n",
		d.ImageUrl, d.Description, d.MainName, d.OfficialName, d.DateOfBirth,
		d.Age, d.Gender, d.Height, d.Weight, d.AgeRange, d.AgeRangeDesc,
		d.Entity, d.EntityDesc, d.Role, d.RoleDesc, d.CreatorIds, d.AnimeIds, d.AnimeEpsOccur,
	)
}

type CharacterInfoLists []CharacterInfo

func (d CharacterInfoLists) WriteToFile() error {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("jsons/characters.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
