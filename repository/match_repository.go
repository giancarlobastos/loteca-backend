package repository

import (
	"database/sql"
	"log"

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

func (mr *MatchRepository) InsertRoundsAndMatches(competitionId uint32, year uint, rounds *[]domain.Round) error {
	stmt, err := mr.db.Prepare("INSERT IGNORE INTO `match`(id, round_id, home_id, away_id, stadium_id, start_at, home_score, away_score) " +
		"VALUES(?, (SELECT r.id FROM round r WHERE r.name = ?), ?, ?, (SELECT s.id FROM stadium s WHERE s.name = ?), ?, ?, ?) " +
		"ON DUPLICATE KEY UPDATE start_at = VALUES(start_at), home_score = VALUES(home_score), away_score = VALUES(away_score) ")

	if err != nil {
		return err
	}
	defer stmt.Close()

	roundStmt, err := mr.db.Prepare("INSERT IGNORE INTO round(name, ended, competition_id, year) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer roundStmt.Close()

	for _, round := range *rounds {
		roundStmt.Exec(round.Name, false, competitionId, year)

		if err != nil {
			log.Fatalf("Error: %v - [%v, %v, %v, %v]", err, round.Name, false, competitionId, year)
			return err
		}

		for _, match := range *round.Matches {
			stmt.Exec(match.Id, round.Name, match.Home.Id, match.Away.Id, match.Stadium, match.StartAt, match.HomeScore, match.AwayScore)

			if err != nil {
				log.Fatalf("Error: %v - [%v, %v, %v, %v, %v, %v, %v, %v]", err, match.Id, round.Name, match.Home.Id, match.Away.Id, match.Stadium, match.StartAt, match.HomeScore, match.AwayScore)
				return err
			}
		}
	}

	return nil
}
