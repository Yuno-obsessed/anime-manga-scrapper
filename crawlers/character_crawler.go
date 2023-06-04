package crawlers

import (
	"crawler/config"
	"crawler/model"
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func CrawlCharacter(idInt int) (model.CharacterInfo, error) {
	crawlUrl := "https://anidb.net/character/"

	id := strconv.Itoa(idInt)

	response, err := config.PrepareRequest(crawlUrl + id)

	if response.StatusCode == 404 {
		return model.CharacterInfo{}, config.ErrNotFound
	} else if response.StatusCode != 200 {
		return model.CharacterInfo{}, fmt.Errorf("%v", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if doc.Find("div#layout-main h1.character").Text() == "ERROR" {
		log.Println("Got captcha")
		return model.CharacterInfo{}, config.ErrNotFound
	}

	out := model.CharacterInfo{
		Id: id,
	}

	img := doc.Find("div.image div.container picture img")
	out.ImageUrl, _ = img.Attr("src")

	desc := doc.Find("div.g_bubble.g_section.desc")
	out.Description = desc.Text()

	mainName := doc.Find("div.g_definitionlist tbody tr.g_odd.romaji.mainname td.value span")
	out.MainName = mainName.First().Text()

	offName := doc.Find("div.g_definitionlist tbody tr.official.verified.no td.value label")
	out.OfficialName = offName.First().Text()

	dateOfBirth := doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info table tr.birthdate td.value span")
	out.DateOfBirth = dateOfBirth.Text()

	// gender := doc.Find("div.g_definitionlist tbody tr.g_odd.gender td.value span")
	gender := doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.g_odd.gender td.value span")
	out.Gender = gender.Text()

	if out.Gender == "" {
		genderIdentity := doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.gender td.value span")
		out.Gender = genderIdentity.Text()
	}

	age := doc.Find("div.g_definitionlist tbody tr.age td.value")
	out.Age = age.First().Text()

	// height := doc.Find("div.g_definitionlist tbody tr.g_odd.height td.value span")
	// out.Height = height.Text()
	height := doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info table tr.height td.value span")
	out.Height = height.Text()

	weight := doc.Find("div.g_definitionlist tbody tr.weight td.value span")
	out.Weight = weight.First().Text()

	abilities := []string{}
	doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.abilities td.value span span.tagname").Each(func(i int, e *goquery.Selection) {
		abilities = append(abilities, e.Text())
	})
	out.Abilities = abilities

	abilitiesDesc := []string{}
	doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.g_odd.abilities td.value span span.wrapper span").Each(func(i int, e *goquery.Selection) {
		abilitiesDesc = append(abilitiesDesc, e.Text())
	})
	out.AbilitiesDesc = abilitiesDesc

	ageRange := doc.Find("div.g_definitionlist tbody tr.g_odd.age.range td.value span a span")
	out.AgeRange = ageRange.First().Text()
	out.AgeRangeDesc = ageRange.Last().Text()

	entities := []string{}
	doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.entity td.value span a span.tagname").Each(func(i int, e *goquery.Selection) {
		entities = append(entities, e.Text())
	})
	out.Entity = entities

	entityDescs := []string{}
	doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.entity td.value span a span.wrapper span").Each(func(i int, e *goquery.Selection) {
		entityDescs = append(entityDescs, e.Text())
	})
	out.EntityDesc = entityDescs

	roles := []string{}
	doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.role td.value span span.tagname").Each(func(i int, e *goquery.Selection) {
		roles = append(roles, e.Text())
	})
	out.Role = roles

	roleDescs := []string{}
	doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.g_section.info tbody tr.role td.value span span.wrapper span").Each(func(i int, e *goquery.Selection) {
		roleDescs = append(roleDescs, e.Text())
	})
	out.RoleDesc = roleDescs

	creatorIds := []string{}
	doc.Find("div.g_content.character_all.sidebar div.container.g_bubble table#seiyuulist_" + id + " tbody tr:not(.rowspan)").Each(func(i int, e *goquery.Selection) {
		creator, _ := e.Find("td.name.creator a").Attr("href")
		creator = creator[9:]
		creatorIds = append(creatorIds, creator)
	})
	out.CreatorIds = creatorIds

	animeIds := []string{}
	animeEpsOccur := []string{}
	doc.Find("div.g_content.character_all.sidebar div:not(.g_section.character_entry.hide) div.pane.anime_appearance table#animelist_" + id + " tbody tr").Each(func(i int, e *goquery.Selection) {
		anime, _ := e.Find("td.name.anime a").Attr("href")
		anime = anime[7:]
		animeIds = append(animeIds, anime)

		animeEps := e.Find("td.eprange").Text()
		animeEpsOccur = append(animeEpsOccur, animeEps)
	})
	out.AnimeIds = animeIds

	out.AnimeEpsOccur = animeEpsOccur
	fmt.Println(id)

	return out, nil
}
