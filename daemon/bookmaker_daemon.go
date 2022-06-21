package daemon

import (
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
)

type BookmakerDaemon struct {
	updateService     *service.UpdateService
	lotteryRepository *repository.LotteryRepository
}

func NewBookmakerDaemon(
	updateService *service.UpdateService,
	lotteryRepository *repository.LotteryRepository) *BookmakerDaemon {
	return &BookmakerDaemon{
		updateService:     updateService,
		lotteryRepository: lotteryRepository,
	}
}

func (bd *BookmakerDaemon) UpdateOdds() {
	for {
		bd.updateOdds()
		time.Sleep(24 * time.Hour)
	}
}

func (bd *BookmakerDaemon) updateOdds() {
	log.Println("[BookmakerDaemon]: checking odds...")
	defer catchErrors()

	lottery, err := bd.lotteryRepository.GetCurrentLottery()

	if err != nil {
		log.Printf("[BookmakerDaemon.updateOdds]: %v", err)
		return
	}

	now := time.Now().UTC()

	if now.Before(*lottery.EndAt) {
		for _, match := range *lottery.Matches {
			bd.updateService.ImportOdds(*match.Id)
		}
	}
}
