package crawler

import (
	"fmt"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/gocolly/colly"
)

type Crawler struct {
}

func NewCrawler() *Crawler {
	return &Crawler{}
}

func (c *Crawler) GetMatch(matchId int) (*domain.Match, error) {
	col := colly.NewCollector(colly.AllowedDomains("www.playmakerstats.com"))
	col.OnHTML("tr.parent", func(e *colly.HTMLElement) {
		id := e.Attr("id")
		// Print link
		fmt.Printf("tr found: %s\n", id)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//col.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	col.Visit(fmt.Sprint("https://www.playmakerstats.com/edition_matches.php?id=", matchId))
	return &domain.Match{}, nil
}
