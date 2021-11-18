package domain

import (
	"time"
)

type Competition struct {
	Id       uint32
	Name     string
	Code     string
	CodeName string
	Division string
	Logo     string
}

type Match struct {
	Id        uint32
	Season    string
	Round     string
	Group     string
	HomeId    uint32
	AwayId    uint32
	Stadium   string
	StartAt   time.Time
	HomeScore int
	AwayScore int
}

type Round struct {
	Id     uint32
	Name   string
	Number int
	Code   string
	Season Season
}

type Season struct {
	Id          uint32
	Name        string
	Code        string
	Competition Competition
}
