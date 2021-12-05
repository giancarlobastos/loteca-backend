package service

import (
	"fmt"
	"github.com/giancarlobastos/loteca-backend/repository"
)

type ScraperService struct {
	competitionRepository *repository.CompetitionRepository
}

func NewScraperService(competitionRepository *repository.CompetitionRepository) *ScraperService {
	return &ScraperService{
		competitionRepository: competitionRepository,
	}
}

func (ss *ScraperService) ImportMatches() error {
	seasons, _ := ss.competitionRepository.GetOpenSeasons()
	fmt.Printf("%v", seasons)
	return nil
}
