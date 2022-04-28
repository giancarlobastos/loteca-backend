package domain

import (
	"time"
)

type Team struct {
	Id           int      `json:",omitempty"`
	Name         string   `json:",omitempty"`
	Abbreviation string   `json:",omitempty"`
	Logo         string   `json:",omitempty"`
	Country      string   `json:",omitempty"`
	Stadium      *Stadium `json:",omitempty"`
}

type Competition struct {
	Id       int             `json:",omitempty"`
	Name     string          `json:",omitempty"`
	Division string          `json:",omitempty"`
	Logo     string          `json:",omitempty"`
	Type     CompetitionType `json:",omitempty"`
	Country  string          `json:",omitempty"`
	Seasons  *[]Season       `json:",omitempty"`
}

type Match struct {
	Id        int       `json:",omitempty"`
	Home      *Team     `json:",omitempty"`
	Away      *Team     `json:",omitempty"`
	Stadium   *Stadium  `json:",omitempty"`
	StartAt   time.Time `json:",omitempty"`
	HomeScore *int      `json:",omitempty"`
	AwayScore *int      `json:",omitempty"`
	Order     int       `json:",omitempty"`
}

type Round struct {
	Id      int      `json:",omitempty"`
	Name    string   `json:",omitempty"`
	Number  int      `json:",omitempty"`
	Ended   bool     `json:",omitempty"`
	Matches *[]Match `json:",omitempty"`
}

type Season struct {
	Year   int      `json:",omitempty"`
	Name   string   `json:",omitempty"`
	Ended  bool     `json:",omitempty"`
	Rounds *[]Round `json:",omitempty"`
}

type Group struct {
	Id   int    `json:",omitempty"`
	Name string `json:",omitempty"`
}

type Stadium struct {
	Id      int    `json:",omitempty"`
	Name    string `json:",omitempty"`
	City    string `json:",omitempty"`
	State   string `json:",omitempty"`
	Country string `json:",omitempty"`
}

type CompetitionType string

const (
	CUP    CompetitionType = "Cup"
	LEAGUE CompetitionType = "League"
)

type Lottery struct {
	Id               int        `json:",omitempty"`
	Number           int        `json:",omitempty"`
	Name             string     `json:",omitempty"`
	EstimatedPrize   float32    `json:",omitempty"`
	MainPrize        float32    `json:",omitempty"`
	MainPrizeWinners int        `json:",omitempty"`
	SidePrize        float32    `json:",omitempty"`
	SidePrizeWinners int        `json:",omitempty"`
	SpecialPrize     float32    `json:",omitempty"`
	Accumulated      bool       `json:",omitempty"`
	EndAt            *time.Time `json:",omitempty"`
	ResultAt         *time.Time `json:",omitempty"`
	Matches          *[]Match   `json:",omitempty"`
}

type Bookmaker struct {
	Id   int    `json:",omitempty"`
	Name string `json:",omitempty"`
	Url  string `json:",omitempty"`
}

type Odd struct {
	Id        int        `json:",omitempty"`
	Bookmaker Bookmaker  `json:",omitempty"`
	Home      float32    `json:",omitempty"`
	Draw      float32    `json:",omitempty"`
	Away      float32    `json:",omitempty"`
	UpdatedAt *time.Time `json:",omitempty"`
}

type User struct {
	Id         *int   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	FacebookId string `json:"facebook_id,omitempty"`
	Email      string `json:"email,omitempty"`
	Picture    string `json:"picture,omitempty"`
}

type Poll struct {
	LotteryId int    `json:"lottery_id,omitempty"`
	Votes     []Vote `json:"votes,omitempty"`
}

type Vote struct {
	MatchId int    `json:"match_id,omitempty"`
	Result  string `json:"result,omitempty"`
}
