package service

import (
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

func (us *UpdateService) ImportTeams() error {
	countries := [...]string{"Brazil", "Argentina", "Italy", "Germany", "Spain"}
	for _, country := range countries {
		teams, err := us.getTeams(country)

		if err == nil {
			us.teamRepository.InsertTeams(teams)
		}
	}
	return nil
}

func (us *UpdateService) getTeams(country string) (*[]domain.Team, error) {
	response, _ := us.apiClient.GetTeams(country)
	teams := make([]domain.Team, response.Size)

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

func (us *UpdateService) ImportCompetitions() error {
	countries := [...]string{"Brazil", "Argentina", "Italy", "Germany", "Spain", "France", "England"}
	for _, country := range countries {
		competitions, err := us.getCompetitions(country)

		if err == nil {
			us.competitionRepository.InsertCompetitions(competitions)
		}
	}
	return nil
}

func (us *UpdateService) getCompetitions(country string) (*[]domain.Competition, error) {
	response, _ := us.apiClient.GetLeagues(country)
	competitions := make([]domain.Competition, response.Size)

	for _, result := range response.Results {
		seasons := make([]domain.Season, len(result.Seasons))

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

func (us *UpdateService) ImportMatches() error {
	// countries := [...]string{"Brazil", "Argentina", "Italy", "Germany", "Spain"}
	// for _, country := range countries {
	competitions, err := us.competitionRepository.GetCompetitions("Brazil")

	if err != nil {
		return err
	}

	for _, competition := range *competitions {
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
	}
	return nil
}

func (us *UpdateService) getRoundsWithMatches(competitionId uint32, year uint) (*[]domain.Round, error) {
	response, _ := us.apiClient.GetFixtures(competitionId, year)
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

		startAt, _ := time.Parse(time.RFC3339, result.Fixture.DateAndTime)
		match := domain.Match{
			Id: result.Fixture.Id,
			Home: &domain.Team{
				Id: result.Teams[0].Id,
			},
			Away: &domain.Team{
				Id: result.Teams[1].Id,
			},
			Stadium:   result.Fixture.Venue.Name,
			StartAt:   startAt,
			HomeScore: uint(result.Goals.Home),
			AwayScore: uint(result.Goals.Away),
		}

		*round.Matches = append(*round.Matches, match)
	}

	return &rounds, nil
}
