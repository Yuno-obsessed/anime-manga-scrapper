package crawlers

import (
	"crawler/model"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type AnimeDocument struct {
	*goquery.Document
	id string
}

func NewAnimeDocument(body io.ReadCloser, id string) CharacterDocument {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatalln(err)
	}
	defer body.Close()
	return CharacterDocument{
		Document: doc,
		id:       id,
	}
}

func CrawlAnime(idInt int) (model.AnimeInfo, error) {
	return model.AnimeInfo{}, nil
}
