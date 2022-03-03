package service

import (
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
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

func (as *ApiService) GetCurrentLottery() (loterry *domain.Lottery, err error) {
	return as.lotteryRepository.GetCurrentLottery()
}
