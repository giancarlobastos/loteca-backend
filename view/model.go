package view

import (
	"time"
)

type Lottery struct {
	Id               *int       `json:"id,omitempty"`
	Number           *int       `json:"number,omitempty"`
	Name             *string    `json:"name,omitempty"`
	EstimatedPrize   *float32   `json:"estimated_prize,omitempty"`
	MainPrize        *float32   `json:"main_prize,omitempty"`
	MainPrizeWinners *int       `json:"main_prize_winners,omitempty"`
	SidePrize        *float32   `json:"side_prize,omitempty"`
	SidePrizeWinners *int       `json:"side_prize_winners,omitempty"`
	SpecialPrize     *float32   `json:"special_prize,omitempty"`
	Accumulated      *bool      `json:"accumulated,omitempty"`
	EndAt            *time.Time `json:"end_at,omitempty"`
	ResultAt         *time.Time `json:"result_at,omitempty"`
	EarliestMatchAt  *time.Time `json:"earliest_at,omitempty"`
	LatestMatchAt    *time.Time `json:"latest_at,omitempty"`
	Matches          *[]Match   `json:"matches,omitempty"`
	LotteryIds       *[]int     `json:"lottery_ids,omitempty"`
	UpdatedAt        *time.Time `json:"-"`
	Enabled        bool `json:"-"`
}

type Match struct {
	Id               *int       `json:"id,omitempty"`
	HomeId           *int       `json:"home_id,omitempty"`
	HomeName         *string    `json:"home_name,omitempty"`
	HomeAbbreviation *string    `json:"home_abbr,omitempty"`
	HomeLogo         *string    `json:"home_logo,omitempty"`
	AwayId           *int       `json:"away_id,omitempty"`
	AwayName         *string    `json:"away_name,omitempty"`
	AwayAbbreviation *string    `json:"away_abbr,omitempty"`
	AwayLogo         *string    `json:"away_logo,omitempty"`
	Stadium          *string    `json:"stadium,omitempty"`
	StartAt          *time.Time `json:"start_at,omitempty"`
	HomeScore        *int       `json:"home_score,omitempty"`
	AwayScore        *int       `json:"away_score,omitempty"`
	RoundNumber      *int       `json:"round_number,omitempty"`
	RoundName        *string    `json:"round_name,omitempty"`
	CompetitionId    *int       `json:"competition_id,omitempty"`
	CompetitionName  *string    `json:"competition_name,omitempty"`
	Year             *int       `json:"year,omitempty"`
	Order            *int       `json:"order,omitempty"`
	Raffle           bool       `json:"raffle"`
	RaffleResult     *string    `json:"raffle_result,omitempty"`
	Ended            bool       `json:"ended"`
	Status           *string    `json:"status,omitempty"`
	ElapsedTime      *int       `json:"elapsed_time,omitempty"`
}

type MatchDetails struct {
	Id                         *int         `json:"id"`
	Match                      *Match       `json:"match"`
	Stats                      *[]TeamStats `json:"stats"`
	StatsHome                  *TeamStats   `json:"stats_home"`
	StatsAway                  *TeamStats   `json:"stats_away"`
	H2H                        *[]Match     `json:"h2h"`
	LastMatchesHome            *[]Match     `json:"last_matches_home"`
	LastMatchesAway            *[]Match     `json:"last_matches_away"`
	LastMatchesHomeCompetition *[]Match     `json:"last_matches_home_competition"`
	LastMatchesAwayCompetition *[]Match     `json:"last_matches_away_competition"`
	LastMatchesAtHome          *[]Match     `json:"last_matches_at_home"`
	LastMatchesAtAway          *[]Match     `json:"last_matches_at_away"`
	NextMatchesHome            *[]Match     `json:"next_matches_home"`
	NextMatchesAway            *[]Match     `json:"next_matches_away"`
	Odds                       *[]Odd       `json:"odds"`
	Votes                      *Vote        `json:"votes"`
	TotalVotes                 int          `json:"total_votes"`
}

type PollResults struct {
	LotteryId  int    `json:"lottery_id"`
	Votes      []Vote `json:"votes"`
	TotalVotes int    `json:"total_votes"`
}

type Vote struct {
	MatchId   int `json:"match_id"`
	HomeVotes int `json:"home_votes"`
	DrawVotes int `json:"draw_votes"`
	AwayVotes int `json:"away_votes"`
}

type TeamStats struct {
	M  int `json:"m"`
	W  int `json:"w"`
	D  int `json:"d"`
	L  int `json:"l"`
	GP int `json:"gp"`
	GC int `json:"gc"`
	SG int `json:"sg"`
}

type Odd struct {
	BookmakerId   int     `json:"bookmaker_id"`
	BookmakerName string  `json:"bookmaker_name"`
	Home          float32 `json:"home"`
	Draw          float32 `json:"draw"`
	Away          float32 `json:"away"`
}

type LiveScore struct {
	Id        *int       `json:"id,omitempty"`
	Order     *int       `json:"order,omitempty"`
	StartAt   *time.Time `json:"start_at,omitempty"`
	HomeScore *int       `json:"home_score,omitempty"`
	AwayScore *int       `json:"away_score,omitempty"`
	Live      *bool      `json:"live,omitempty"`
	HalfTime  *bool      `json:"half_time,omitempty"`
	Status    *string    `json:"status,omitempty"`
	Elapsed   *int       `json:"elapsed,omitempty"`
}
