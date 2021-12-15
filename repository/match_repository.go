package repository

import (
	"database/sql"

	"github.com/giancarlobastos/loteca-backend/domain"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{
		db: db,
	}
}

func (mr *MatchRepository) InsertMatches(matches *[]domain.Match) error {
	stmt, err := mr.db.Prepare("INSERT IGNORE INTO `match`(id, home_id, away_id, stadium_id, start_at, home_score, away_score) " +
		"SELECT ?, ?, ?, s.id, ?, ?, ? FROM stadium s WHERE s.name = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, match := range *matches {
		stmt.Exec(match.Id, match.Home.Id, match.Away.Id, match.StartAt, match.HomeScore, match.AwayScore, match.Stadium)
	}

	return nil
}
