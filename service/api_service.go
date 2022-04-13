package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/view"
)

type ApiService struct {
	userRepository        *repository.UserRepository
	lotteryRepository     *repository.LotteryRepository
	pollRepository        *repository.PollRepository
	matchRepository       *repository.MatchRepository
	bookmakerRepository   *repository.BookmakerRepository
	competitionRepository *repository.CompetitionRepository
	updateService         *UpdateService
	facebookClient        *client.FacebookClient
	cacheService          *CacheService
}

func NewApiService(
	userRepository *repository.UserRepository,
	lotteryRepository *repository.LotteryRepository,
	pollRepository *repository.PollRepository,
	matchRepository *repository.MatchRepository,
	bookmakerRepository *repository.BookmakerRepository,
	competitionRepository *repository.CompetitionRepository,
	updateService *UpdateService,
	facebookClient *client.FacebookClient,
	cacheService *CacheService) *ApiService {
	return &ApiService{
		userRepository:        userRepository,
		lotteryRepository:     lotteryRepository,
		pollRepository:        pollRepository,
		matchRepository:       matchRepository,
		bookmakerRepository:   bookmakerRepository,
		competitionRepository: competitionRepository,
		updateService:         updateService,
		facebookClient:        facebookClient,
		cacheService:          cacheService,
	}
}

func (as *ApiService) GetCurrentLottery() (*view.Lottery, error) {
	lottery, err := as.cacheService.Get("currentLottery")

	if err != nil {
		lottery, err = as.lotteryRepository.GetCurrentLottery()

		if err != nil {
			return nil, err
		}

		as.cacheService.Put("currentLottery", lottery)
	}

	return lottery.(*view.Lottery), nil
}

func (as *ApiService) GetLottery(number int) (*view.Lottery, error) {
	key := fmt.Sprint("lottery_", number)
	lottery, err := as.cacheService.Get(key)

	if err != nil {
		lottery, err = as.lotteryRepository.GetLottery(number)

		if err != nil {
			return nil, err
		}

		as.cacheService.Put(key, lottery)
	}

	return lottery.(*view.Lottery), nil
}

func (as *ApiService) GetPollResults(lotteryId int) (*view.PollResults, error) {
	key := fmt.Sprint("poll_results_", lotteryId)
	pollResults, err := as.cacheService.Get(key)

	if err != nil {
		pollResults, err = as.pollRepository.GetPollResults(lotteryId)

		if err != nil {
			return nil, err
		}

		as.cacheService.Put(key, pollResults)
	}

	return pollResults.(*view.PollResults), nil
}

func (as *ApiService) Vote(poll domain.Poll, user domain.User) error {
	lottery, err := as.GetCurrentLottery()

	if err == nil && lottery.Id != nil {

		if time.Now().After(*lottery.EndAt) {
			return errors.New("voting period is over")
		}

		err = as.pollRepository.Vote(poll, user)

		if err == nil {
			key := fmt.Sprint("poll_results_", lottery.Id)
			as.cacheService.Delete(key)
		}
	}

	return err
}

