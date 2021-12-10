package repository

import (
	"database/sql"

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
	stmt, err := cr.db.Prepare("INSERT IGNORE INTO competition(id, name, logo, type, country) VALUES(?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, competition := range *competitions {
		stmt.Exec(competition.Id, competition.Name, competition.Logo, competition.Type, competition.Country)
	}

	return nil
}

func (cr *CompetitionRepository) InsertSeasons(seasons *[]domain.Season) error {
	stmt, err := cr.db.Prepare("INSERT IGNORE INTO season(competition_id, year, name, ended) VALUES(?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, season := range *seasons {
		if (domain.Season{}) != season {
			stmt.Exec(season.Competition.Id, season.Year, season.Name, season.Ended)
		}
	}

	return nil
}
