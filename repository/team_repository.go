package repository

import (
	"database/sql"

	"github.com/giancarlobastos/loteca-backend/domain"
)

type TeamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

func (tr *TeamRepository) InsertTeams(teams *[]domain.Team) error {
	stmt, err := tr.db.Prepare("INSERT IGNORE INTO team(id, name, logo, country) VALUES(?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	stadiumStmt, err := tr.db.Prepare("INSERT IGNORE INTO stadium(id, name, city, state, country) VALUES(?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stadiumStmt.Close()

	for _, team := range *teams {
		stmt.Exec(team.Id, team.Name, team.Logo, team.Country)

		if (domain.Stadium{}) != *team.Stadium {
			stadiumStmt.Exec(team.Stadium.Id, team.Stadium.Name, team.Stadium.City, team.Stadium.State, team.Stadium.Country)
		}
	}

	return nil
}
