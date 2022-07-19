package daemon

import (
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
)

type H2HDaemon struct {
	apiClient         *client.ApiFootballClient
	updateService     *service.UpdateService
	lotteryRepository *repository.LotteryRepository
}

func NewH2HDaemon(apiClient *client.ApiFootballClient,
	updateService *service.UpdateService,
	lotteryRepository *repository.LotteryRepository) *H2HDaemon {
	return &H2HDaemon{
		apiClient:         apiClient,
		updateService:     updateService,
		lotteryRepository: lotteryRepository,
	}
}

func (h2hd *H2HDaemon) CheckMatches() {
	now := time.Now().UTC()
	nextExecution := time.Now().UTC().Truncate(24 * time.Hour).Add(1 * time.Hour)

	if now.After(nextExecution) {
		nextExecution = nextExecution.Add(24 * time.Hour)
	}

	time.Sleep(nextExecution.Sub(now))

	for {
		h2hd.checkMatches()
		time.Sleep(24 * time.Hour)
	}
}

func (h2hd *H2HDaemon) checkMatches() {
	log.Println("[H2HDaemon]: checking h2h matches...")
	defer catchErrors()

	lottery, err := h2hd.lotteryRepository.GetCurrentLottery()

	if err != nil {
		log.Printf("[H2HDaemon.checkMatches]: %v", err)
		return
	}

	now := time.Now().UTC()

	if now.Before(*lottery.EndAt) {
		for _, match := range *lottery.Matches {
			h2hd.updateService.ImportHeadToHead(*match.HomeId, *match.AwayId)
			time.Sleep(1 * time.Minute)
		}
	}
}
