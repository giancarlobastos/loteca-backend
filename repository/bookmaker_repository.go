package repository

import (
	"database/sql"
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
)

type BookmakerRepository struct {
	db *sql.DB
}

func NewBookmakerRepository(db *sql.DB) *BookmakerRepository {
	return &BookmakerRepository{
		db: db,
	}
}

func (br *BookmakerRepository) InsertOdds(odds *[]domain.Odd) error {
	bookmakerStmt, err := br.db.Prepare(`INSERT IGNORE INTO betting_platform(id, name) VALUES(?, ?)`)

	if err != nil {
		log.Printf("Error [InsertOdds]: %v", err)
		return err
	}

	defer bookmakerStmt.Close()

	oddsStmt, err := br.db.Prepare(
		`INSERT INTO match_odds(platform_id, match_id, home_odds, draw_odds, away_odds, updated_at)
		 VALUES(?, ?, ?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE 
			home_odds = coalesce(VALUES(home_odds), home_odds), 
			draw_odds = coalesce(VALUES(draw_odds), draw_odds), 
			away_odds = coalesce(VALUES(away_odds), away_odds), 
			updated_at = coalesce(VALUES(updated_at), updated_at)`)

	if err != nil {
		log.Printf("Error [InsertOdds]: %v", err)
		return err
	}

	defer oddsStmt.Close()

	for _, odd := range *odds {
		_, err = bookmakerStmt.Exec(odd.Bookmaker.Id, odd.Bookmaker.Name)

		if err != nil {
			log.Printf("Error [InsertOdds]: %v - [%v %v]", err, odd.Bookmaker.Id, odd.Bookmaker.Name)
			continue
		}

		_, err = oddsStmt.Exec(odd.Bookmaker.Id, odd.Id, odd.Home, odd.Draw, odd.Away, *odd.UpdatedAt)

		if err != nil {
			log.Printf("Error [InsertOdds]: %v - [%v %v %v %v %v %v]", err, odd.Bookmaker.Id, odd.Id, odd.Home, odd.Draw, odd.Away, *odd.UpdatedAt)
			continue
		}
	}

	return nil
}

func (tr *BookmakerRepository) GetOdds(matchId int) (*[]domain.Odd, error) {
	stmt, err := tr.db.Prepare(
		`SELECT bp.id, bp.name, mo.home_odds, mo.draw_odds, mo.away_odds
		 FROM match_odds mo
		 JOIN betting_platform bp ON mo.platform_id = bp.id 
		 WHERE mo.match_id = ? AND bp.preference IS NOT NULL
		 ORDER BY isnull(bp.preference), bp.preference`)

	odds := make([]domain.Odd, 0)

	if err != nil {
		log.Printf("Error [GetOdds]: %v", err)
		return &odds, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(matchId)

	if err != nil {
		log.Printf("Error [GetOdds]: %v - [%v]", err, matchId)
		return &odds, err
	}

	defer rows.Close()

	odd := domain.Odd{}

	for rows.Next() {
		bookmaker := domain.Bookmaker{}

		err = rows.Scan(&bookmaker.Id, &bookmaker.Name, &odd.Home, &odd.Draw, &odd.Away)

		if err != nil {
			log.Printf("Error [GetOdds]: %v - [%v]", err, matchId)
			continue
		}

		odd.Bookmaker = bookmaker
		odds = append(odds, odd)
	}

	return &odds, err
}
