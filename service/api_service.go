package service

import (
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/view"
)

type ApiService struct {
	lotteryRepository *repository.LotteryRepository
}

func NewApiService(
	lotteryRepository *repository.LotteryRepository) *ApiService {
	return &ApiService{
		lotteryRepository: lotteryRepository,
	}
}

func (as *ApiService) GetCurrentLottery() (loterry *view.Lottery, err error) {
	return as.lotteryRepository.GetCurrentLottery()
}

func (as *ApiService) GetLottery(number int) (loterry *view.Lottery, err error) {
	return as.lotteryRepository.GetLottery(number)
}

func (as *ApiService) CreateLottery(lottery domain.Lottery) (loterry *domain.Lottery, err error) {
	return as.lotteryRepository.CreateLottery(lottery)
}