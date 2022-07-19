package daemon

import (
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
)

type MatchDaemon struct {
	apiClient         *client.ApiFootballClient
	updateService     *service.UpdateService
	lotteryRepository *repository.LotteryRepository
}

func NewMatchDaemon(apiClient *client.ApiFootballClient,
	updateService *service.UpdateService,
	lotteryRepository *repository.LotteryRepository) *MatchDaemon {
	return &MatchDaemon{
		apiClient:         apiClient,
		updateService:     updateService,
		lotteryRepository: lotteryRepository,
	}
}

func (md *MatchDaemon) CheckMatches() {
	now := time.Now().UTC()
	nextExecution := time.Now().UTC().Truncate(24 * time.Hour).Add(6 * time.Hour)

	if now.After(nextExecution) {
		nextExecution = nextExecution.Add(24 * time.Hour)
	}

	time.Sleep(nextExecution.Sub(now))

	for {
		md.checkMatches()
		time.Sleep(24 * time.Hour)
	}
}

func (md *MatchDaemon) checkMatches() {
	log.Println("[MatchDaemon]: checking match updates...")
	defer catchErrors()

	lottery, err := md.lotteryRepository.GetCurrentLottery()

	if err != nil {
		log.Printf("[MatchDaemon.checkMatches]: %v", err)
		return
	}

	now := time.Now().UTC()

	if now.Before(*lottery.EndAt) {
		for _, match := range *lottery.Matches {
			md.updateService.ImportLastMatches(*match.HomeId)
			md.updateService.ImportLastMatches(*match.AwayId)
		}
	}
}
