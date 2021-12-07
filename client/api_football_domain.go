package client

type GetTeamResponse struct {
	Get        string            `json:"get"`
	Parameters map[string]string `json:"parameters"`
	Errors     map[string]string `json:"errors"`
	Size       int               `json:"results"`
	Paging     Paging            `json:"paging"`
	Results    []TeamResult      `json:"response"`
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
