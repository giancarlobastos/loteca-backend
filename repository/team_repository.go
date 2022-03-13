package repository

import (
	"database/sql"
	"log"

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
	stmt, err := tr.db.Prepare(
		`INSERT INTO team(id, name, logo, country)
	 	 VALUES(?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE 
			name = coalesce(VALUES(name), name), 
			logo = coalesce(VALUES(logo), logo)`)

	if err != nil {
		log.Printf("Error [InsertTeams]: %v", err)
		return err
	}

	defer stmt.Close()

	stadiumStmt, err := tr.db.Prepare(
		`INSERT INTO stadium(id, name, city, state, country)
		 VALUES(?, ?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE 
			name = coalesce(VALUES(name), name), 
			city = coalesce(VALUES(city), city), 
			state = coalesce(VALUES(state), state), 
			country = coalesce(VALUES(country), country)`)

	if err != nil {
		log.Printf("Error [InsertTeams]: %v", err)
		return err
	}

	defer stadiumStmt.Close()

	for _, team := range *teams {
		_, err = stmt.Exec(team.Id, team.Name, team.Logo, team.Country)

		if err != nil {
			log.Printf("Error [InsertTeams]: %v - [%v %v %v %v]", err, team.Id, team.Name, team.Logo, team.Country)
			continue
		}

		if (domain.Stadium{}) != *team.Stadium {
			_, err = stadiumStmt.Exec(team.Stadium.Id, team.Stadium.Name, team.Stadium.City, team.Stadium.State, team.Stadium.Country)
			log.Printf("Error [InsertTeams]: %v - [%v %v %v %v %v]", err, team.Stadium.Id, team.Stadium.Name, team.Stadium.City, team.Stadium.State, team.Stadium.Country)
		}
	}

	return nil
}

func (tr *TeamRepository) GetTeams(country string) (*[]domain.Team, error) {
	stmt, err := tr.db.Prepare(
		`SELECT id, name, country
		 FROM team
		 WHERE country = ?
		 ORDER BY 2`)
	teams := make([]domain.Team, 0)

	if err != nil {
		log.Printf("Error [GetTeams]: %v", err)
		return &teams, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(country)

	if err != nil {
		log.Printf("Error [GetTeams]: %v - [%v]", err, country)
		return &teams, err
	}

	defer rows.Close()

	team := domain.Team{}

	for rows.Next() {
		err = rows.Scan(&team.Id, &team.Name, &team.Country)

		if err != nil {
			log.Printf("Error [GetTeams]: %v - [%v %v %v]", err, team.Id, team.Name, team.Country)
			continue
		}

		teams = append(teams, team)
	}

	return &teams, err
}
