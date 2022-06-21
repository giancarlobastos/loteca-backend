package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/view"
)

type PollRepository struct {
	db *sql.DB
}

func NewPollRepository(db *sql.DB) *PollRepository {
	return &PollRepository{
		db: db,
	}
}

func (pr *PollRepository) GetPollResults(lotteryId int) (*view.PollResults, error) {
	stmt, err := pr.db.Prepare(
		`SELECT lp.lottery_id, lp.match_id, 
		 	count(case when lp.result = 'H' then 1 else null end) as home_votes,
		 	count(case when lp.result = 'D' then 1 else null end) as draw_votes,
		 	count(case when lp.result = 'A' then 1 else null end) as away_votes,
			count(DISTINCT lp.user_id) total_votes
		 FROM lottery_poll lp
		 JOIN lottery_match lm ON lp.lottery_id = lm.lottery_id AND lp.match_id = lm.match_id
		 WHERE lp.lottery_id = ?
		 GROUP BY lp.lottery_id, lp.match_id
		 ORDER BY lm.order`)

	if err != nil {
		log.Printf("Error [GetPollResults]: %v - [%v]", err, lotteryId)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(lotteryId)

	if err != nil {
		log.Printf("Error [GetPollResults]: %v - [%v]", err, lotteryId)
		return nil, err
	}

	defer rows.Close()

	results := view.PollResults{}
	votes := make([]view.Vote, 0)

	for rows.Next() {
		vote := view.Vote{}
		err = rows.Scan(&results.LotteryId, &vote.MatchId, &vote.HomeVotes, &vote.DrawVotes, &vote.AwayVotes, &results.TotalVotes)

		if err != nil {
			log.Printf("Error [GetPollResults]: %v", err)
			return nil, err
		}

		votes = append(votes, vote)
	}

	results.Votes = votes

	return &results, err
}

func (pr *PollRepository) GetVotes(matchId int) (*view.Vote, int, error) {
	stmt, err := pr.db.Prepare(
		`SELECT 
		    count(case when lp.result = 'H' then 1 else null end) as home_votes,
		 	count(case when lp.result = 'D' then 1 else null end) as draw_votes,
		 	count(case when lp.result = 'A' then 1 else null end) as away_votes,
			count(DISTINCT lp.user_id) total_votes
		 FROM lottery_poll lp
		 WHERE lp.match_id = ?
		 GROUP BY lp.match_id`)

	if err != nil {
		log.Printf("Error [GetVotes]: %v - [%v]", err, matchId)
		return nil, 0, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(matchId)

	if err != nil {
		log.Printf("Error [GetVotes]: %v - [%v]", err, matchId)
		return nil, 0, err
	}

	defer rows.Close()

	if rows.Next() {
		vote := view.Vote{
			MatchId: matchId,
		}

		totalVotes := 0

		err = rows.Scan(&vote.HomeVotes, &vote.DrawVotes, &vote.AwayVotes, &totalVotes)

		if err != nil {
			log.Printf("Error [GetVotes]: %v - [%v]", err, matchId)
			return nil, 0, err
		}

		return &vote, totalVotes, err
	}

	return nil, 0, err
}

func (lr *PollRepository) Vote(poll domain.Poll, user domain.User) error {
	tx, err := lr.db.BeginTx(context.Background(), nil)

	if err != nil {
		log.Printf("Error [Vote]: %v", err)
		return err
	}

	defer tx.Rollback()

	stmt, err := lr.db.Prepare(`INSERT INTO lottery_poll VALUES(?, ?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error [Vote]: %v - [%v %v]", err, poll, *user.Id)
		return err
	}

	defer stmt.Close()

	votedAt := time.Now().UTC()

	for _, vote := range poll.Votes {
		_, err = stmt.Exec(poll.LotteryId, vote.MatchId, *user.Id, vote.Result, votedAt)

		if err != nil {
			log.Printf("Error [CreateLottery]: %v - [%v, %v, %v, %v, %v]", err, poll.LotteryId, vote.MatchId, *user.Id, vote.Result, votedAt)
			return err
		}
	}

	err = tx.Commit()
	return err
}
