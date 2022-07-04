package repository

import (
	"database/sql"
	"log"

	"github.com/giancarlobastos/loteca-backend/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) InsertUser(user *domain.User) (*domain.User, error) {
	stmt, err := ur.db.Prepare(
		`INSERT INTO user(name, facebook_id, device_id, photo, email) VALUES(?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error [InsertUser]: %v", err)
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.FacebookId, user.DeviceId, user.Picture, user.Email)

	if err != nil {
		log.Printf("Error [InsertUser]: %v - [%v %v %v %v %v]", err, user.Name, user.FacebookId, user.DeviceId, user.Picture, user.Email)
		return nil, err
	}

	_id, err := result.LastInsertId()

	if err != nil {
		log.Printf("Error [InsertUser.LastInsertId]: %v - [%v]", err, result)
		return nil, err
	}

	id := int(_id)
	user.Id = &id
	return user, nil
}

func (ur *UserRepository) GetUser(id int) (*domain.User, error) {
	return ur.getUser(
		`SELECT id, name, facebook_id, photo, email
		 FROM user
		 WHERE id = ?`, &id)
}

func (ur *UserRepository) GetUserByFacebookId(facebookId string) (*domain.User, error) {
	return ur.getUser(
		`SELECT id, name, facebook_id, NULL, photo, email
		 FROM user
		 WHERE facebook_id = ?`, &facebookId)
}

func (ur *UserRepository) GetUserByDeviceId(deviceId string) (*domain.User, error) {
	return ur.getUser(
		`SELECT id, name, NULL, device_id, photo, email
		 FROM user
		 WHERE device_id = ?`, &deviceId)
}

func (ur *UserRepository) getUser(query string, args ...interface{}) (*domain.User, error) {
	stmt, err := ur.db.Prepare(query)

	if err != nil {
		log.Printf("Error [GetUser]: %v - [%v]", err, query)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)

	if err != nil {
		log.Printf("Error [GetUser]: %v - [%v]", err, args)
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		user := domain.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.FacebookId, &user.DeviceId, &user.Picture, &user.Email)

		if err != nil {
			log.Printf("Error [GetUser]: %v - [%v]", err, args)
			return nil, err
		}

		return &user, nil
	}

	return nil, nil
}
