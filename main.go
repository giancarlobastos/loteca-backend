package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/giancarlobastos/loteca-backend/api"
	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/daemon"
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
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		router.Start(":9000")
		wg.Done()
	}()

	go func() {
		router.Start(":9001")
		wg.Done()
	}()

	wg.Wait()
}

func init() {
	topic := os.Getenv("TOPIC")
	scoreDaemonEnabled := os.Getenv("SCORE_DAEMON_ENABLED") == "true"
	lotteryDaemonEnabled := os.Getenv("LOTTERY_DAEMON_ENABLED") == "true"
	bookmakerDaemonEnabled := os.Getenv("BOOKMAKER_DAEMON_ENABLED") == "true"

	database = getDatabaseConnection()

	teamRepository := repository.NewTeamRepository(database)
	competitionRepository := repository.NewCompetitionRepository(database)
	matchRepository := repository.NewMatchRepository(database)
	lotteryRepository := repository.NewLotteryRepository(database, matchRepository)
	bookmakerRepository := repository.NewBookmakerRepository(database)
	userRepository := repository.NewUserRepository(database)
	pollRepository := repository.NewPollRepository(database)

	apiClient := client.NewApiFootballClient()
	facebookClient := client.NewFacebookClient()
	firebaseClient := client.NewFirebaseClient()
	notificationService := service.NewNotificationService(firebaseClient, topic)

	cacheService := service.NewCacheService()
	updateService := service.NewUpdateService(teamRepository, competitionRepository, matchRepository, bookmakerRepository, apiClient)
	apiService := service.NewApiService(userRepository, lotteryRepository, pollRepository, matchRepository, bookmakerRepository, competitionRepository, updateService, facebookClient, cacheService, notificationService)

	log.Printf("topic: %s", topic)

	liveScoreDaemon := daemon.NewLiveScoreDaemon(apiClient, updateService, cacheService, lotteryRepository, notificationService)
	lottteryDaemon := daemon.NewLotteryDaemon(cacheService, lotteryRepository, notificationService)
	bookmakerDaemon := daemon.NewBookmakerDaemon(updateService, lotteryRepository)

	if scoreDaemonEnabled {
		go liveScoreDaemon.CheckLiveScores()
	}

	if lotteryDaemonEnabled {
		go lottteryDaemon.CheckUpdates()
	}

	if bookmakerDaemonEnabled {
		go bookmakerDaemon.UpdateOdds()
	}

	router = api.NewRouter(apiService, updateService, notificationService)
}

func getDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:secret@tcp(mysql:3306)/loteca?parseTime=true")

	if err != nil {
		log.Fatalf("Error [getDatabaseConnection]: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error [getDatabaseConnection]: %v", err)
	}

	return db
}

func destroy() {
	err := database.Close()

	if err != nil {
		panic(err.Error())
	}
}
