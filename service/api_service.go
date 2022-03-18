package service

import (
	"errors"
	"log"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/view"
)

type ApiService struct {
	userRepository    *repository.UserRepository
	lotteryRepository *repository.LotteryRepository
	updateSerice      *UpdateService
	facebookClient    *client.FacebookClient
}

func NewApiService(
	userRepository *repository.UserRepository,
	lotteryRepository *repository.LotteryRepository,
	updateService *UpdateService,
	facebookClient *client.FacebookClient) *ApiService {
	return &ApiService{
		userRepository:    userRepository,
		lotteryRepository: lotteryRepository,
		updateSerice:      updateService,
		facebookClient:    facebookClient,
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
		log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Number)
	}

	lotteryVO, err := as.lotteryRepository.GetLottery(lottery.Number)

	if err != nil {
		log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Number)
		return nil, err
	}

	go func() {
		for _, match := range *lotteryVO.Matches {
			err = as.updateSerice.ImportHeadToHead(*match.HomeId, *match.AwayId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportHeadToHead]: %v - [%v %v %v]", err, lottery.Number, *match.HomeId, *match.AwayId)
			}

			err = as.updateSerice.ImportLastMatches(*match.HomeId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportLastMatches]: %v - [%v %v]", err, lottery.Number, *match.HomeId)
			}

			err = as.updateSerice.ImportLastMatches(*match.AwayId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportLastMatches]: %v - [%v %v]", err, lottery.Number, *match.AwayId)
			}

			err = as.updateSerice.ImportNextMatches(*match.HomeId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportNextMatches]: %v - [%v %v]", err, lottery.Number, *match.HomeId)
			}

			err = as.updateSerice.ImportNextMatches(*match.AwayId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportNextMatches]: %v - [%v %v]", err, lottery.Number, *match.AwayId)
			}

			err = as.updateSerice.ImportOdds(*match.Id)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportOdds]: %v - [%v %v]", err, lottery.Number, *match.Id)
			}
		}
	}()

	return &lottery, nil
}

func (as *ApiService) Authenticate(token string) (*domain.User, error) {
	user, err := as.getFacebookUser(token)

	if err != nil {
		return nil, errors.New("invalid facebook id")
	}

	authenticatedUser, err := as.userRepository.GetUserByFacebookId(user.FacebookId)

	if err != nil {
		log.Printf("Error [Authenticate]: %v - [%v]", err, user.FacebookId)
		return nil, err
	}

	if authenticatedUser == nil {
		return as.userRepository.InsertUser(user)
	}

	return authenticatedUser, nil
}

func (as *ApiService) getFacebookUser(token string) (*domain.User, error) {
	user, err := as.facebookClient.GetUser(token)

	if err != nil {
		log.Printf("Error [facebook.validateToken]: %v - [%v]", err, token)
		return nil, err
	}

	return user, nil
}
