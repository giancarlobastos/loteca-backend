package client

type Response struct {
	Get        string            `json:"get"`
	Parameters map[string]string `json:"parameters"`
	Errors     map[string]string `json:"errors"`
	Size       int               `json:"results"`
	Paging     Paging            `json:"paging"`
}

type GetTeamsResponse struct {
	Response
	Results []TeamResult `json:"response"`
}

type GetLeaguesResponse struct {
	Response
	Results []LeagueResult `json:"response"`
}

type GetFixturesResponse struct {
	Response
	Results []FixtureResult `json:"response"`
}

type GetOddsResponse struct {
	Response
	Results []OddsResult `json:"response"`
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
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	LogoUrl  string `json:"logo"`
}

type Venue struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Capacity int    `json:"capacity"`
	Surface  string `json:"surface"`
	ImageUrl string `json:"image"`
}

type LeagueResult struct {
	League  League   `json:"league"`
	Country Country  `json:"country"`
	Seasons []Season `json:"seasons"`
}

type League struct {
	Id      int    `json:"id"`
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
	Year      int      `json:"year"`
	StartDate string   `json:"start"`
	EndDate   string   `json:"end"`
	Current   bool     `json:"current"`
	Coverage  Coverage `json:"coverage"`
}

type Coverage struct {
	AvailableFixtures AvailableFixtures `json:"fixtures"`
	Standings         bool              `json:"standings"`
	Players           bool              `json:"players"`
	TopScorers        bool              `json:"top_scorers"`
	Predictions       bool              `json:"predictions"`
	Odds              bool              `json:"odds"`
}

type AvailableFixtures struct {
	Events             bool `json:"events"`
	Lineups            bool `json:"lineups"`
	StatisticsFixtures bool `json:"statistics_fixtures"`
	StatisticsPlayers  bool `json:"statistics_players"`
}

type FixtureResult struct {
	Fixture Fixture       `json:"fixture"`
	League  FixtureLeague `json:"league"`
	Teams   FixtureTeams  `json:"teams"`
	Goals   FixtureGoals  `json:"goals"`
	Score   FixtureScore  `json:"score"`
}

type Fixture struct {
	Id          int           `json:"id"`
	Referee     string        `json:"referee"`
	Timezone    string        `json:"timezone"`
	DateAndTime string        `json:"date"`
	Timestamp   uint64        `json:"timestamp"`
	Period      FixturePeriod `json:"periods"`
	Venue       Venue         `json:"venue"`
	Status      FixtureStatus `json:"status"`
}

type FixturePeriod struct {
	First  uint64 `json:"first"`
	Second uint64 `json:"second"`
}

type FixtureStatus struct {
	Name           string `json:"long"`
	Code           string `json:"short"`
	ElapsedMinutes int    `json:"elapsed"`
}

type FixtureScore struct {
	HalfTime  FixtureGoals `json:"halftime"`
	FullTime  FixtureGoals `json:"fulltime"`
	ExtraTime FixtureGoals `json:"extratime"`
	Penalty   FixtureGoals `json:"penalty"`
}

type FixtureGoals struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type FixtureLeague struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	LogoUrl string `json:"logo"`
	FlagUrl string `json:"flag"`
	Season  int    `json:"season"`
	Round   string `json:"round"`
}

type FixtureTeams struct {
	Home FixtureTeam `json:"home"`
	Away FixtureTeam `json:"away"`
}

type FixtureTeam struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	LogoUrl string `json:"logo"`
	Winner  bool   `json:"winner"`
}

type OddsResult struct {
	Fixture    Fixture       `json:"fixture"`
	League     FixtureLeague `json:"league"`
	UpdatedAt  string        `json:"update"`
	Bookmakers []Bookmaker   `json:"bookmakers"`
}

type Bookmaker struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bets []Bet  `json:"bets"`
}

type Bet struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Odds []Odd  `json:"values"`
}

type Odd struct {
	Name  string `json:"value"`
	Value string `json:"odd"`
}
