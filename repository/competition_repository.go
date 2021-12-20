package repository

import (
	"database/sql"

	"github.com/giancarlobastos/loteca-backend/domain"
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
	stmt, err := cr.db.Prepare("INSERT IGNORE INTO competition(id, name, logo, type, country) VALUES(?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, competition := range *competitions {
		stmt.Exec(competition.Id, competition.Name, competition.Logo, competition.Type, competition.Country)
		cr.insertSeasons(competition.Id, competition.Seasons)
	}

	return nil
}

func (cr *CompetitionRepository) insertSeasons(competitionId uint32, seasons *[]domain.Season) error {
	stmt, err := cr.db.Prepare("INSERT IGNORE INTO season(competition_id, year, name, ended) VALUES(?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, season := range *seasons {
		if (domain.Season{}) != season {
			stmt.Exec(competitionId, season.Year, season.Name, season.Ended)
		}
	}

	return nil
}

func (cr *CompetitionRepository) GetCompetitions(country string) (*[]domain.Competition, error) {
	stmt, err := cr.db.Prepare("SELECT c.id, s.year, s.ended " +
		"FROM competition c " +
		"JOIN season s ON s.competition_id = c.id " +
		"WHERE c.country = ? " +
		"ORDER BY 1, 2")
	competitions := make([]domain.Competition, 0)

	if err != nil {
		return &competitions, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(country)

	if err != nil {
		return &competitions, err
	}

	defer rows.Close()

	competition := domain.Competition{}
	var competitionId uint32 = 0

	for rows.Next() {
		season := domain.Season{}
		rows.Scan(&competitionId, &season.Year, &season.Ended)

		if competition.Id != competitionId {
			seasons := make([]domain.Season, 0)
			competition = domain.Competition{
				Id:      competitionId,
				Seasons: &seasons,
			}

			competitions = append(competitions, competition)
		}

		*competition.Seasons = append(*competition.Seasons, season)
	}

	return &competitions, err
}
