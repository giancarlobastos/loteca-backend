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
	stmt, err := mr.db.Prepare(
		`INSERT INTO ` + "`match`" + `(id, round_id, home_id, away_id, stadium_id, start_at, home_score, away_score)
		 VALUES(?, (SELECT r.id FROM round r WHERE r.name = ? AND r.competition_id = ? AND r.year = ?), ?, ?, (SELECT s.id FROM stadium s WHERE s.name = ?), ?, ?, ?)
		 ON DUPLICATE KEY UPDATE round_id = VALUES(round_id), start_at = VALUES(start_at), home_score = VALUES(home_score), away_score = VALUES(away_score)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	roundStmt, err := mr.db.Prepare(
		`INSERT IGNORE INTO round(name, ended, competition_id, year)
		 VALUES (?, ?, ?, ?)`)

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
			stmt.Exec(match.Id, round.Name, competitionId, year, match.Home.Id, match.Away.Id, match.Stadium, match.StartAt, match.HomeScore, match.AwayScore)

			if err != nil {
				log.Fatalf("Error: %v - [%v, %v, %v, %v, %v, %v, %v, %v]", err, match.Id, round.Name, match.Home.Id, match.Away.Id, match.Stadium, match.StartAt, match.HomeScore, match.AwayScore)
				return err
			}
		}
	}

	return nil
}

func (mr *MatchRepository) GetMatches(competitionId uint32, year uint) (*[]domain.MatchVO, error) {
	stmt, err := mr.db.Prepare(
		`SELECT m.id Id, r.number RoundNumber, r.name RoundName, r.year Year,
			 c.id CompetitionId, c.name CompetitionName,
			 t1.id HomeId, t1.name HomeName, t2.id AwayId, t2.name AwayName, 
			 s.name Stadium, m.start_at StartAt, m.home_score HomeScore, m.away_score AwayScore
		 FROM` + " `match` " + `m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE c.id = ? AND r.year = ?
		 ORDER BY r.id, m.start_at`)

	matches := make([]domain.MatchVO, 0)

	if err != nil {
		return &matches, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(competitionId, year)

	if err != nil {
		log.Fatalf("Error: %v - [%v, %v]", err, competitionId, year)
		return &matches, err
	}

	defer rows.Close()

	var match domain.MatchVO
	
	for rows.Next() {
		match = domain.MatchVO{}
		columns, _ := rows.Columns()
		columnTypes, _ := rows.ColumnTypes()
		rows.
		log.Printf("%v", columns)
		log.Printf("%v", columnTypes)
		rows.Scan(&match.Id, &match.RoundNumber, &match.RoundName, &match.Year, &match.CompetitionId, &match.CompetitionName,
			&match.HomeId, &match.HomeName, &match.AwayId, &match.AwayName, &match.Stadium,
			&match.StartAt, &match.HomeScore, &match.AwayScore)
		
		matches = append(matches, match)
	}

	return &matches, err
}
