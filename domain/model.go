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
	Ended    bool
}

type Match struct {
	Id        uint32
	Round     Round
	Group     Group
	Home      Team
	Away      Team
	Stadium   string
	StartAt   time.Time
	HomeScore int
	AwayScore int
}

type Round struct {
	Id          uint32
	Name        string
	Number      int
	Ended       bool
	Competition Competition
	Season      Season
}

type Season struct {
	Id   uint32
	Name string
	Code string
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
