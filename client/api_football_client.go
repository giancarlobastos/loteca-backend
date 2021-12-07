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
	body, _ := c.callApi("https://api-football-v1.p.rapidapi.com/v3/teams?country=" + country)
	json.Unmarshal(body, &teamResponse)
	return teamResponse, nil
}

func (c *ApiFootballClient) callApi(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "dbcf853914msh8b02e0bab593853p14026djsne2424dfc2ab8")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return make([]byte, 0), err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
