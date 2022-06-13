package client

import (
	"encoding/json"
	// "errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type ApiFootballClient struct {
	remainingRequestCount int
}

func NewApiFootballClient() *ApiFootballClient {
	return &ApiFootballClient{
		remainingRequestCount: 100,
	}
}

func (c *ApiFootballClient) GetTeams(country string) (teamsResponse *GetTeamsResponse, err error) {
	params := map[string]string{"country": country}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/teams", &params)
	json.Unmarshal(body, &teamsResponse)
	return teamsResponse, nil
}

func (c *ApiFootballClient) GetLeagues(country string) (leaguesResponse *GetLeaguesResponse, err error) {
	params := map[string]string{"country": country}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/leagues", &params)
	json.Unmarshal(body, &leaguesResponse)
	return leaguesResponse, nil
}

func (c *ApiFootballClient) GetFixtures(leagueId int, year int) (fixturesResponse *GetFixturesResponse, err error) {
	params := map[string]string{"league": strconv.Itoa(leagueId), "season": strconv.Itoa(year), "timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures", &params)
	json.Unmarshal(body, &fixturesResponse)
	return fixturesResponse, nil
}

func (c *ApiFootballClient) GetFixturesById(matchIds []int) (fixturesResponse *GetFixturesResponse, err error) {
	ids := ""

	for i, matchId := range matchIds {
		ids += strconv.Itoa(matchId)

		if i < len(matchIds) - 1 {
			ids += "-"
		}
	}

	params := map[string]string{"ids": ids, "timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures", &params)
	json.Unmarshal(body, &fixturesResponse)
	return fixturesResponse, nil
}

func (c *ApiFootballClient) GetLastFixtures(teamId int, limit int) (fixturesResponse *GetFixturesResponse, err error) {
	params := map[string]string{"team": strconv.Itoa(teamId), "last": strconv.Itoa(limit), "timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures", &params)
	json.Unmarshal(body, &fixturesResponse)
	return fixturesResponse, nil
}

func (c *ApiFootballClient) GetLastCompetitionFixtures(leagueId int, year int, teamId int, limit int) (fixturesResponse *GetFixturesResponse, err error) {
	params := map[string]string{"league": strconv.Itoa(leagueId), "season": strconv.Itoa(year), "team": strconv.Itoa(teamId), "last": strconv.Itoa(limit), "timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures", &params)
	json.Unmarshal(body, &fixturesResponse)
	return fixturesResponse, nil
}

func (c *ApiFootballClient) GetNextFixtures(teamId int, limit int) (fixturesResponse *GetFixturesResponse, err error) {
	params := map[string]string{"team": strconv.Itoa(teamId), "next": strconv.Itoa(limit), "timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures", &params)
	json.Unmarshal(body, &fixturesResponse)
	return fixturesResponse, nil
}

func (c *ApiFootballClient) GetNextCompetitionFixtures(leagueId int, year int, teamId int, limit int) (fixturesResponse *GetFixturesResponse, err error) {
	params := map[string]string{"league": strconv.Itoa(leagueId), "season": strconv.Itoa(year), "team": strconv.Itoa(teamId), "next": strconv.Itoa(limit), "timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures", &params)
	json.Unmarshal(body, &fixturesResponse)
	return fixturesResponse, nil
}

func (c *ApiFootballClient) GetHeadToHead(homeId int, awayId int) (fixturesResponse *GetFixturesResponse, err error) {
	params := map[string]string{
		"h2h":      strconv.Itoa(homeId) + "-" + strconv.Itoa(awayId),
		"last":     "5",
		"timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures/headtohead", &params)
	json.Unmarshal(body, &fixturesResponse)
	return fixturesResponse, nil
}

func (c *ApiFootballClient) GetOdds(matchId int) (oddsResponse *GetOddsResponse, err error) {
	params := map[string]string{
		"fixture":  strconv.Itoa(matchId),
		"bet":      "1",
		"timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/odds", &params)
	json.Unmarshal(body, &oddsResponse)
	return oddsResponse, nil
}

func (c *ApiFootballClient) callApi(url string, params *map[string]string) ([]byte, error) {
	log.Printf("[%s]: %v", url, params)

	if c.remainingRequestCount <= 0 {
		log.Printf("no remaining calls to ApiFootball [%v]", c.remainingRequestCount)
		// return make([]byte, 0), errors.New("no remaining calls to ApiFootball")
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "dbcf853914msh8b02e0bab593853p14026djsne2424dfc2ab8")

	query := req.URL.Query()

	for key, value := range *params {
		query.Add(key, value)
	}

	req.URL.RawQuery = query.Encode()
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Error: %v - [%s]: %v", err, url, params)
		return make([]byte, 0), err
	}

	defer res.Body.Close()

	c.remainingRequestCount, _ = strconv.Atoi(res.Header.Get("x-ratelimit-requests-remaining"))
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
