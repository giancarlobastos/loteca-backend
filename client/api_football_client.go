package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func (c *ApiFootballClient) GetTeams(country string) (teamResponse *GetTeamResponse, err error) {
	params := map[string]string{"country": country}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/teams", &params)
	json.Unmarshal(body, &teamResponse)
	return teamResponse, nil
}

func (c *ApiFootballClient) GetLeagues(country string) (leagueResponse *GetLeagueResponse, err error) {
	params := map[string]string{"country": country}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/leagues", &params)
	json.Unmarshal(body, &leagueResponse)
	return leagueResponse, nil
}

func (c *ApiFootballClient) GetFixtures(leagueId uint32, year uint) (fixtureResponse *GetFixtureResponse, err error) {
	params := map[string]string{"league": strconv.Itoa(int(leagueId)), "season": strconv.Itoa(int(year)), "timezone": "America/Sao_Paulo"}
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/fixtures", &params)
	json.Unmarshal(body, &fixtureResponse)
	return fixtureResponse, nil
}

func (c *ApiFootballClient) callApi(url string, params *map[string]string) ([]byte, error) {
	if c.remainingRequestCount <= 0 {
		return make([]byte, 0), errors.New("no remaining calls to ApiFootball")
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
		return make([]byte, 0), err
	}

	defer res.Body.Close()

	c.remainingRequestCount, _ = strconv.Atoi(res.Header.Get("x-ratelimit-requests-remaining"))
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
