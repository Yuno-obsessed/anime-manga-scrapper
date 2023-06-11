package crawlers

import (
	"crawler/config"
	"crawler/model"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type animeDocument struct {
	*goquery.Document
	id string
}

func NewAnimeDocument(body io.ReadCloser, id string) animeDocument {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatalln(err)
	}
	defer body.Close()
	return animeDocument{
		Document: doc,
		id:       id,
	}
}

func CrawlAnime(idInt int) (*model.AnimeInfo, error) {
	crawlUrl := "https://anidb.net/anime/"

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

	doc := NewAnimeDocument(response.Body, id)
	return &model.AnimeInfo{
		Id:                  id,
		ImageUrl:            doc.parseImage(),
		Title:               doc.parseTitle(),
		OfficialTitleEng:    doc.parseOfficialTitleEN(),
		OfficialTitleJa:     doc.parseOfficialTitleJa(),
		Description:         doc.parseDescription(),
		Type:                doc.parseType(),
		Year:                doc.parseYear(),
		Tags:                doc.parseTags(),
		TagsDesc:            doc.parseTagsDesc(),
		StaffNCredit:        doc.parseStaffNCredit(),
		MainCharacters:      doc.parseMainCharacters(),
		SecondaryCharacters: doc.parseSecondaryCharacters(),
		AppearsCharacters:   doc.parseAppearsCharacters(),
		Episodes:            doc.getEpisodes(),
	}, nil
}

func (p animeDocument) parseImage() string {
	img := p.Find("div.image div.container picture img")
	imageUrl, _ := img.Attr("src")
	return imageUrl
}

func (p animeDocument) parseTitle() string {
	return p.Find("div.g_definitionlist tbody tr.g_odd.romaji td span").First().Text()
}

func (p animeDocument) parseOfficialTitleEN() string {
	found := ""
	p.Find("div.g_definitionlist tbody tr.official.verified td.value").Each(func(i int, e *goquery.Selection) {
		if e.Find("span.icons span[title*=\"language: english\"] span").First().Text() != "" {
			found = e.Find("label").First().Text()
		}
	})
	return found
}

func (p animeDocument) parseOfficialTitleJa() string {
	found := ""
	p.Find("div.g_definitionlist tbody tr.official.verified td.value").Each(func(i int, e *goquery.Selection) {
		if e.Find("span.icons span[title*=\"language: japanese\"] span").First().Text() != "" {
			found = e.Find("label").First().Text()
		}
	})
	return found
}

func (p animeDocument) parseDescription() string {
	return trimChars(p.Find("div.g_bubble.g_section.desc").Text())
}

func (p animeDocument) parseType() string {
	return p.Find("div.g_definitionlist tbody tr.type td").First().Text()
}

func (p animeDocument) parseYear() string {
	return p.Find("div.g_definitionlist tbody tr.g_odd.year td.value").First().Text()
}

func (p animeDocument) parseTags() []string {
	tags := []string{}
	p.Find("div.g_definitionlist tbody tr.tags td.value span.g_tag").Each(func(i int, e *goquery.Selection) {
		tags = append(tags, e.Find("div.g_definitionlist tbody tr.tags td.value span a span.tagname").Text())
	})
	return tags
}

func (p animeDocument) parseTagsDesc() []string {
	tagsDesc := []string{}
	p.Find("div.g_definitionlist tbody tr.tags td.value span.g_tag").Each(func(i int, e *goquery.Selection) {
		tagsDesc = append(tagsDesc, e.Find("a span.wrapper").Text())
	})
	return tagsDesc
}

func (p animeDocument) parseStaffNCredit() []model.StaffCredit {
	var staffCredits []model.StaffCredit
	var staffCredit model.StaffCredit
	var credit model.Credit
	var moreStaff []model.Staff
	p.Find("div.container div.staffblock").Each(func(i int, block *goquery.Selection) {
		block.Find("div.container div.staffblock table tbody tr").Each(func(i int, c *goquery.Selection) {
			if c.Find("td.credit").Text() == "" {
				el := c.Find("td.name.creator")
				moreStaff = append(moreStaff, getStaff(el))
			}
			if moreStaff != nil && hasCredit(c) {
				staffCredit = model.StaffCredit{credit, moreStaff}
				moreStaff = nil
			} else if moreStaff == nil && hasCredit(c) {
				moreStaff = append(moreStaff, getStaff(c.Find("td.name.creator")))
				credit = getCredit(c.Find("td.credit"))
				staffCredit = model.StaffCredit{credit, moreStaff}
				moreStaff = nil
			}
			if hasCredit(c.Next()) {
				staffCredits = append(staffCredits, staffCredit)
			}
		})
	})
	return staffCredits
}

func (p animeDocument) parseMainCharacters() []string {
	mainCharacters := []string{}
	p.Find("div.container div.g_section.main.character div.g_bubble.stripe.medium").Each(func(i int, e *goquery.Selection) {
		val, _ := e.Find("div.thumb.image.char a").Attr("href")
		mainCharacters = append(mainCharacters, val[11:])
	})
	return mainCharacters
}

func (p animeDocument) parseSecondaryCharacters() []string {
	secondaryCharacters := []string{}
	p.Find("div.container div.g_section.secondary.cast div.g_bubble.stripe.medium").Each(func(i int, e *goquery.Selection) {
		val, _ := e.Find("div.thumb.image.char a").Attr("href")
		secondaryCharacters = append(secondaryCharacters, val[11:])
	})
	return secondaryCharacters
}

func (p animeDocument) parseAppearsCharacters() []string {
	appearsCharacters := []string{}
	p.Find("div.container div.g_section.appears div.g_bubble.stripe.medium").Each(func(i int, e *goquery.Selection) {
		val, exists := e.Find("div.thumb.image.char a").Attr("href")
		if exists {
			appearsCharacters = append(appearsCharacters, val[11:])
		}
	})
	return appearsCharacters
}

func (p animeDocument) getEpisodes() []model.Episode {
	episodes := []model.Episode{}
	p.Find("div.g_section.container table#eplist tbody tr").Each(func(i int, e *goquery.Selection) {
		episode := model.Episode{}
		episode.Num = trimChars(e.Find("td.id a abbr").Text())
		idInt, _ := e.Find("td.id a").Attr("href")
		episode.Id = idInt[10:]
		episode.Title = trimChars(e.Find("td.title label").Text())
		episode.Duration = trimChars(e.Find("td.duration").Text())
		episode.AirDate = trimChars(e.Find("td.date.airdate").Text())
		if e.Find(" td.id a abbr[title*=\"Opening/Ending\"]").Text() == "" {
			episodes = append(episodes, episode)
		}
	})
	return episodes
}

func getStaff(el *goquery.Selection) model.Staff {
	idEl, _ := el.Find("a").Attr("href")
	id := idEl[9:]
	name := el.Find("a").Text()
	return model.Staff{id, name}
}

func getCredit(el *goquery.Selection) model.Credit {
	idEl, _ := el.Find("a").Attr("href")
	id := idEl[17:]
	name := el.Find("a").Text()
	return model.Credit{id, name}
}

func hasCredit(el *goquery.Selection) bool {
	idEl, _ := el.Find("td.credit a").Attr("href")
	return idEl != ""
}

func trimChars(str string) string {
	result := strings.ReplaceAll(str, "\t", "")
	result = strings.ReplaceAll(result, "\n", "")
	return result
}
