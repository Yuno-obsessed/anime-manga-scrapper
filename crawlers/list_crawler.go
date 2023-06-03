package crawlers

import (
	"fmt"
	"log"
	"net/http"
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

	request, err := http.NewRequest("GET", crawlUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.57")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)

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
