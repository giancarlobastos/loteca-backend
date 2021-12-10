package client

type Response struct {
	Get        string            `json:"get"`
	Parameters map[string]string `json:"parameters"`
	Errors     map[string]string `json:"errors"`
	Size       int               `json:"results"`
	Paging     Paging            `json:"paging"`
}

type GetTeamResponse struct {
	Response
	Results []TeamResult `json:"response"`
}

type GetLeagueResponse struct {
	Response
	Results []LeagueResult `json:"response"`
}
type Paging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type TeamResult struct {
	Team  Team  `json:"team"`
	Venue Venue `json:"venue"`
}

type Team struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Founded  uint   `json:"founded"`
	National bool   `json:"national"`
	LogoUrl  string `json:"logo"`
}

type Venue struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Capacity uint   `json:"capacity"`
	Surface  string `json:"surface"`
	ImageUrl string `json:"image"`
}

type LeagueResult struct {
	League  League   `json:"league"`
	Country Country  `json:"country"`
	Seasons []Season `json:"seasons"`
}

type League struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	LogoUrl string `json:"logo"`
}

type Country struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	FlagUrl string `json:"flag"`
}

type Season struct {
	Year      uint     `json:"year"`
	StartDate string   `json:"start"`
	EndDate   string   `json:"end"`
	Current   bool     `json:"current"`
	Coverage  Coverage `json:"coverage"`
}

type Coverage struct {
	Fixtures    Fixtures `json:"fixtures"`
	Standings   bool     `json:"standings"`
	Players     bool     `json:"players"`
	TopScorers  bool     `json:"top_scorers"`
	Predictions bool     `json:"predictions"`
	Odds        bool     `json:"odds"`
}

type Fixtures struct {
	Events             bool `json:"events"`
	Lineups            bool `json:"lineups"`
	StatisticsFixtures bool `json:"statistics_fixtures"`
	StatisticsPlayers  bool `json:"statistics_players"`
}
