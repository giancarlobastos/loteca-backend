package repository

import (
	"database/sql"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{
		db: db,
	}
}

func (mr *MatchRepository) GetMatchIdsWithoutScore(id int) (ids []uint32, err error) {
	rows, err := mr.db.Query(
		"SELECT id " +
			"FROM match " +
			"WHERE home_score IS NULL " +
			"AND start_at < NOW()")

	if err != nil {
		return make([]uint32, 0), err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
