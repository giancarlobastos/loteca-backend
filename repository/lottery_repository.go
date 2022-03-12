package repository

import (
	"database/sql"
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/view"
)

type LotteryRepository struct {
	db *sql.DB
}

func NewLotteryRepository(db *sql.DB) *LotteryRepository {
	return &LotteryRepository{
		db: db,
	}
}

func (lr *LotteryRepository) GetCurrentLottery() (*view.Lottery, error) {
	query :=
		`SELECT id, number, estimated_prize, main_prize, main_prize_winners, side_prize, 
			side_prize_winners, special_prize, accumulated, end_at
	     FROM lottery 
		 ORDER BY number DESC
		 LIMIT 1`

	return lr.getLottery(query)
}

func (lr *LotteryRepository) GetLottery(number int) (*view.Lottery, error) {
	query :=
		`SELECT id, number, estimated_prize, main_prize, main_prize_winners, side_prize, 
			side_prize_winners, special_prize, accumulated, end_at
	     FROM lottery 
		 WHERE number = ?`

	return lr.getLottery(query, &number)
}

func (lr *LotteryRepository) CreateLottery(lottery domain.Lottery) (*domain.Lottery, error) {
	stmt, err := lr.db.Prepare(
		`INSERT IGNORE INTO lottery(id, name, number, estimated_prize, main_prize, special_prize, accumulated, end_at)
		 VALUES(?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	matchStmt, err := lr.db.Prepare(
		`INSERT IGNORE INTO lottery_match(lottery_id, match_id,` + " `order`) " + `VALUES (?, ?, ?)`)

	if err != nil {
		return nil, err
	}

	defer matchStmt.Close()

	_, err = stmt.Exec(lottery.Number, lottery.Name, lottery.Number, lottery.EstimatedPrize, lottery.MainPrize, lottery.SpecialPrize, lottery.Accumulated, lottery.EndAt)

	if err != nil {
		log.Fatalf("Error: %v - [%v, %v, %v, %v, %v, %v, %v]", err, lottery.Name, lottery.Number, lottery.EstimatedPrize, lottery.MainPrize, lottery.SpecialPrize, lottery.Accumulated, lottery.EndAt)
		return nil, err
	}

	for _, match := range *lottery.Matches {
		_, err = matchStmt.Exec(lottery.Number, match.Id, match.Order)

		if err != nil {
			log.Fatalf("Error: %v - [%v, %v, %v]", err, lottery.Number, match.Id, match.Order)
			return nil, err
		}
	}

	return &lottery, nil
}

func (lr *LotteryRepository) getLottery(query string, args ...interface{}) (*view.Lottery, error) {
	stmt, err := lr.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if len(args) == 0 {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(args)
	}

	if err != nil {
		log.Fatalf("Error: %v - Could not get lottery number %v", err, args)
		return nil, err
	}

	defer rows.Close()

	lottery := view.Lottery{}

	if rows.Next() {
		rows.Scan(&lottery.Id, &lottery.Number, &lottery.EstimatedPrize, &lottery.MainPrize, &lottery.MainPrizeWinners,
			&lottery.SidePrize, &lottery.SidePrizeWinners, &lottery.SpecialPrize, &lottery.Accumulated, &lottery.EndAt)
	}

	lottery.Matches, _ = lr.getMatches(lottery.Id)

	return &lottery, err
}

func (lr *LotteryRepository) getMatches(lotteryId int) (*[]view.Match, error) {
	stmt, err := lr.db.Prepare(
		`SELECT m.id, r.number, r.name, r.year, c.id, c.name, t1.id, t1.name, t2.id, t2.name, s.name, 
		 	m.start_at, m.home_score, m.away_score, lm.order
		 FROM lottery_match lm
		 JOIN` + " `match` " + `m ON  lm.match_id = m.id
		 JOIN round r ON m.round_id = r.id
		 JOIN competition c ON r.competition_id = c.id
		 JOIN team t1 ON m.home_id = t1.id
		 JOIN team t2 ON m.away_id = t2.id
		 LEFT JOIN stadium s ON m.stadium_id = s.id
		 WHERE lm.lottery_id = ?
		 ORDER BY lm.order`)

	matches := make([]view.Match, 0)

	if err != nil {
		return &matches, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(lotteryId)

	if err != nil {
		log.Fatalf("Error: %v - [%v]", err, lotteryId)
		return &matches, err
	}

	defer rows.Close()

	for rows.Next() {
		match := view.Match{
		}
		err = rows.Scan(&match.Id, &match.RoundNumber, &match.RoundName, &match.Year, &match.CompetitionId, &match.CompetitionName,
			&match.HomeId, &match.HomeName, &match.AwayId, &match.AwayName, &match.Stadium,
			&match.StartAt, &match.HomeScore, &match.AwayScore, &match.Order)

		if err != nil {
			log.Printf("Error: %v", err)
		}

		matches = append(matches, match)
	}

	return &matches, err
}
