package crawlers

import (
	"crawler/config"
	"crawler/model"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type creatorDocument struct {
	*goquery.Document
	id string
}

func NewCreatorDocument(body io.ReadCloser, id string) characterDocument {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatalln(err)
	}
	defer body.Close()
	return characterDocument{
		Document: doc,
		id:       id,
	}
}

func CrawlCreator(idInt int) (*model.CreatorInfo, error) {
	crawlUrl := "https://anidb.net/creator/"

	id := strconv.Itoa(idInt)

	response, err := config.PrepareRequest(crawlUrl + id)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 404 {
		return nil, config.ErrNotFound
	} else if response.StatusCode != 200 {
		return nil, fmt.Errorf("%v", response.StatusCode)
	}

	doc := NewCreatorDocument(response.Body, id)

	return &model.CreatorInfo{
		Id:       id,
		ImageUrl: doc.parseImage(),
	}, nil
}

func (p creatorDocument) parseImage() string {
	img := p.Find("div.image div.container picture img")
	imageUrl, _ := img.Attr("src")
	return imageUrl
}
