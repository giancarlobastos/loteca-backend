package repository

import (
	"database/sql"
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
)

type CompetitionRepository struct {
	db *sql.DB
}

func NewCompetitionRepository(db *sql.DB) *CompetitionRepository {
	return &CompetitionRepository{
		db: db,
	}
}

func (cr *CompetitionRepository) InsertCompetitions(competitions *[]domain.Competition) error {
	stmt, err := cr.db.Prepare(
		`INSERT INTO competition(id, name, logo, type, country)
		 VALUES(?, ?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE 
		 	logo = coalesce(VALUES(logo), logo), 
			type = coalesce(VALUES(type), type), 
			name = coalesce(VALUES(name), name)`)

	if err != nil {
		log.Printf("Error [InsertCompetitions]: %v", err)
		return err
	}

	defer stmt.Close()

	for _, competition := range *competitions {
		_, err = stmt.Exec(competition.Id, competition.Name, competition.Logo, competition.Type, competition.Country)

		if err != nil {
			log.Printf("Error [InsertCompetitions]: %v - [%v, %v, %v, %v, %v]", err, competition.Id, competition.Name, competition.Logo, competition.Type, competition.Country)
			continue
		}

		cr.insertSeasons(competition.Id, competition.Seasons)
	}

	return nil
}

func (cr *CompetitionRepository) insertSeasons(competitionId int, seasons *[]domain.Season) error {
	stmt, err := cr.db.Prepare(
		`INSERT INTO season(competition_id, year, name, ended)
		 VALUES(?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE 
		 	year = coalesce(VALUES(year), year), 
			name = coalesce(VALUES(name), name), 
			ended = coalesce(VALUES(ended), ended)`)

	if err != nil {
		log.Printf("Error [insertSeasons]: %v - [%v]", err, competitionId)
		return err
	}

	defer stmt.Close()

	for _, season := range *seasons {
		if (domain.Season{}) != season {
			_, err = stmt.Exec(competitionId, season.Year, season.Name, season.Ended)

			if err != nil {
				log.Printf("Error: %v - [%v, %v, %v, %v]", err, competitionId, season.Year, season.Name, season.Ended)
			}
		}
	}

	return nil
}

func (cr *CompetitionRepository) GetCompetitions(country string, year int, ended bool) (*[]domain.Competition, error) {
	stmt, err := cr.db.Prepare(
		`SELECT c.id, c.name, s.year, s.ended
		 FROM competition c
		 JOIN season s ON s.competition_id = c.id
		 WHERE c.country = ? AND s.ended = ? AND s.year = ?
		 ORDER BY 1, 2`)
	competitions := make([]domain.Competition, 0)

	if err != nil {
		log.Printf("Error [GetCompetitions]: %v - [%v %v]", err, country, year)
		return &competitions, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(country, ended, year)

	if err != nil {
		log.Printf("Error [GetCompetitions]: %v - [%v, %v, %v]", err, country, year, ended)
		return &competitions, err
	}

	defer rows.Close()

	competition := domain.Competition{}
	var competitionId int = 0
	var competitionName string

	for rows.Next() {
		season := domain.Season{}

		err = rows.Scan(&competitionId, &competitionName, &season.Year, &season.Ended)

		if err != nil {
			log.Printf("Error [GetCompetitions]: %v", err)
			continue
		}

		if competition.Id != competitionId {
			seasons := make([]domain.Season, 0)
			competition = domain.Competition{
				Id:      competitionId,
				Name:    competitionName,
				Seasons: &seasons,
			}

			competitions = append(competitions, competition)
		}

		*competition.Seasons = append(*competition.Seasons, season)
	}

	return &competitions, err
}
func (cr *CompetitionRepository) GetCompetition(competitionId int, year int) (*domain.Competition, error) {
	stmt, err := cr.db.Prepare(
		`SELECT c.id, c.name, s.year, s.ended
		 FROM competition c
		 JOIN season s ON s.competition_id = c.id
		 WHERE c.id = ? AND s.year = ?
		 ORDER BY 1, 2`)

	if err != nil {
		log.Printf("Error [GetCompetition]: %v - [%v, %v]", err, competitionId, year)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(competitionId, year)

	if err != nil {
		log.Printf("Error [GetCompetition]: %v - [%v, %v]", err, competitionId, year)
		return nil, err
	}

	defer rows.Close()

	competition := domain.Competition{}
	var competitionName string

	for rows.Next() {
		season := domain.Season{}
		err = rows.Scan(&competitionId, &competitionName, &season.Year, &season.Ended)

		if err != nil {
			log.Printf("Error [GetCompetition]: %v", err)
			continue
		}

		seasons := make([]domain.Season, 0)
		competition = domain.Competition{
			Id:      competitionId,
			Name:    competitionName,
			Seasons: &seasons,
		}

		*competition.Seasons = append(*competition.Seasons, season)
	}

	return &competition, err
}
