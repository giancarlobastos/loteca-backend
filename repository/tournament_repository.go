package repository

import (
	"database/sql"
	"github.com/giancarlobastos/loteca-backend/domain"
)

type TournamentRepository struct {
	db *sql.DB
}

func NewTournamentRepository(db *sql.DB) *TournamentRepository {
	return &TournamentRepository{
		db: db,
	}
}

func (pr *TournamentRepository) GetTournament(id int) (domain.Tournament, error) {
	tournaments, err := pr.getTournaments("SELECT id, name, division, logo FROM tournament WHERE id = ?", id)

	if err != nil {
		return domain.Tournament{}, err
	}

	if len(tournaments) == 0 {
		return domain.Tournament{}, sql.ErrNoRows
	}

	return tournaments[0], nil
}

func (pr *TournamentRepository) getTournaments(query string, args ...interface{}) (tournaments []domain.Tournament, err error) {
	rows, err := pr.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var tournament domain.Tournament
		if err := rows.Scan(
			&tournament.Id,
			&tournament.Name,
			&tournament.Division,
			&tournament.Logo); err != nil {
			return nil, err
		}
		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}
