package main

import (
	"database/sql"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

func main() {
	competitionRepository := repository.NewCompetitionRepository(database)
	apiClient := client.NewApiFootballClient()
	updateService := service.NewUpdateService(competitionRepository, apiClient)
	updateService.ImportTeams()
	defer destroy()
	//image.ConvertSvgToPngWithChrome("https://s.glbimg.com/es/sde/f/organizacoes/2020/02/12/botsvg.svg", "./assets/test.png")
}

func init() {
	var err error
	database, err = sql.Open("mysql", "root:secret@tcp(mysql:3306)/loteca")

	if err != nil {
		panic(err.Error())
	}
}

func destroy() {
	err := database.Close()

	if err != nil {
		panic(err.Error())
	}
}
