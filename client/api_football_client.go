package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiFootballClient struct {
}

func NewApiFootballClient() *ApiFootballClient {
	return &ApiFootballClient{}
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

func (c *ApiFootballClient) callApi(url string, params *map[string]string) ([]byte, error) {
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
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
