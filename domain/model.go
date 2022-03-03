package domain

import (
	"time"
)

type Team struct {
	Id      uint32   `json:",omitempty"`
	Name    string   `json:",omitempty"`
	Logo    string   `json:",omitempty"`
	Country string   `json:",omitempty"`
	Stadium *Stadium `json:",omitempty"`
}

type Competition struct {
	Id       uint32          `json:",omitempty"`
	Name     string          `json:",omitempty"`
	Division string          `json:",omitempty"`
	Logo     string          `json:",omitempty"`
	Type     CompetitionType `json:",omitempty"`
	Country  string          `json:",omitempty"`
	Seasons  *[]Season       `json:",omitempty"`
}

type Match struct {
	Id        uint32    `json:",omitempty"`
	Home      *Team     `json:",omitempty"`
	Away      *Team     `json:",omitempty"`
	Stadium   string    `json:",omitempty"`
	StartAt   time.Time `json:",omitempty"`
	HomeScore uint      `json:",omitempty"`
	AwayScore uint      `json:",omitempty"`
}

type Round struct {
	Id      uint32   `json:",omitempty"`
	Name    string   `json:",omitempty"`
	Number  uint     `json:",omitempty"`
	Ended   bool     `json:",omitempty"`
	Matches *[]Match `json:",omitempty"`
}

type Season struct {
	Year   uint     `json:",omitempty"`
	Name   string   `json:",omitempty"`
	Ended  bool     `json:",omitempty"`
	Rounds *[]Round `json:",omitempty"`
}

type Group struct {
	Id   uint32 `json:",omitempty"`
	Name string `json:",omitempty"`
}

type Stadium struct {
	Id      uint32 `json:",omitempty"`
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
	Number           int             `json:",omitempty"`
	EstimatedPrize   float32         `json:",omitempty"`
	MainPrize        float32         `json:",omitempty"`
	MainPrizeWinners int             `json:",omitempty"`
	SidePrize        float32         `json:",omitempty"`
	SidePrizeWinners int             `json:",omitempty"`
	SpecialPrize     float32         `json:",omitempty"`
	Accumulated      bool            `json:",omitempty"`
	EndAt            time.Time       `json:",omitempty"`
	Matches          *[]LotteryMatch `json:",omitempty"`
}

type LotteryMatch struct {
	Order       int          `json:",omitempty"`
	Match       *Match       `json:",omitempty"`
	Round       *Round       `json:",omitempty"`
	Competition *Competition `json:",omitempty"`
}

type MatchVO struct {
	Id              uint32    `json:",omitempty"`
	HomeId          uint32    `json:",omitempty"`
	HomeName        string    `json:",omitempty"`
	AwayId          uint32    `json:",omitempty"`
	AwayName        string    `json:",omitempty"`
	Stadium         string    `json:",omitempty"`
	StartAt         time.Time `json:",omitempty"`
	HomeScore       uint      `json:",omitempty"`
	AwayScore       uint      `json:",omitempty"`
	RoundNumber     uint      `json:",omitempty"`
	RoundName       string    `json:",omitempty"`
	CompetitionId   uint32    `json:",omitempty"`
	CompetitionName string    `json:",omitempty"`
	Year            uint      `json:",omitempty"`
}
