package domain

type Tournament struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Logo     string `json:"logo"`
}
