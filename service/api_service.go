package service

import (
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/view"
)

type ApiService struct {
	lotteryRepository *repository.LotteryRepository
	updateSerice      *UpdateService
}

func NewApiService(
	lotteryRepository *repository.LotteryRepository,
	updateService *UpdateService) *ApiService {
	return &ApiService{
		lotteryRepository: lotteryRepository,
		updateSerice:      updateService,
	}
}

func (as *ApiService) GetCurrentLottery() (lottery *view.Lottery, err error) {
	return as.lotteryRepository.GetCurrentLottery()
}

func (as *ApiService) GetLottery(number int) (lottery *view.Lottery, err error) {
	return as.lotteryRepository.GetLottery(number)
}

func (as *ApiService) CreateLottery(lottery domain.Lottery) (*domain.Lottery, error) {
	_, err := as.lotteryRepository.CreateLottery(lottery)

	if err != nil {
		log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Id)
	}

	for _, match := range *lottery.Matches {
		err = as.updateSerice.ImportHeadToHead(match.Home.Id, match.Away.Id)

		if err != nil {
			log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Id)
		}

		err = as.updateSerice.ImportLastMatches(match.Home.Id)

		if err != nil {
			log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Id)
		}

		err = as.updateSerice.ImportLastMatches(match.Away.Id)

		if err != nil {
			log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Id)
		}

		err = as.updateSerice.ImportNextMatches(match.Home.Id)

		if err != nil {
			log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Id)
		}

		err = as.updateSerice.ImportNextMatches(match.Away.Id)

		if err != nil {
			log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Id)
		}
	}

	return &lottery, nil
}
