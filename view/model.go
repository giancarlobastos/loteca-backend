package view

import (
	"time"
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

type Match struct {
	Id              *int       `json:",omitempty"`
	HomeId          *int       `json:",omitempty"`
	HomeName        *string    `json:",omitempty"`
	AwayId          *int       `json:",omitempty"`
	AwayName        *string    `json:",omitempty"`
	Stadium         *string    `json:",omitempty"`
	StartAt         *time.Time `json:",omitempty"`
	HomeScore       *int       `json:",omitempty"`
	AwayScore       *int       `json:",omitempty"`
	RoundNumber     *int       `json:",omitempty"`
	RoundName       *string    `json:",omitempty"`
	CompetitionId   *int       `json:",omitempty"`
	CompetitionName *string    `json:",omitempty"`
	Year            *int       `json:",omitempty"`
	Order           *int       `json:",omitempty"`
}
