package service

import (
	"errors"
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/view"
)

type ApiService struct {
	userRepository    *repository.UserRepository
	lotteryRepository *repository.LotteryRepository
	pollRepository    *repository.PollRepository
	updateService     *UpdateService
	facebookClient    *client.FacebookClient
}

func NewApiService(
	userRepository *repository.UserRepository,
	lotteryRepository *repository.LotteryRepository,
	pollRepository *repository.PollRepository,
	updateService *UpdateService,
	facebookClient *client.FacebookClient) *ApiService {
	return &ApiService{
		userRepository:    userRepository,
		lotteryRepository: lotteryRepository,
		pollRepository:    pollRepository,
		updateService:     updateService,
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
			err = as.updateService.ImportHeadToHead(*match.HomeId, *match.AwayId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportHeadToHead]: %v - [%v %v %v]", err, lottery.Number, *match.HomeId, *match.AwayId)
			}

			err = as.updateService.ImportLastMatches(*match.HomeId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportLastMatches]: %v - [%v %v]", err, lottery.Number, *match.HomeId)
			}

			err = as.updateService.ImportLastMatches(*match.AwayId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportLastMatches]: %v - [%v %v]", err, lottery.Number, *match.AwayId)
			}

			err = as.updateService.ImportNextMatches(*match.HomeId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportNextMatches]: %v - [%v %v]", err, lottery.Number, *match.HomeId)
			}

			err = as.updateService.ImportNextMatches(*match.AwayId)

			if err != nil {
				log.Printf("Error [CreateLottery.ImportNextMatches]: %v - [%v %v]", err, lottery.Number, *match.AwayId)
			}

			err = as.updateService.ImportOdds(*match.Id)

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

func (as *ApiService) AuthenticateManager(token string) error {
	if token != "0RjZAhNDhXOHZAXZAXNTNTQwWXdsZAmVPZAktVX1RIOXg2YjczMwZDZD" {
		return errors.New("invalid manager")
	}

	return nil
}

func (as *ApiService) GetPollResults(lotteryId int) (*view.PollResults, error) {
	return as.pollRepository.GetPollResults(lotteryId)
}

func (as *ApiService) Vote(poll domain.Poll, user domain.User) error {
	lottery, err := as.lotteryRepository.GetCurrentLottery()

	if err != nil && lottery.Id != nil {

		if time.Now().After(*lottery.EndAt) {
			return errors.New("voting period is over")
		}

		return as.pollRepository.Vote(poll, user)
	}

	return err
}

func (as *ApiService) getFacebookUser(token string) (*domain.User, error) {
	user, err := as.facebookClient.GetUser(token)

	if err != nil {
		log.Printf("Error [facebook.validateToken]: %v - [%v]", err, token)
		return nil, err
	}

	return user, nil
}
