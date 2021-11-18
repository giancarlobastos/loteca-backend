package domain

type Tournament struct {
	Id       int
	Name     string `json:"name"`
	Division string `json:"division"`
	Logo     string `json:"logo"`
}

type Match struct {
	Id        int   `json:"id"`
	RoundId   int   `json:"roundId"`
	HomeId    int   `json:"homeId"`
	AwayId    int   `json:"awayId"`
	StadiumId int   `json:"stadiumId"`
	StartAt   int64 `json:"startAt"`
	HomeScore int   `json:"homeScore"`
	AwayScore int   `json:"awayScore"`
}
