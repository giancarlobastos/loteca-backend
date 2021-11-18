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

func (pr *CompetitionRepository) GetCompetition(id int) (domain.Competition, error) {
	competitions, err := pr.getCompetitions("SELECT id, name, division, logo FROM competition WHERE id = ?", id)

	if err != nil {
		return domain.Competition{}, err
	}

	if len(competitions) == 0 {
		return domain.Competition{}, sql.ErrNoRows
	}

	return competitions[0], nil
}

func (pr *CompetitionRepository) getTournaments(query string, args ...interface{}) (competitions []domain.Competition, err error) {
	rows, err := pr.db.Query(query, args...)

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
