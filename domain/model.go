package domain

import (
	"time"
)

type Team struct {
	Id      int      `json:",omitempty"`
	Name    string   `json:",omitempty"`
	Logo    string   `json:",omitempty"`
	Country string   `json:",omitempty"`
	Stadium *Stadium `json:",omitempty"`
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
	Matches          *[]Match   `json:",omitempty"`
}
