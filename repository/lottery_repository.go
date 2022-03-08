package repository

import (
	"database/sql"
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
)

type LotteryRepository struct {
	db *sql.DB
}

func NewLotteryRepository(db *sql.DB) *LotteryRepository {
	return &LotteryRepository{
		db: db,
	}
}

func (lr *LotteryRepository) GetCurrentLottery() (*domain.Lottery, error) {
	query :=
		`SELECT number, estimated_prize, main_prize, main_prize_winners, side_prize, 
			side_prize_winners, special_prize, accumulated, end_at
	     FROM lottery 
		 ORDER BY number DESC
		 LIMIT 1`

	return lr.getLottery(query)
}

func (lr *LotteryRepository) GetLottery(number int) (*domain.Lottery, error) {
	query :=
		`SELECT number, estimated_prize, main_prize, main_prize_winners, side_prize, 
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
		_, err = matchStmt.Exec(lottery.Number, match.Match.Id, match.Order)

		if err != nil {
			log.Fatalf("Error: %v - [%v, %v, %v]", err, lottery.Number, match.Match.Id, match.Order)
			return nil, err
		}
	}

	return &lottery, nil
}

func (lr *LotteryRepository) getLottery(query string, args ...interface{}) (*domain.Lottery, error) {
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

	lottery := domain.Lottery{}

	if rows.Next() {
		rows.Scan(&lottery.Number, &lottery.EstimatedPrize, &lottery.MainPrize, &lottery.MainPrizeWinners,
			&lottery.SidePrize, &lottery.SidePrizeWinners, &lottery.SpecialPrize, &lottery.Accumulated, &lottery.EndAt)
	}

	return &lottery, err
}
