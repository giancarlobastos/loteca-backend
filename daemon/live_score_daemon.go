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

		if nextCheck == 0 {
			time.Sleep(5 * time.Minute)
		} else {
			time.Sleep(nextCheck)
		}
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

	now := time.Now().UTC()
	update := false
	earliestMatchAt := time.Unix(99999999999999, 0)

	for _, match := range *lottery.Matches {
		if match.Ended || match.Raffle ||
			(match.Status != nil &&
				(*match.Status == "CANC" ||
					*match.Status == "SUSP" ||
					*match.Status == "INT" ||
					*match.Status == "PST" ||
					*match.Status == "ABD" ||
					*match.Status == "AWD" ||
					*match.Status == "WO")) {
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
		lsd.notifyUpdates(*lottery.Id, &(*lottery.Matches)[i], &(*updatedLottery.Matches)[i], now.Unix())
	}

	return 5 * time.Minute
}

func (lsd *LiveScoreDaemon) notifyUpdates(lotteryId int, oldMatch *view.Match, newMatch *view.Match, timestamp int64) error {
	if (newMatch.HomeScore != nil && newMatch.AwayScore != nil) &&
		((oldMatch.HomeScore == nil && *newMatch.HomeScore > 0) ||
			(oldMatch.AwayScore == nil && *newMatch.AwayScore > 0) ||
			(oldMatch.HomeScore != nil && *oldMatch.HomeScore != *newMatch.HomeScore) ||
			(oldMatch.AwayScore != nil && *oldMatch.AwayScore != *newMatch.AwayScore)) {
		log.Printf("Goal: %v", *newMatch)
		lsd.notificationService.NotifyMatchScore(lotteryId, newMatch, timestamp)
	}

	if oldMatch.Ended != newMatch.Ended {
		log.Printf("Finished: %v", *newMatch)
		return lsd.notificationService.NotifyMatchFinished(lotteryId, newMatch, timestamp)
	}

	if *oldMatch.Status != *newMatch.Status && *newMatch.Status == "HT" {
		log.Printf("Interval: %v", *newMatch)
		return lsd.notificationService.NotifyMatchInterval(lotteryId, newMatch, timestamp)
	}

	return nil
}

func catchErrors() {
	if err := recover(); err != nil {
		log.Printf("[daemon error] %v", err)
	}
}