func (as *ApiService) GetMatchDetails(matchId int) (*view.MatchDetails, error) {
	key := fmt.Sprint("details_", matchId)
	matchDetails, err := as.cacheService.Get(key)

	if err != nil {
		match, err := as.matchRepository.GetMatch(matchId)

		if err != nil {
			return nil, err
		}

		h2h, _ := as.matchRepository.GetH2HMatches(*match.HomeId, *match.AwayId, *match.StartAt)
		lastMatchesHome, _ := as.matchRepository.GetLastMatches(*match.HomeId, time.Now())
		lastMatchesHomeCompetition, _ := as.matchRepository.GetLastCompetitionMatches(*match.CompetitionId, *match.Year, *match.HomeId, time.Now())
		lastMatchesAtHome, _ := as.matchRepository.GetLastCompetitionHomeMatches(*match.CompetitionId, *match.Year, *match.HomeId, time.Now())
		nextMatchesHome, _ := as.matchRepository.GetNextMatches(*match.Id, *match.HomeId, *match.StartAt)
		lastMatchesAway, _ := as.matchRepository.GetLastMatches(*match.AwayId, time.Now())
		lastMatchesAwayCompetition, _ := as.matchRepository.GetLastCompetitionMatches(*match.CompetitionId, *match.Year, *match.AwayId, time.Now())
		lastMatchesAtAway, _ := as.matchRepository.GetLastCompetitionAwayMatches(*match.CompetitionId, *match.Year, *match.AwayId, time.Now())
		nextMatchesAway, _ := as.matchRepository.GetNextMatches(*match.Id, *match.AwayId, *match.StartAt)
		votes, totalVotes, _ := as.pollRepository.GetVotes(*match.Id)
		odds, _ := as.getOdds(*match.Id)

		teamsStats := make([]view.TeamStats, 0)
		teamStats, _ := as.competitionRepository.GetTeamStats(*match.CompetitionId, *match.Year, *match.HomeId)

		if teamStats != nil {
			teamsStats = append(teamsStats, *teamStats)
		}

		teamStats, _ = as.competitionRepository.GetTeamStats(*match.CompetitionId, *match.Year, *match.AwayId)

		if teamStats != nil {
			teamsStats = append(teamsStats, *teamStats)
		}

		matchDetails = &view.MatchDetails{
			Id:                         &matchId,
			Match:                      match,
			TeamStats:                  &teamsStats,
			H2H:                        h2h,
			LastMatchesHome:            lastMatchesHome,
			LastMatchesHomeCompetition: lastMatchesHomeCompetition,
			LastMatchesAtHome:          lastMatchesAtHome,
			NextMatchesHome:            nextMatchesHome,
			LastMatchesAway:            lastMatchesAway,
			LastMatchesAwayCompetition: lastMatchesAwayCompetition,
			LastMatchesAtAway:          lastMatchesAtAway,
			NextMatchesAway:            nextMatchesAway,
			Votes:                      votes,
			TotalVotes:                 totalVotes,
			Odds:                       odds,
		}

		as.cacheService.Put(key, matchDetails)
	}

	return matchDetails.(*view.MatchDetails), nil
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

func (as *ApiService) Login(token string) (*string, error) {
	extendedToken, err := as.facebookClient.GetExtendedToken(token)

	if err != nil {
		return nil, errors.New("invalid facebook id")
	}

	return extendedToken, nil
}

func (as *ApiService) Authenticate(token string) (*domain.User, error) {
	key := fmt.Sprint("user_", token)
	user, err := as.cacheService.Get(key)

	if err != nil {
		facebookUser, err := as.getFacebookUser(token)

		if err != nil {
			return nil, errors.New("invalid facebook id")
		}

		authenticatedUser, err := as.userRepository.GetUserByFacebookId(facebookUser.FacebookId)

		if err != nil {
			log.Printf("Error [Authenticate]: %v - [%v]", err, facebookUser.FacebookId)
			return nil, err
		}

		if authenticatedUser == nil {
			authenticatedUser, err = as.userRepository.InsertUser(facebookUser)
		}

		if err != nil {
			log.Printf("Error [Authenticate.InsertUser]: %v - [%v]", err, facebookUser.FacebookId)
			return nil, err
		}

		as.cacheService.Put(key, authenticatedUser)
		return authenticatedUser, nil
	}

	return user.(*domain.User), nil
}

func (as *ApiService) AuthenticateManager(token string) error {
	if token != "0RjZAhNDhXOHZAXZAXNTNTQwWXdsZAmVPZAktVX1RIOXg2YjczMwZDZD" {
		return errors.New("invalid manager")
	}

	return nil
}

func (as *ApiService) getFacebookUser(token string) (*domain.User, error) {
	user, err := as.facebookClient.GetUser(token)

	if err != nil {
		log.Printf("Error [facebook.validateToken]: %v - [%v]", err, token)
		return nil, err
	}

	return user, nil
}

func (as *ApiService) getOdds(matchId int) (*[]view.Odd, error) {
	odds, err := as.bookmakerRepository.GetOdds(matchId)

	if err != nil {
		log.Printf("Error [getOdds]: %v - [%v]", err, matchId)
		return nil, err
	}

	viewOdds := make([]view.Odd, 0)

	for _, odd := range *odds {
		viewOdds = append(viewOdds, view.Odd{
			BookmakerId:   odd.Bookmaker.Id,
			BookmakerName: odd.Bookmaker.Name,
			Home:          odd.Home,
			Draw:          odd.Draw,
			Away:          odd.Away,
		})
	}

	return &viewOdds, nil
}
