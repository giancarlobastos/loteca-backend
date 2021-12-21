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

type GetFixtureResponse struct {
	Response
	Results []FixtureResult `json:"response"`
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
	Id          uint32        `json:"id"`
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
	ElapsedMinutes uint8  `json:"elapsed"`
}

type FixtureScore struct {
	HalfTime  FixtureGoals `json:"halftime"`
	FullTime  FixtureGoals `json:"fulltime"`
	ExtraTime FixtureGoals `json:"extratime"`
	Penalty   FixtureGoals `json:"penalty"`
}

type FixtureGoals struct {
	Home uint8 `json:"home"`
	Away uint8 `json:"away"`
}

type FixtureLeague struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	LogoUrl string `json:"logo"`
	FlagUrl string `json:"flag"`
	Season  uint   `json:"season"`
	Round   string `json:"round"`
}

type FixtureTeams struct {
	Home FixtureTeam `json:"home"`
	Away FixtureTeam `json:"away"`
}

type FixtureTeam struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	LogoUrl string `json:"logo"`
	Winner  bool   `json:"winner"`
}
