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

type characterDocument struct {
	*goquery.Document
	id string
}

func NewCharacterDocument(body io.ReadCloser, id string) characterDocument {
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

func CrawlCharacter(idInt int) (*model.CharacterInfo, error) {
	crawlUrl := "https://anidb.net/character/"

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

	doc := NewCharacterDocument(response.Body, id)

	return &model.CharacterInfo{
		Id:            id,
		ImageUrl:      doc.parseImage(),
		Description:   doc.parseDescription(),
		MainName:      doc.parseMainName(),
		OfficialName:  doc.parseOfficialName(),
		DateOfBirth:   doc.parseDateOfBirth(),
		Age:           doc.parseAge(),
		Gender:        doc.parseGender(),
		Height:        doc.parseHeight(),
		Weight:        doc.parseWeight(),
		Abilities:     doc.parseAbilities(),
		AbilitiesDesc: doc.parseAbilitiesDesc(),
		AgeRange:      doc.parseAgeRange(),
		AgeRangeDesc:  doc.parseAgeRangeDesc(),
		Entities:      doc.parseEntities(),
		EntityDesc:    doc.parseEntitiesDesc(),
		Roles:         doc.parseRoles(),
		RolesDesc:     doc.parseRolesDesc(),
		CreatorIds:    doc.parseCreatorIds(),
		AnimeIds:      doc.parseAnimeIds(),
		AnimeEpsOccur: doc.parseAnimeEpsOccurrences(),
	}, nil
}

func (p characterDocument) parseImage() string {
	img := p.Find("div.image div.container picture img")
	imageUrl, _ := img.Attr("src")
	return imageUrl
}

func (p characterDocument) parseDescription() string {
	return trimChars(p.Find("div.g_bubble.g_section.desc").Text())
}

func (p characterDocument) parseMainName() string {
	return p.Find("div.g_definitionlist tbody tr.g_odd.romaji.mainname td.value span").First().Text()
}

func (p characterDocument) parseOfficialName() string {
	return p.Find("div.g_definitionlist tbody tr.official.verified.no td.value label").First().Text()
}

func (p characterDocument) parseDateOfBirth() string {
	return p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info table tr.birthdate td.value span").Text()

}

func (p characterDocument) parseGender() string {
	gender := p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.g_odd.gender td.value span").Text()
	if gender == "" {
		gender = p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.gender td.value span").Text()
	}
	return gender
}

func (p characterDocument) parseAge() string {
	return p.Find("div.g_definitionlist tbody tr.age td.value").First().Text()
}

func (p characterDocument) parseHeight() string {
	return p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info table tr.height td.value span").Text()
}

func (p characterDocument) parseWeight() string {
	return p.Find("div.g_definitionlist tbody tr.weight td.value span").First().Text()
}

func (p characterDocument) parseAbilities() []string {
	abilities := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.abilities td.value span span.tagname").Each(func(i int, e *goquery.Selection) {
		abilities = append(abilities, e.Text())
	})
	return abilities
}

func (p characterDocument) parseAbilitiesDesc() []string {
	abilitiesDesc := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.g_odd.abilities td.value span span.wrapper span").Each(func(i int, e *goquery.Selection) {
		abilitiesDesc = append(abilitiesDesc, e.Text())
	})
	return abilitiesDesc
}

func (p characterDocument) parseAgeRange() string {
	return p.Find("div.g_definitionlist tbody tr.g_odd.age.range td.value span a span").First().Text()
}

func (p characterDocument) parseAgeRangeDesc() string {
	return p.Find("div.g_definitionlist tbody tr.g_odd.age.range td.value span a span").Last().Text()
}

func (p characterDocument) parseEntities() []string {
	entities := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.entity td.value span a span.tagname").Each(func(i int, e *goquery.Selection) {
		entities = append(entities, e.Text())
	})
	return entities
}

func (p characterDocument) parseEntitiesDesc() []string {
	entityDescs := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.entity td.value span a span.wrapper span").Each(func(i int, e *goquery.Selection) {
		entityDescs = append(entityDescs, e.Text())
	})
	return entityDescs
}

func (p characterDocument) parseRoles() []string {
	roles := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.role td.value span span.tagname").Each(func(i int, e *goquery.Selection) {
		roles = append(roles, e.Text())
	})
	return roles
}

func (p characterDocument) parseRolesDesc() []string {
	roleDescs := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.role td.value span span.wrapper span").Each(func(i int, e *goquery.Selection) {
		roleDescs = append(roleDescs, e.Text())
	})
	return roleDescs
}

func (p characterDocument) parseCreatorIds() []string {
	creatorIds := []string{}
	p.Find("div.g_content.character_all.sidebar div.container.g_bubble table#seiyuulist_" + p.id + " tbody tr:not(.rowspan)").Each(func(i int, e *goquery.Selection) {
		creator, _ := e.Find("td.name.creator a").Attr("href")
		creator = creator[9:]
		creatorIds = append(creatorIds, creator)
	})
	return creatorIds
}

func (p characterDocument) parseAnimeIds() []string {
	animeIds := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.pane.anime_appearance table#animelist_" + p.id + " tbody tr").Each(func(i int, e *goquery.Selection) {
		anime, _ := e.Find("td.name.anime a").Attr("href")
		anime = anime[7:]
		animeIds = append(animeIds, anime)
	})
	return animeIds
}

func (p characterDocument) parseAnimeEpsOccurrences() []string {
	animeEpsOccur := []string{}
	p.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.pane.anime_appearance table#animelist_" + p.id + " tbody tr").Each(func(i int, e *goquery.Selection) {
		animeEps := e.Find("td.eprange").Text()
		animeEpsOccur = append(animeEpsOccur, animeEps)
	})
	return animeEpsOccur
}
