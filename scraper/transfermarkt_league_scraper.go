package scraper

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/nleeper/goment"
	"golang.org/x/net/html"
)

const ()

type TransferMarktLeagueScraper struct {
}

func NewTransferMarktLeagueScraper() *TransferMarktLeagueScraper {
	return &TransferMarktLeagueScraper{}
}

func (c *TransferMarktLeagueScraper) GetMatchList(round domain.Round) (*[]domain.Match, error) {
	url := fmt.Sprintf("https://www.transfermarkt.com.br/%s/spieltag/wettbewerb/%s/saison_id/%s/spieltag/%s",
		round.Season.Competition.CodeName,
		round.Season.Competition.Code,
		round.Season.Code,
		round.Code)

	doc := c.getDocument(url)
	matches := make([]domain.Match, 10)

	doc.Find("div.footer a[href*='bericht/index/spielbericht/']").Each(
		func(i int, s *goquery.Selection) {
			matchDetailsUrl, _ := s.Attr("href")
			match, _ := c.getMatchDetails(matchDetailsUrl)
			matches = append(matches, *match)
		})

	return &matches, nil
}

func (c *TransferMarktLeagueScraper) getMatchDetails(url string) (*domain.Match, error) {
	doc := c.getDocument("https://www.transfermarkt.com.br" + url)
	match := domain.Match{
		HomeScore: -1,
		AwayScore: -1,
	}

	urlSplit := strings.Split(url, "/")
	matchId, _ := strconv.ParseUint(urlSplit[len(urlSplit)-1], 10, 32)
	match.Id = uint32(matchId)

	teamUrls := doc.Find("div.sb-team  a.sb-vereinslink").Nodes

	homeTeamUrl := c.attr(teamUrls[0], "href")
	homeTeamUrlSplit := strings.Split(homeTeamUrl, "/")
	homeTeamId, _ := strconv.ParseUint(homeTeamUrlSplit[4], 10, 32)
	match.HomeId = uint32(homeTeamId)

	awayTeamUrl := c.attr(teamUrls[1], "href")
	awayTeamUrlSplit := strings.Split(awayTeamUrl, "/")
	awayTeamId, _ := strconv.ParseUint(awayTeamUrlSplit[4], 10, 32)
	match.AwayId = uint32(awayTeamId)

	scorePattern := regexp.MustCompile("[^0-9]*([0-9]+):([0-9]+).*")
	score := doc.Find(".sb-endstand").Text()
	hasScore := scorePattern.MatchString(score)

	if hasScore {
		match.HomeScore, _ = strconv.Atoi(scorePattern.FindStringSubmatch(score)[1])
		match.AwayScore, _ = strconv.Atoi(scorePattern.FindStringSubmatch(score)[2])
	}

	datePattern := regexp.MustCompile(`[^0-9]*([0-9]{2}/[0-9]{2}/[0-9]{2})[^0-9]*([0-9]{2}:[0-9]{2})`)
	dateSplit := datePattern.FindStringSubmatch(doc.Find(".sb-datum").Text())
	formatter, _ := goment.New(dateSplit[1]+dateSplit[2], "DD/MM/YYHH:mm")

	match.StartAt = formatter.ToTime()
	match.Stadium = doc.Find(".sb-zusatzinfos span a").Text()

	return &match, nil
}

func (c *TransferMarktLeagueScraper) getDocument(url string) *goquery.Document {
	fmt.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36")

	res, _ := client.Do(req)
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	return doc
}

func (c *TransferMarktLeagueScraper) attr(n *html.Node, key string) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == key {
				return a.Val
			}
		}
	}
	return ""
}
