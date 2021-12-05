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

func (cr *CompetitionRepository) GetOpenSeasons() (seasons []domain.Season, err error) {
	rows, err := cr.db.Query(
		"SELECT s.id, s.name, s.code, c.code c_code, c.code_name c_code_name " +
			"FROM season s " +
			"JOIN competition c ON s.competition_id = c.id " +
			"WHERE c.ended IS NOT TRUE")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var season domain.Season
		if err := rows.Scan(
			&season.Id,
			&season.Name,
			&season.Code,
			&season.Competition.Code,
			&season.Competition.CodeName); err != nil {
			return nil, err
		}
		seasons = append(seasons, season)
	}

	return seasons, nil
}

func (cr *CompetitionRepository) GetCompetition(id int) (domain.Competition, error) {
	competitions, err := cr.getCompetitions(
		"SELECT id, name, division, code, code_name, logo "+
			"FROM competition "+
			"WHERE id = ?", id)

	if err != nil {
		return domain.Competition{}, err
	}

	if len(competitions) == 0 {
		return domain.Competition{}, sql.ErrNoRows
	}

	return competitions[0], nil
}

func (cr *CompetitionRepository) getCompetitions(query string, args ...interface{}) (competitions []domain.Competition, err error) {
	rows, err := cr.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var competition domain.Competition
		if err := rows.Scan(
			&competition.Id,
			&competition.Name,
			&competition.Division,
			&competition.Logo); err != nil {
			return nil, err
		}
		competitions = append(competitions, competition)
	}

	return competitions, nil
}
