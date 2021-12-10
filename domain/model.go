package domain

import (
	"time"
)

type Team struct {
	Id      uint32
	Name    string
	Logo    string
	Country string
}

type Competition struct {
	Id       uint32
	Name     string
	Division string
	Logo     string
	Type     CompetitionType
	Country  string
}

type Match struct {
	Id        uint32
	Round     *Round
	Group     *Group
	Home      *Team
	Away      *Team
	Stadium   string
	StartAt   time.Time
	HomeScore uint
	AwayScore uint
}

type Round struct {
	Id          uint32
	Name        string
	Number      uint
	Ended       bool
	Competition *Competition
	Year        uint
}

type Season struct {
	Competition *Competition
	Year        uint
	Name        string
	Ended       bool
}

type Group struct {
	Id   uint32
	Name string
}

type Stadium struct {
	Id      uint32
	Name    string
	City    string
	State   string
	Country string
}

type CompetitionType string

const (
	CUP    CompetitionType = "Cup"
	LEAGUE CompetitionType = "League"
)
