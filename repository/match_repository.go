package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/view"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{
		db: db,
	}
}

func (mr *MatchRepository) InsertRoundsAndMatches(competitionId int, year int, rounds *[]domain.Round) error {
	tx, err := mr.db.BeginTx(context.Background(), nil)

	if err != nil {
		log.Printf("Error [CreateLottery]: %v", err)
		return err
	}

	defer tx.Rollback()

	stmt, err := mr.db.Prepare(
		`INSERT INTO ` + "`match`" + `(id, round_id, home_id, away_id, stadium_id, start_at, home_score, away_score)
		 VALUES(?, (SELECT r.id FROM round r WHERE r.name = ? AND r.competition_id = ? AND r.year = ?), ?, ?, ?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE round_id = VALUES(round_id), start_at = VALUES(start_at), home_score = VALUES(home_score), away_score = VALUES(away_score), stadium_id = VALUES(stadium_id) `)

	if err != nil {
		log.Printf("Error [InsertRoundsAndMatches]: %v - [%v %v]", err, competitionId, year)
		return err
	}
	defer stmt.Close()

	roundStmt, err := mr.db.Prepare(
		`INSERT IGNORE INTO round(name, ended, competition_id, year)
		 VALUES (?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error [InsertRoundsAndMatches]: %v - [%v %v]", err, competitionId, year)
		return err
	}

	defer roundStmt.Close()

	for _, round := range *rounds {
		_, err = roundStmt.Exec(round.Name, false, competitionId, year)

		if err != nil {
			log.Printf("Error [InsertRoundsAndMatches]: %v - [%v, %v, %v, %v]", err, round.Name, false, competitionId, year)
			return err
		}

		for _, match := range *round.Matches {
			_, err = stmt.Exec(match.Id, round.Name, competitionId, year, match.Home.Id, match.Away.Id, match.Stadium.Id, match.StartAt, match.HomeScore, match.AwayScore)

			if err != nil {
				log.Printf("Error [InsertRoundsAndMatches]: %v - [%v, %v, %v, %v, %v, %v, %v, %v]", err, match.Id, round.Name, match.Home.Id, match.Away.Id, match.Stadium, match.StartAt, match.HomeScore, match.AwayScore)
				return err
			}
		}
	}

	err = tx.Commit()
	return err
}

func (mr *MatchRepository) GetMatch(matchId int) (*view.Match, error) {
	matches, err := mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE m.id = ?`, &matchId)

	if err != nil {
		return nil, err
	}

	if len(*matches) == 0 {
		log.Printf("Error [GetMatch]: match not found - [%v]", matchId)

	}

	return &(*matches)[0], nil
}

func (mr *MatchRepository) GetH2HMatches(homeId int, awayId int, before time.Time) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE t1.id IN (?, ?) AND t2.id IN (?, ?) AND m.start_at < ? AND m.home_score IS NOT NULL
		 ORDER BY m.start_at DESC
		 LIMIT 5`, &homeId, &awayId, &homeId, &awayId, &before)
}

func (mr *MatchRepository) GetLastMatches(teamId int, before time.Time) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE (t1.id = ? OR t2.id = ?) AND m.start_at < ? AND m.home_score IS NOT NULL
		 ORDER BY m.start_at DESC
		 LIMIT 5`, &teamId, &teamId, &before)
}

func (mr *MatchRepository) GetNextMatches(matchId int, teamId int, after time.Time) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE (t1.id = ? OR t2.id = ?) AND m.start_at > ? AND m.home_score IS NULL AND m.id != ?
		 ORDER BY m.start_at
		 LIMIT 5`, &teamId, &teamId, &after, &matchId)
}

func (mr *MatchRepository) GetLastCompetitionMatches(competitionId int, year int, teamId int, before time.Time) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE (t1.id = ? OR t2.id = ?) AND m.start_at < ? AND m.home_score IS NOT NULL AND c.id = ? AND r.year = ?
		 ORDER BY m.start_at DESC
		 LIMIT 5`, &teamId, &teamId, &before, &competitionId, &year)
}

func (mr *MatchRepository) GetLastCompetitionHomeMatches(competitionId int, year int, teamId int, before time.Time) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE t1.id = ? AND m.start_at < ? AND m.home_score IS NOT NULL AND c.id = ? AND r.year = ?
		 ORDER BY m.start_at DESC
		 LIMIT 5`, &teamId, &before, &competitionId, &year)
}

func (mr *MatchRepository) GetLastCompetitionAwayMatches(competitionId int, year int, teamId int, before time.Time) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE t2.id = ? AND m.start_at < ? AND m.home_score IS NOT NULL AND c.id = ? AND r.year = ?
		 ORDER BY m.start_at DESC
		 LIMIT 5`, &teamId, &before, &competitionId, &year)
}

func (mr *MatchRepository) GetMatches(competitionId int, year int) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, NULL
		 FROM`+" `match` "+`m
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE c.id = ? AND r.year = ?
		 ORDER BY r.id, m.start_at`, &competitionId, &year)
}

func (mr *MatchRepository) GetLotteryMatches(lotteryId int) (*[]view.Match, error) {
	return mr.getMatches(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, 
			t1.id, t1.name, coalesce(t1.abbreviation, upper(substr(t1.name, 1, 3))), t1.logo, 
			t2.id, t2.name, coalesce(t2.abbreviation, upper(substr(t2.name, 1, 3))), t2.logo, 
			s.name, m.start_at, m.home_score, m.away_score, lm.order
		 FROM lottery_match lm
		 JOIN`+" `match` "+`m ON  lm.match_id = m.id
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE lm.lottery_id = ?
		 ORDER BY lm.order`, &lotteryId)
}

func (mr *MatchRepository) getMatches(query string, args ...interface{}) (*[]view.Match, error) {
	stmt, err := mr.db.Prepare(query)

	matches := make([]view.Match, 0)

	if err != nil {
		log.Printf("Error [getMatches]: %v - [%v %v]", err, query, args)
		return &matches, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)

	if err != nil {
		log.Printf("Error [getMatches]: %v - [%v]", err, args)
		return &matches, err
	}

	defer rows.Close()

	var match view.Match

	for rows.Next() {
		match = view.Match{}
		err = rows.Scan(&match.Id, &match.RoundNumber, &match.RoundName, &match.Year, &match.CompetitionId, &match.CompetitionName,
			&match.HomeId, &match.HomeName, &match.HomeAbbreviation, &match.HomeLogo, 
			&match.AwayId, &match.AwayName, &match.AwayAbbreviation, &match.AwayLogo, 
			&match.Stadium, &match.StartAt, &match.HomeScore, &match.AwayScore, &match.Order)

		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		matches = append(matches, match)
	}

	return &matches, err
}
