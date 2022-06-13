package daemon

import (
	"fmt"
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
	"github.com/giancarlobastos/loteca-backend/view"
)

type LiveScoreDaemon struct {
	apiClient           *client.ApiFootballClient
	updateService       *service.UpdateService
	cacheService        *service.CacheService
	lotteryRepository   *repository.LotteryRepository
	notificationService *service.NotificationService
}

func NewLiveScoreDaemon(apiClient *client.ApiFootballClient,
	updateService *service.UpdateService,
	cacheService *service.CacheService,
	lotteryRepository *repository.LotteryRepository,
	notificationService *service.NotificationService) *LiveScoreDaemon {
	return &LiveScoreDaemon{
		apiClient:           apiClient,
		updateService:       updateService,
		cacheService:        cacheService,
		lotteryRepository:   lotteryRepository,
		notificationService: notificationService,
	}
}

func (lsd *LiveScoreDaemon) CheckLiveScores() {
	for {
		nextCheck := lsd.checkLiveScores()
		time.Sleep(nextCheck)
	}
}

func (lsd *LiveScoreDaemon) checkLiveScores() time.Duration {
	log.Println("[LiveScoreDaemon]: checking live scores...")
	defer catchErrors()

	lottery, err := lsd.lotteryRepository.GetCurrentLottery()

	if err != nil {
		log.Printf("[LiveScoreDaemon.checkLiveScores]: %v", err)
		return 5 * time.Minute
	}

	now := time.Now()
	update := false
	earliestMatchAt := time.Unix(99999999999999, 0)

	for _, match := range *lottery.Matches {
		if match.Ended {
			continue
		}

		if (*match.StartAt).Before(earliestMatchAt) {
			update = true
			earliestMatchAt = *match.StartAt
		}
	}

	if !update {
		return 24 * time.Hour
	}

	if now.Before(earliestMatchAt) {
		return earliestMatchAt.Sub(now)
	}

	lsd.updateService.UpdateMatches(lottery.Matches)
	lsd.cacheService.Delete("currentLottery")
	lsd.cacheService.Delete(fmt.Sprint("lottery_", lottery.Id))

	updatedLottery, err := lsd.lotteryRepository.GetCurrentLottery()

	if err != nil {
		log.Printf("[LiveScoreDaemon.checkLiveScores]: %v", err)
		return 5 * time.Minute
	}

	for i := range *lottery.Matches {
		lsd.notifyUpdates(&(*lottery.Matches)[i], &(*updatedLottery.Matches)[i])
	}

	return 5 * time.Minute
}

func (lsd *LiveScoreDaemon) notifyUpdates(oldMatch *view.Match, newMatch *view.Match) {
	if *oldMatch.HomeScore != *newMatch.HomeScore || *oldMatch.AwayScore != *newMatch.AwayScore {
		log.Printf("Goal: %v", newMatch)
		return
	}

	if oldMatch.Ended != newMatch.Ended {
		log.Printf("Finished: %v", newMatch)
		return
	}

	if *oldMatch.Status != *newMatch.Status && *newMatch.Status == "HT" {
		log.Printf("Interval: %v", newMatch)
		return
	}
}

func catchErrors() {
	if err := recover(); err != nil {
		log.Printf("[LiveScoreDaemon] panic occurred: %v", err)
	}
}
