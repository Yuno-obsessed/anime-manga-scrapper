package app

import (
	"crawler/config"
	"crawler/crawlers"
	"crawler/model"
	"errors"
)

func CrawlCharacters() (model.CharacterInfoList, error) {
	idInt := 1
	missedIds := 0
	crawledCharacters := model.CharacterInfoList{}
	var err error
	// for missedIds < 100 {
	for len(crawledCharacters) < 3 {

		character, err := crawlers.CrawlCharacter(idInt)
		if errors.Is(err, config.ErrNotFound) {
			missedIds++
		} else if err != nil {
			return nil, err
		} else if err == nil {
			missedIds = 0
		}

		crawledCharacters = append(crawledCharacters, character)

		idInt++
	}

	return crawledCharacters, err
}
