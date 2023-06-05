package app

import (
	"crawler/config"
	"crawler/crawlers"
	"crawler/security"
	"errors"
)

func CrawlCharacters() error {
	idInt := 102
	missedIds := 0
	// crawledCharacters := model.CharacterInfoList{}
	var err error
	// for missedIds < 100 {
	for idInt < 104 {

		character, err := crawlers.CrawlCharacter(idInt)
		if errors.Is(err, config.ErrNotFound) {
			missedIds++
			idInt++
			continue
		} else if err != nil {
			return err
		} else if err == nil {
			missedIds = 0
		}

		err = character.WriteToFile()
		if err != nil {
			return err
		}
		// crawledCharacters = append(crawledCharacters, character)

		idInt++
		security.TimeoutRequest()
	}
	return err
}

func CrawlAnime() error {
	idInt := 102
	missedIds := 0
	// crawledAnime := model.AnimeInfoList{}
	var err error
	// for missedIds < 100 {
	for idInt < 104 {

		anime, err := crawlers.CrawlAnime(idInt)
		if errors.Is(err, config.ErrNotFound) {
			missedIds++
			idInt++
			continue
		} else if err != nil {
			return err
		} else if err == nil {
			missedIds = 0
		}

		err = anime.WriteToFile()
		if err != nil {
			return err
		}
		// crawledAnime = append(crawledAnime, &anime)

		idInt++
		security.TimeoutRequest()
	}

	return err
}
