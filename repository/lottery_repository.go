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

func (mr *LotteryRepository) GetCurrentLottery() (*domain.Lottery, error) {
	query :=
		`SELECT number, estimated_prize, main_prize, main_prize_winners, side_prize, 
			side_prize_winners, special_prize, accumulated, end_at
	     FROM lottery 
		 ORDER BY number DESC
		 LIMIT 1`

	return mr.getLottery(query)
}

func (mr *LotteryRepository) GetLottery(number int) (*domain.Lottery, error) {
	query :=
		`SELECT number, estimated_prize, main_prize, main_prize_winners, side_prize, 
			side_prize_winners, special_prize, accumulated, end_at
	     FROM lottery 
		 WHERE number = ?`

	return mr.getLottery(query, &number)
}

func (mr *LotteryRepository) getLottery(query string, args ...interface{}) (*domain.Lottery, error) {
	stmt, err := mr.db.Prepare(query)

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
