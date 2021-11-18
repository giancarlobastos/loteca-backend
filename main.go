package main

import (
	"database/sql"
	"github.com/giancarlobastos/loteca-backend/crawler"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

func main() {
	//tournamentRepository := repository.NewTournamentRepository(database)
	//tournament, _ := tournamentRepository.GetTournament(154298)
	//json, _ := json.Marshal(tournament)
	//fmt.Println(string(json))
	c := crawler.NewCrawler()
	c.GetMatch(154298)
	defer destroy()
}

func init() {
	var err error
	database, err = sql.Open("mysql", "root:secret@tcp(mysql:3306)/loteca")

	if err != nil {
		//		panic(err.Error())
	}
}

func destroy() {
	err := database.Close()

	if err != nil {
		panic(err.Error())
	}
}
