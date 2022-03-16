package main

import (
	"database/sql"
	"log"

	"github.com/giancarlobastos/loteca-backend/api"
	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
	router   *api.Router
)

func main() {
	defer destroy()
	router.Start(":8080")
}

func init() {
	var err error
	database, err = sql.Open("mysql", "root:secret@tcp(mysql:3306)/loteca?parseTime=true")

	if err != nil {
		log.Fatalf("Error [init]: %v", err)
	}

	teamRepository := repository.NewTeamRepository(database)
	competitionRepository := repository.NewCompetitionRepository(database)
	matchRepository := repository.NewMatchRepository(database)
	lotteryRepository := repository.NewLotteryRepository(database, matchRepository)
	bookmakerRepository := repository.NewBookmakerRepository(database)
	userRepository := repository.NewUserRepository(database)

	apiClient := client.NewApiFootballClient()
	facebookClient := client.NewFacebookClient()

	updateService := service.NewUpdateService(teamRepository, competitionRepository, matchRepository, bookmakerRepository, apiClient)
	apiService := service.NewApiService(userRepository, lotteryRepository, updateService, facebookClient)

	router = api.NewRouter(apiService, updateService)
}

func destroy() {
	err := database.Close()

	if err != nil {
		panic(err.Error())
	}
}
