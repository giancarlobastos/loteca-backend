package service

import (
	"log"
	"strconv"
	"time"

	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/repository"
	"github.com/giancarlobastos/loteca-backend/view"
)

type UpdateService struct {
	teamRepository        *repository.TeamRepository
	competitionRepository *repository.CompetitionRepository
	matchRepository       *repository.MatchRepository
	bookmakerRepository   *repository.BookmakerRepository
	apiClient             *client.ApiFootballClient
}

func NewUpdateService(
	teamRepository *repository.TeamRepository,
	competitionRepository *repository.CompetitionRepository,
	matchRepository *repository.MatchRepository,
	bookmakerRepository *repository.BookmakerRepository,
	apiClient *client.ApiFootballClient) *UpdateService {
	return &UpdateService{
		teamRepository:        teamRepository,
		competitionRepository: competitionRepository,
		matchRepository:       matchRepository,
		bookmakerRepository:   bookmakerRepository,
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
		log.Printf("Error [getTeams]: %v - [%v]", err, country)
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

func (us *UpdateService) GetCompetitions(country string, year int) (*[]domain.Competition, error) {
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
		log.Printf("Error [getCompetitions]: %v - [%v]", err, country)
		return &[]domain.Competition{}, err
	}

	competitions := make([]domain.Competition, 0)

	for _, result := range response.Results {
		seasons := make([]domain.Season, 0)

		for _, season := range result.Seasons {
			seasons = append(seasons, domain.Season{
				Year:  season.Year,
				Name:  strconv.Itoa(season.Year),
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

func (us *UpdateService) GetMatches(competitionId int, year int) (*[]view.Match, error) {
	return us.matchRepository.GetMatches(competitionId, year)
}

func (us *UpdateService) ImportMatches(competitionId int, year int) error {
	response, err := us.apiClient.GetFixtures(competitionId, year)

	if err != nil {
		log.Printf("Error [ImportMatches]: %v - [%v, %v]", err, competitionId, year)
	}

	competitions, err := us.getCompetitionAndMatches(response)

	if err != nil {
		log.Printf("Error [ImportMatches]: %v - [%v, %v]", err, competitionId, year)
		return err
	}

	return us.insertCompetitionAndMatches(competitions)
}

func (us *UpdateService) UpdateMatches(matches *[]view.Match) error {
	var ids []int

	for _, match := range *matches {
		ids = append(ids, *match.Id)
	}

	response, err := us.apiClient.GetFixturesById(ids)

	if err != nil {
		log.Printf("Error [ImportMatches]: %v - [%v]", err, ids)
	}

	competitions, err := us.getCompetitionAndMatches(response)

	if err != nil {
		log.Printf("Error [ImportMatches]: %v - [%v]", err, ids)
		return err
	}

	return us.insertCompetitionAndMatches(competitions)
}

func (us *UpdateService) ImportHeadToHead(homeId int, awayId int) error {
	response, err := us.apiClient.GetHeadToHead(homeId, awayId)

	if err != nil {
		log.Printf("Error [ImportHeadToHead]: %v - [%v, %v]", err, homeId, awayId)
	}

	competitions, err := us.getCompetitionAndMatches(response)

	if err != nil {
		log.Printf("Error [ImportHeadToHead]: %v - [%v, %v]", err, homeId, awayId)
		return err
	}

	return us.insertCompetitionAndMatches(competitions)
}

func (us *UpdateService) ImportLastMatches(teamId int) error {
	response, err := us.apiClient.GetLastFixtures(teamId, 6)

	if err != nil {
		log.Printf("Error [ImportLastMatches]: %v - [%v]", err, teamId)
	}

	competitions, err := us.getCompetitionAndMatches(response)

	if err != nil {
		return err
	}

	return us.insertCompetitionAndMatches(competitions)
}

func (us *UpdateService) ImportLastCompetitionMatches(competitionId int, year int, teamId int) error {
	response, err := us.apiClient.GetLastCompetitionFixtures(competitionId, year, teamId, 6)

	if err != nil {
		log.Printf("Error [ImportLastCompetitionMatches]: %v - [%v, %v, %v]", err, competitionId, year, teamId)
	}

	competitions, err := us.getCompetitionAndMatches(response)

	if err != nil {
		return err
	}

	return us.insertCompetitionAndMatches(competitions)
}

func (us *UpdateService) ImportNextMatches(teamId int) error {
	response, err := us.apiClient.GetNextFixtures(teamId, 6)

	if err != nil {
		log.Printf("Error [ImportNextMatches]: %v - [%v]", err, teamId)
	}

	competitions, err := us.getCompetitionAndMatches(response)

	if err != nil {
		return err
	}

	return us.insertCompetitionAndMatches(competitions)
}

func (us *UpdateService) ImportNextCompetitionMatches(competitionId int, year int, teamId int) error {
	response, err := us.apiClient.GetNextCompetitionFixtures(competitionId, year, teamId, 6)

	if err != nil {
		log.Printf("Error [ImportLastCompetitionMatches]: %v - [%v, %v, %v]", err, competitionId, year, teamId)
	}

	competitions, err := us.getCompetitionAndMatches(response)

	if err != nil {
		return err
	}

	return us.insertCompetitionAndMatches(competitions)
}

func (us *UpdateService) ImportOdds(matchId int) error {
	response, err := us.apiClient.GetOdds(matchId)

	if err != nil {
		log.Printf("Error [ImportOdds]: %v - [%v]", err, matchId)
		return err
	}

	odds := make([]domain.Odd, 0)

	if len(response.Results) == 0 {
		log.Printf("Error [ImportOdds]: has no results - [%v %v]", matchId, response)
		return nil
	}

	result := response.Results[0]
	updatedAt, err := time.Parse(time.RFC3339, result.UpdatedAt)

	if err != nil {
		log.Printf("Error [ImportOdds]: %v - [%v, %v]", err, matchId, result.UpdatedAt)
		return err
	}

	for _, bookmaker := range result.Bookmakers {
		odd := domain.Odd{
			Id: result.Fixture.Id,
			Bookmaker: domain.Bookmaker{
				Id:   bookmaker.Id,
				Name: bookmaker.Name,
			},
			Home:      us.getOdd(bookmaker.Bets[0], "Home"),
			Draw:      us.getOdd(bookmaker.Bets[0], "Draw"),
			Away:      us.getOdd(bookmaker.Bets[0], "Away"),
			UpdatedAt: &updatedAt,
		}

		odds = append(odds, odd)
	}

	return us.bookmakerRepository.InsertOdds(&odds)
}

func (us *UpdateService) getOdd(bet client.Bet, name string) float32 {
	for _, odd := range bet.Odds {
		if odd.Name == name {
			value, _ := strconv.ParseFloat(odd.Value, 32)
			return float32(value)
		}
	}

	return 0
}

func (us *UpdateService) insertCompetitionAndMatches(competitions *[]domain.Competition) error {
	for _, competition := range *competitions {
		err := us.competitionRepository.InsertCompetitions(&([]domain.Competition{competition}))

		if err != nil {
			log.Printf("Error [insertCompetitionAndMatches.InsertCompetitions]: %v - [%v]", err, competition)
			return err
		}

		seasons := *competition.Seasons
		rounds := *seasons[0].Rounds
		matches := *rounds[0].Matches

		err = us.teamRepository.InsertTeams(&[]domain.Team{*matches[0].Home, *matches[0].Away})

		if err != nil {
			log.Printf("Error [insertCompetitionAndMatches.InsertTeams]: %v", err)
			return err
		}

		err = us.teamRepository.InsertStadium(matches[0].Stadium)

		if err != nil {
			log.Printf("Error [insertCompetitionAndMatches.InsertStadium]: %v - [%v]", err, (*matches[0].Stadium).Id)
			return err
		}

		err = us.matchRepository.InsertRoundsAndMatches(competition.Id, seasons[0].Year, seasons[0].Rounds)

		if err != nil {
			log.Printf("Error [insertCompetitionAndMatches.InsertRoundsAndMatches]: %v - [%v]", err, competition)
			return err
		}
	}

	return nil
}

func (us *UpdateService) getCompetitionAndMatches(fixtures *client.GetFixturesResponse) (*[]domain.Competition, error) {
	competitions := make([]domain.Competition, 0)

	for _, result := range fixtures.Results {
		startAt, err := time.Parse(time.RFC3339, result.Fixture.DateAndTime)

		if err != nil {
			log.Printf("Error [getHeadToHead]: %v - [%v, %v, %v]", err, result.League.Id, result.League.Season, result.Fixture.DateAndTime)
			continue
		}

		matchEnded := result.Fixture.Status.Code == "FT" ||
			result.Fixture.Status.Code == "AET" ||
			result.Fixture.Status.Code == "PEN" ||
			result.Fixture.Status.Code == "P" ||
			result.Fixture.Status.Code == "ET"

		homeScore := result.Goals.Home
		awayScore := result.Goals.Away

		if result.Score.FullTime.Home != nil {
			homeScore = result.Score.FullTime.Home
		}

		if result.Score.FullTime.Away != nil {
			awayScore = result.Score.FullTime.Away
		}

		match := domain.Match{
			Id: result.Fixture.Id,
			Home: &domain.Team{
				Id:      result.Teams.Home.Id,
				Name:    result.Teams.Home.Name,
				Logo:    result.Teams.Home.LogoUrl,
				Stadium: &domain.Stadium{},
			},
			Away: &domain.Team{
				Id:      result.Teams.Away.Id,
				Name:    result.Teams.Away.Name,
				Logo:    result.Teams.Away.LogoUrl,
				Stadium: &domain.Stadium{},
			},
			Stadium: &domain.Stadium{
				Id:   result.Fixture.Venue.Id,
				Name: result.Fixture.Venue.Name,
				City: result.Fixture.Venue.City,
			},
			StartAt:     startAt,
			HomeScore:   homeScore,
			AwayScore:   awayScore,
			Ended:       matchEnded,
			Status:      result.Fixture.Status.Code,
			ElapsedTime: result.Fixture.Status.ElapsedMinutes,
		}

		competition := domain.Competition{
			Id:      result.League.Id,
			Name:    result.League.Name,
			Logo:    result.League.LogoUrl,
			Country: result.League.Country,
			Seasons: &([]domain.Season{
				{
					Year: result.League.Season,
					Rounds: &([]domain.Round{
						{
							Name:    result.League.Round,
							Matches: &([]domain.Match{match}),
						}}),
				}}),
		}

		competitions = append(competitions, competition)
	}

	return &competitions, nil
}
