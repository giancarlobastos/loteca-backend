package daemon

import (
	"fmt"
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/service"
	"github.com/giancarlobastos/loteca-backend/view"
)

type LotteryDaemon struct {
	cacheService        *service.CacheService
	lotteryRepository   *repository.LotteryRepository
	notificationService *service.NotificationService
	lastUpdates         map[int]*view.Lottery
}

func NewLotteryDaemon(cacheService *service.CacheService,
	lotteryRepository *repository.LotteryRepository,
	notificationService *service.NotificationService) *LotteryDaemon {
	return &LotteryDaemon{
		cacheService:        cacheService,
		lotteryRepository:   lotteryRepository,
		notificationService: notificationService,
		lastUpdates:         make(map[int]*view.Lottery),
	}
}

func (ld *LotteryDaemon) CheckUpdates() {
	lotteries, err := ld.lotteryRepository.GetLotteryUpdates()

	if err != nil {
		log.Printf("Error populating daemon: %v", err)
		return
	}

	for _, lottery := range *lotteries {
		ld.lastUpdates[*lottery.Id] = &lottery
	}

	for {
		ld.checkUpdates()
		time.Sleep(5 * time.Minute)
	}
}

func (ld *LotteryDaemon) checkUpdates() {
	log.Println("[LotteryDaemon]: checking lottery updates...")
	defer catchErrors()

	updates, err := ld.lotteryRepository.GetLotteryUpdates()

	if err != nil {
		log.Printf("[LotteryDaemon.checkUpdates]: %v", err)
		return
	}

	for _, update := range *updates {
		if lastUpdate, ok := ld.lastUpdates[*update.Id]; ok {
			if (*lastUpdate.UpdatedAt).Before(*update.UpdatedAt) {
				switch {
				case lastUpdate.MainPrizeWinners == nil && update.MainPrizeWinners != nil:
					ld.notificationService.NotifyLotteryResultEvent(&update)
				case !lastUpdate.Enabled && update.Enabled:
					ld.notificationService.NotifyNewLotteryEvent(&update)
				case lastUpdate.EstimatedPrize == nil && update.EstimatedPrize != nil:
					ld.notificationService.NotifyNewLotteryEvent(&update)
				default:
					ld.notificationService.NotifyLotteryUpdateEvent(*update.Id)
				}

				ld.cacheService.Delete("currentLottery")
				ld.cacheService.Delete(fmt.Sprint("lottery_", *update.Id))
			}
		} else if update.Enabled {
			ld.notificationService.NotifyNewLotteryEvent(&update)
		}

		ld.lastUpdates[*update.Id] = &update
	}
}
