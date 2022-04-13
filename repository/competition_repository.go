package repository

import (
	"database/sql"
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/view"
)

type CompetitionRepository struct {
	db *sql.DB
}

func NewCompetitionRepository(db *sql.DB) *CompetitionRepository {
	return &CompetitionRepository{
		db: db,
	}
}

func (cr *CompetitionRepository) InsertCompetitions(competitions *[]domain.Competition) error {
	stmt, err := cr.db.Prepare(
		`INSERT INTO competition(id, name, logo, type, country)
		 VALUES(?, ?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE 
		 	logo = coalesce(VALUES(logo), logo), 
			type = coalesce(VALUES(type), type), 
			name = coalesce(VALUES(name), name)`)

	if err != nil {
		log.Printf("Error [InsertCompetitions]: %v", err)
		return err
	}

	defer stmt.Close()

	for _, competition := range *competitions {
		_, err = stmt.Exec(competition.Id, competition.Name, competition.Logo, competition.Type, competition.Country)

		if err != nil {
			log.Printf("Error [InsertCompetitions]: %v - [%v, %v, %v, %v, %v]", err, competition.Id, competition.Name, competition.Logo, competition.Type, competition.Country)
			continue
		}

		cr.insertSeasons(competition.Id, competition.Seasons)
	}

	return nil
}

func (cr *CompetitionRepository) insertSeasons(competitionId int, seasons *[]domain.Season) error {
	stmt, err := cr.db.Prepare(
		`INSERT INTO season(competition_id, year, name, ended)
		 VALUES(?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE 
		 	year = coalesce(VALUES(year), year), 
			name = coalesce(VALUES(name), name), 
			ended = coalesce(VALUES(ended), ended)`)

	if err != nil {
		log.Printf("Error [insertSeasons]: %v - [%v]", err, competitionId)
		return err
	}

	defer stmt.Close()

	for _, season := range *seasons {
		if (domain.Season{}) != season {
			_, err = stmt.Exec(competitionId, season.Year, season.Name, season.Ended)

			if err != nil {
				log.Printf("Error: %v - [%v, %v, %v, %v]", err, competitionId, season.Year, season.Name, season.Ended)
			}
		}
	}

	return nil
}

func (cr *CompetitionRepository) GetCompetitions(country string, year int, ended bool) (*[]domain.Competition, error) {
	stmt, err := cr.db.Prepare(
		`SELECT c.id, c.name, s.year, s.ended
		 FROM competition c
		 JOIN season s ON s.competition_id = c.id
		 WHERE c.country = ? AND s.ended = ? AND s.year = ?
		 ORDER BY 1, 2`)
	competitions := make([]domain.Competition, 0)

	if err != nil {
		log.Printf("Error [GetCompetitions]: %v - [%v %v]", err, country, year)
		return &competitions, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(country, ended, year)

	if err != nil {
		log.Printf("Error [GetCompetitions]: %v - [%v, %v, %v]", err, country, year, ended)
		return &competitions, err
	}

	defer rows.Close()

	competition := domain.Competition{}
	var competitionId int = 0
	var competitionName string

	for rows.Next() {
		season := domain.Season{}

		err = rows.Scan(&competitionId, &competitionName, &season.Year, &season.Ended)

		if err != nil {
			log.Printf("Error [GetCompetitions]: %v", err)
			continue
		}

		if competition.Id != competitionId {
			seasons := make([]domain.Season, 0)
			competition = domain.Competition{
				Id:      competitionId,
				Name:    competitionName,
				Seasons: &seasons,
			}

			competitions = append(competitions, competition)
		}

		*competition.Seasons = append(*competition.Seasons, season)
	}

	return &competitions, err
}
func (cr *CompetitionRepository) GetCompetition(competitionId int, year int) (*domain.Competition, error) {
	stmt, err := cr.db.Prepare(
		`SELECT c.id, c.name, s.year, s.ended
		 FROM competition c
		 JOIN season s ON s.competition_id = c.id
		 WHERE c.id = ? AND s.year = ?
		 ORDER BY 1, 2`)

	if err != nil {
		log.Printf("Error [GetCompetition]: %v - [%v, %v]", err, competitionId, year)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(competitionId, year)

	if err != nil {
		log.Printf("Error [GetCompetition]: %v - [%v, %v]", err, competitionId, year)
		return nil, err
	}

	defer rows.Close()

	competition := domain.Competition{}
	var competitionName string

	for rows.Next() {
		season := domain.Season{}
		err = rows.Scan(&competitionId, &competitionName, &season.Year, &season.Ended)

		if err != nil {
			log.Printf("Error [GetCompetition]: %v", err)
			continue
		}

		seasons := make([]domain.Season, 0)
		competition = domain.Competition{
			Id:      competitionId,
			Name:    competitionName,
			Seasons: &seasons,
		}

		*competition.Seasons = append(*competition.Seasons, season)
	}

	return &competition, err
}

func (cr *CompetitionRepository) GetTeamStats(competitionId int, year int, teamId int) (*view.TeamStats, error) {
	stmt, err := cr.db.Prepare(
		`SELECT
			count(*) M,
			count(case when 
			  (m.home_score > m.away_score AND m.home_id = ?) OR
		  	  (m.away_score > m.home_score AND m.away_id = ?) then 1 else null end) as W,
			count(case when m.home_score = m.away_score then 1 else null end) as D,
			count(case when 
		  	  (m.home_score < m.away_score AND m.home_id = ?) OR
		      (m.away_score < m.home_score AND m.away_id = ?) then 1 else null end) as L,
			sum(case when m.home_id = ? then m.home_score else m.away_score end) as GP,
			sum(case when m.home_id = ? then m.away_score else m.home_score end) as GC
	  	 FROM round r
	  	 JOIN` + " `match` " + `m ON m.round_id = r.id
	  	 WHERE r.competition_id = ? AND r.year = ? 
			AND (m.home_id = ? OR m.away_id = ?) AND m.home_score IS NOT NULL`)

	if err != nil {
		log.Printf("Error [GetCompetition]: %v - [%v, %v]", err, competitionId, year)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(teamId, teamId, teamId, teamId, teamId, teamId, competitionId, year, teamId, teamId)

	if err != nil {
		log.Printf("Error [GetTeamStats]: %v - [%v, %v %v]", err, competitionId, year, teamId)
		return nil, err
	}

	defer rows.Close()

	teamStats := view.TeamStats{}

	if rows.Next() {
		err = rows.Scan(&teamStats.M, &teamStats.W, &teamStats.D, &teamStats.L, &teamStats.GP, &teamStats.GC)

		if err != nil {
			log.Printf("Error [GetTeamStats]: %v", err)
			return nil, err
		}

		teamStats.TeamId = teamId
		teamStats.SG = teamStats.GP - teamStats.GC
	}

	return &teamStats, err
}
