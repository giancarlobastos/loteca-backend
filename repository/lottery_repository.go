package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/view"
)

type LotteryRepository struct {
	db              *sql.DB
	matchRepository *MatchRepository
}

func NewLotteryRepository(
	db *sql.DB,
	matchRepository *MatchRepository) *LotteryRepository {
	return &LotteryRepository{
		db:              db,
		matchRepository: matchRepository,
	}
}

func (lr *LotteryRepository) GetCurrentLottery() (*view.Lottery, error) {
	query :=
		`SELECT id, number, estimated_prize, main_prize, main_prize_winners, side_prize, 
			side_prize_winners, special_prize, accumulated, end_at, result_at
	     FROM lottery 
		 WHERE enabled
		 ORDER BY number DESC
		 LIMIT 1`

	return lr.getLottery(query)
}

func (lr *LotteryRepository) GetLottery(number int) (*view.Lottery, error) {
	query :=
		`SELECT id, number, estimated_prize, main_prize, main_prize_winners, side_prize, 
			side_prize_winners, special_prize, accumulated, end_at, result_at
	     FROM lottery 
		 WHERE number = ? AND enabled`

	return lr.getLottery(query, number)
}

func (lr *LotteryRepository) CreateLottery(lottery domain.Lottery) (*domain.Lottery, error) {
	tx, err := lr.db.BeginTx(context.Background(), nil)

	if err != nil {
		log.Printf("Error [CreateLottery]: %v", err)
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := lr.db.Prepare(
		`INSERT IGNORE INTO lottery(id, name, number, estimated_prize, main_prize, special_prize, accumulated, end_at, result_at)
		 VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Number)
		return nil, err
	}

	defer stmt.Close()

	matchStmt, err := lr.db.Prepare(
		`INSERT IGNORE INTO lottery_match(lottery_id, match_id,` + " `order`) " + `VALUES (?, ?, ?)`)

	if err != nil {
		log.Printf("Error [CreateLottery]: %v - [%v]", err, lottery.Number)
		return nil, err
	}

	defer matchStmt.Close()

	_, err = stmt.Exec(lottery.Number, lottery.Name, lottery.Number, lottery.EstimatedPrize, lottery.MainPrize, lottery.SpecialPrize, lottery.Accumulated, lottery.EndAt, lottery.ResultAt)

	if err != nil {
		log.Printf("Error [CreateLottery]: %v - [%v, %v, %v, %v, %v, %v, %v, %v]", err, lottery.Name, lottery.Number, lottery.EstimatedPrize, lottery.MainPrize, lottery.SpecialPrize, lottery.Accumulated, lottery.EndAt, lottery.ResultAt)
		return nil, err
	}

	for _, match := range *lottery.Matches {
		_, err = matchStmt.Exec(lottery.Number, match.Id, match.Order)

		if err != nil {
			log.Printf("Error [CreateLottery]: %v - [%v, %v, %v]", err, lottery.Number, match.Id, match.Order)
			return nil, err
		}
	}

	err = tx.Commit()
	return &lottery, err
}

func (lr *LotteryRepository) getLottery(query string, args ...interface{}) (*view.Lottery, error) {
	stmt, err := lr.db.Prepare(query)

	if err != nil {
		log.Printf("Error [getLottery]: %v - [%v %v]", err, query, args)
		return nil, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if len(args) == 0 {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(args...)
	}

	if err != nil {
		log.Printf("Error [getLottery]: %v - Could not get lottery number %v", err, args)
		return nil, err
	}

	defer rows.Close()

	lottery := view.Lottery{}

	if rows.Next() {
		err = rows.Scan(&lottery.Id, &lottery.Number, &lottery.EstimatedPrize, &lottery.MainPrize, &lottery.MainPrizeWinners,
			&lottery.SidePrize, &lottery.SidePrizeWinners, &lottery.SpecialPrize, &lottery.Accumulated, &lottery.EndAt, &lottery.ResultAt)

		if err != nil {
			log.Printf("Error [GetCompetition]: %v", err)
			return nil, err
		}

		lottery.Matches, _ = lr.matchRepository.GetLotteryMatches(*lottery.Id)

		earliestMatchAt := time.Unix(99999999999999, 0)
		latestMatchAt := time.Time{}

		for _, match := range *lottery.Matches {
			if (*match.StartAt).Before(earliestMatchAt) {
				earliestMatchAt = *match.StartAt
			}

			if (*match.StartAt).After(latestMatchAt) {
				latestMatchAt = *match.StartAt
			}
		}

		lottery.EarliestMatchAt = &earliestMatchAt
		lottery.LatestMatchAt = &latestMatchAt
		lottery.LotteryIds, _ = lr.getLotteryIdList()
	}

	return &lottery, err
}

func (lr *LotteryRepository) getLotteryIdList() (*[]int, error) {
	stmt, err := lr.db.Prepare(`SELECT id FROM lottery WHERE enabled ORDER BY id`)

	if err != nil {
		log.Printf("Error [getLotteryIdList]: %v", err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	ids := make([]int, 0)

	if err != nil {
		log.Printf("Error [getLotteryIdList]: %v", err)
		return &ids, err
	}

	defer rows.Close()

	var id int

	for rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		ids = append(ids, id)
	}

	return &ids, nil
}
