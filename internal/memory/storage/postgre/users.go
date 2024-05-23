package storage

import (
	"database/sql"
	"time"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) CreateUser(user *entity.User) (int, error) {
	var eventId int
	date, err := time.Parse("2006-01-02", user.DateBirthday)

	if err != nil {
		return 0, err
	}

	query := `insert into users(first_name, middle_name, last_name, date_birth, phone_number, account_id, age_categoty_id) values($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, user.FirstName, user.MiddelName, user.LastName, date, user.PhoneNumber, user.AccountId, user.AgeCategory)
	err = result.Scan(&eventId)

	if err != nil {
		return 0, err
	}

	return eventId, nil
}

func (s *Storage) GetUserById(id int) (*entity.User, error) {
	var user entity.User
	query := `select id, account_id, first_name, middle_name, last_name, date_birth, phone_number from users where id = $1`
	row := s.ConnectionDB.QueryRow(query, id)

	err := row.Scan(
		&user.Id,
		&user.AccountId,
		&user.FirstName,
		&user.MiddelName,
		&user.LastName,
		&user.DateBirthday,
		&user.PhoneNumber,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return &user, err
	}

	return &user, nil
}

func (s *Storage) GetUserByIdAccount(id int) (*entity.User, error) {
	var user entity.User
	query := `select id, account_id, first_name, middle_name, last_name, date_birth, phone_number, age_categoty_id, gener_ids from users where account_id = $1`
	row := s.ConnectionDB.QueryRow(query, id)

	err := row.Scan(
		&user.Id,
		&user.AccountId,
		&user.FirstName,
		&user.MiddelName,
		&user.LastName,
		&user.DateBirthday,
		&user.PhoneNumber,
		&user.AgeCategory,
		&user.GenersIds,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return &user, err
	}

	return &user, nil
}

func (s *Storage) DeleteUserById(id int) error {
	query := `delete from users where id = $1`
	_, err := s.ConnectionDB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateUser(user *entity.User) error {
	query := `update users set first_name=$1, middle_name=$2, last_name=$3, date_birth=$4, phone_number=$5 where id=$6`
	_, err := s.ConnectionDB.Exec(query, user.FirstName, user.MiddelName, user.LastName, user.DateBirthday, user.PhoneNumber, user.Id)

	if err != nil {
		return err
	}

	return nil
}
