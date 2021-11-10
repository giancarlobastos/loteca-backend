package main

import (
	"database/sql"
	"fmt"
	"github.com/giancarlobastos/loteca-backend/repository"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

func main() {
	defer destroy()
}

func init() {
	var err error
	database, err = sql.Open("mysql", "root:secret@tcp(mysql:3306)/loteca")

	if err != nil {
		panic(err.Error())
	}

	tournamentRepository := repository.NewTournamentRepository(database)
	var tournament, _ = tournamentRepository.GetTournament(1)
	fmt.Println(tournament)
}

func destroy() {
	err := database.Close()

	if err != nil {
		panic(err.Error())
	}
}
