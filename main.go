package main

import (
	"database/sql"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

func main() {
	teamRepository := repository.NewTeamRepository(database)
	competitionRepository := repository.NewCompetitionRepository(database)
	matchRepository := repository.NewMatchRepository(database)
	apiClient := client.NewApiFootballClient()

	_ = service.NewUpdateService(teamRepository, competitionRepository, apiClient)

	match := &[]domain.Match{{Id: 123, Home: &domain.Team{
		Id: 126,
	}, Away: &domain.Team{
		Id: 126,
	}, Stadium: "Estadio Islas Malvinas", StartAt: time.Now()}}
	matchRepository.InsertMatches(match)
	// updateService := service.NewUpdateService(teamRepository, competitionRepository, apiClient)
	// updateService.ImportCompetitionsAndSeasons()
	// updateService.ImportTeamsAndStadiums()

	defer destroy()
	//image.ConvertSvgToPngWithChrome("https://s.glbimg.com/es/sde/f/organizacoes/2020/02/12/botsvg.svg", "./assets/test.png")
}

func init() {
	var err error
	database, err = sql.Open("mysql", "root:secret@tcp(mysql:3306)/loteca?parseTime=true")

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
