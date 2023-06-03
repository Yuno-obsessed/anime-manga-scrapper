package crawlers

import (
	"crawler/config"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AnimeMain struct {
	Title    string
	Type     string
	Episodes string
	Picture  string
	URL      string
	Aired    string
	Ended    string
}

func (d AnimeMain) print() {
	fmt.Printf("\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n", d.Title, d.Type, d.Episodes, d.Picture, d.URL, d.Aired, d.Ended)
}

func CrawlList() []string {
	baseUrl := "https://anidb.net"
	crawlUrl := "https://anidb.net/anime/?h=1&noalias=1&orderby.name=0.1&page=0&view=list"

	animeUrlList := []string{}

	response, err := config.PrepareRequest(crawlUrl)

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	out := make([]AnimeMain, 0)

	doc.Find("div.animelist_list tbody tr").Each(func(_ int, tr *goquery.Selection) {
		td := tr.Find("td")
		if td.Length() <= 4 {
			fmt.Print("less")
			return
		}

		title := tr.Find("td.name.main.anime a").Text()
		picture, _ := tr.Find("picture source").Attr("srcset")
		url, _ := td.Find("a").Attr("href")
		typeA := strings.TrimSpace(tr.Find("td.type").Text())
		eps := strings.TrimSpace(tr.Find("td.count.eps").Text())
		aired := strings.TrimSpace(tr.Find("td.date.airdate").Text())
		ended := strings.TrimSpace(tr.Find("td.date.enddate").Text())

		out = append(out, AnimeMain{
			Title:    title,
			Type:     typeA,
			Episodes: eps,
			Picture:  picture,
			URL:      baseUrl + url,
			Aired:    aired,
			Ended:    ended,
		})
	})

	for _, result := range out {
		animeUrlList = append(animeUrlList, result.URL)
		result.print()
	}

	return animeUrlList
}
