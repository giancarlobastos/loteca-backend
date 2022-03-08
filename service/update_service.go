package service

import (
	"log"
	"strconv"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
)

type UpdateService struct {
	teamRepository        *repository.TeamRepository
	competitionRepository *repository.CompetitionRepository
	matchRepository       *repository.MatchRepository
	apiClient             *client.ApiFootballClient
}

func NewUpdateService(
	teamRepository *repository.TeamRepository,
	competitionRepository *repository.CompetitionRepository,
	matchRepository *repository.MatchRepository,
	apiClient *client.ApiFootballClient) *UpdateService {
	return &UpdateService{
		teamRepository:        teamRepository,
		competitionRepository: competitionRepository,
		matchRepository:       matchRepository,
		apiClient:             apiClient,
	}
}

func (us *UpdateService) GetTeams(country string) (*[]domain.Team, error) {
	return us.teamRepository.GetTeams(country)
}

func (us *UpdateService) ImportTeams(country string) error {
	teams, err := us.getTeams(country)

	if err == nil {
		us.teamRepository.InsertTeams(teams)
	}

	return nil
}

func (us *UpdateService) getTeams(country string) (*[]domain.Team, error) {
	response, err := us.apiClient.GetTeams(country)

	if err != nil {
		log.Fatalf("Error [getTeams]: %v - [%v]", err, country)
		return &[]domain.Team{}, err
	}

	teams := make([]domain.Team, 0)

	for _, result := range response.Results {
		teams = append(teams, domain.Team{
			Id:      result.Team.Id,
			Name:    result.Team.Name,
			Country: country,
			Logo:    result.Team.LogoUrl,
			Stadium: &domain.Stadium{
				Id:      result.Venue.Id,
				Name:    result.Venue.Name,
				City:    result.Venue.City,
				Country: country,
			},
		})
	}
	return &teams, nil
}

func (us *UpdateService) GetCompetitions(country string, year uint) (*[]domain.Competition, error) {
	return us.competitionRepository.GetCompetitions(country, year, false)
}

func (us *UpdateService) ImportCompetitions(country string) error {
	competitions, err := us.getCompetitions(country)

	if err == nil {
		us.competitionRepository.InsertCompetitions(competitions)
	}

	return nil
}

func (us *UpdateService) getCompetitions(country string) (*[]domain.Competition, error) {
	response, err := us.apiClient.GetLeagues(country)

	if err != nil {
		log.Fatalf("Error [getCompetitions]: %v - [%v]", err, country)
		return &[]domain.Competition{}, err
	}

	competitions := make([]domain.Competition, 0)

	for _, result := range response.Results {
		seasons := make([]domain.Season, 0)

		for _, season := range result.Seasons {
			seasons = append(seasons, domain.Season{
				Year:  season.Year,
				Name:  strconv.Itoa(int(season.Year)),
				Ended: !season.Current,
			})
		}

		competition := domain.Competition{
			Id:      result.League.Id,
			Name:    result.League.Name,
			Logo:    result.League.LogoUrl,
			Type:    domain.CompetitionType(result.League.Type),
			Country: result.Country.Name,
			Seasons: &seasons,
		}

		competitions = append(competitions, competition)
	}
	return &competitions, nil
}

func (us *UpdateService) GetMatches(competitionId uint32, year uint) (*[]domain.MatchVO, error) {
	return us.matchRepository.GetMatches(competitionId, year)
}

func (us *UpdateService) ImportMatches(competitionId uint32, year uint) error {
	competition, err := us.competitionRepository.GetCompetition(competitionId, year)

	if err != nil {
		return err
	}

	for _, season := range *competition.Seasons {
		rounds, err := us.getRoundsWithMatches(competition.Id, season.Year)

		if err != nil {
			return err
		}

		err = us.matchRepository.InsertRoundsAndMatches(competition.Id, season.Year, rounds)

		if err != nil {
			return err
		}
	}
	return nil
}

func (us *UpdateService) getRoundsWithMatches(competitionId uint32, year uint) (*[]domain.Round, error) {
	response, err := us.apiClient.GetFixtures(competitionId, year)

	if err != nil {
		log.Fatalf("Error [getRoundsWithMatches]: %v - [%v, %v]", err, competitionId, year)
		return &[]domain.Round{}, err
	}

	rounds := make([]domain.Round, 0)

	var roundName string
	var round domain.Round

	for _, result := range response.Results {
		if result.League.Round != roundName {
			roundName = result.League.Round
			matches := make([]domain.Match, 0)

			round = domain.Round{
				Name:    roundName,
				Matches: &matches,
			}

			rounds = append(rounds, round)
		}

		startAt, err := time.Parse(time.RFC3339, result.Fixture.DateAndTime)

		if err != nil {
			log.Fatalf("Error [getRoundsWithMatches]: %v - [%v, %v, %v]", err, competitionId, year, round.Id)
		}

		match := domain.Match{
			Id: result.Fixture.Id,
			Home: &domain.Team{
				Id: result.Teams.Home.Id,
			},
			Away: &domain.Team{
				Id: result.Teams.Away.Id,
			},
			Stadium:   &domain.Stadium{
				Id: result.Fixture.Venue.Id,
				Name: result.Fixture.Venue.Name,
			},
			StartAt:   startAt,
			HomeScore: uint(result.Goals.Home),
			AwayScore: uint(result.Goals.Away),
		}

		*round.Matches = append(*round.Matches, match)
	}

	return &rounds, nil
}
