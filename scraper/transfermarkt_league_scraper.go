package scraper

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/gocolly/colly"
)

const ()

type TransferMarktLeagueScraper struct {
}

func NewTransferMarktLeagueScraper() *TransferMarktLeagueScraper {
	return &TransferMarktLeagueScraper{}
}

func (c *TransferMarktLeagueScraper) GetMatchList(round domain.Round) (*[]domain.Match, error) {
	matchListCollector := colly.NewCollector(colly.AllowedDomains("www.transfermarkt.com.br"))

	matches := make([]domain.Match, 10)

	matchListCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36")
		fmt.Println("Visiting", r.URL.String())
	})

	matchDetailsCollector := matchListCollector.Clone()

	matchListCollector.OnHTML(".liveLink", func(e *colly.HTMLElement) {
		match, _ := c.getMatchDetails(matchDetailsCollector, e.Attr("href"))

		if match != nil {
			matches = append(matches, *match)
		}
	})

	matchListCollector.Visit(fmt.Sprintf("https://www.transfermarkt.com.br/%s/spieltag/wettbewerb/%s/saison_id/%s/spieltag/%s",
		round.Season.Competition.CodeName,
		round.Season.Competition.Code,
		round.Season.Code,
		round.Code))

	return &matches, nil
}

func (c *TransferMarktLeagueScraper) getMatchDetails(collector *colly.Collector, url string) (*domain.Match, error) {
	match := domain.Match{}

	urlSplit := strings.Split(url, "/")
	matchId, _ := strconv.ParseUint(urlSplit[len(urlSplit)-1], 10, 32)
	match.Id = uint32(matchId)

	collector.OnHTML("div.sb-team a.sb-vereinslink", func(e *colly.HTMLElement) {
		homeTeam, _ := e.DOM.Eq(0).Attr("href")
		homeId, _ := strconv.ParseUint(homeTeam, 10, 32)
		//.attr("href").replaceFirst("[^0-9]*", "").split("/")[0];
		match.HomeId = uint32(homeId)
	})

	collector.Visit("https://www.transfermarkt.com.br/" + url)
	return &match, nil
}
