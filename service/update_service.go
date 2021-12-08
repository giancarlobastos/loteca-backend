package service

import (
	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
)

type UpdateService struct {
	teamRepository        *repository.TeamRepository
	competitionRepository *repository.CompetitionRepository
	apiClient             *client.ApiFootballClient
}

func NewUpdateService(
	teamRepository *repository.TeamRepository,
	competitionRepository *repository.CompetitionRepository,
	apiClient *client.ApiFootballClient) *UpdateService {
	return &UpdateService{
		teamRepository:        teamRepository,
		competitionRepository: competitionRepository,
		apiClient:             apiClient,
	}
}

func (us *UpdateService) ImportTeamsAndStadiums() error {
	countries := [...]string{"Brazil", "Argentina", "Italy", "Germany", "Spain"}
	for _, country := range countries {
		teams, stadiums, err := us.getTeamsAndStadiums(country)

		if err == nil {
			us.teamRepository.InsertTeams(teams)
			us.teamRepository.InsertStadiums(stadiums)
		}
	}
	return nil
}

func (us *UpdateService) getTeamsAndStadiums(country string) (*[]domain.Team, *[]domain.Stadium, error) {
	response, _ := us.apiClient.GetTeams(country)
	teams := make([]domain.Team, response.Size)
	stadiums := make([]domain.Stadium, response.Size)

	for _, result := range response.Results {
		teams = append(teams, domain.Team{
			Id:      result.Team.Id,
			Name:    result.Team.Name,
			Country: country,
			Logo:    result.Team.LogoUrl,
		})

		stadiums = append(stadiums, domain.Stadium{
			Id:      result.Venue.Id,
			Name:    result.Venue.Name,
			City:    result.Venue.City,
			Country: country,
		})
	}
	return &teams, &stadiums, nil
}
