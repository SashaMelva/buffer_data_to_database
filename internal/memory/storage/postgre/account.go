package storage

import (
	"database/sql"

	"github.com/SashaMelva/buffer_data_to_database/internal/entity"
)

func (s *Storage) ChekLoginAccount(account *entity.Account) (int, error) {
	var accountId int
	query := `select id from account where login = $1`
	result := s.ConnectionDB.QueryRow(query, account.Login)

	err := result.Scan(&accountId)

	if err != nil {
		return 0, nil
	}

	return accountId, nil
}

func (s *Storage) CreateAccount(account *entity.Account) (int, error) {
	var id int
	query := `insert into account(login, password, role) values($1, $2, $3) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, account.Login, account.Password, account.Role)
	err := result.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetAccountId(account *entity.Account) (*entity.AccountId, error) {
	var accId entity.AccountId
	query := `select id, login, password, role from account where login = $1 AND password = $2`
	row := s.ConnectionDB.QueryRow(query, account.Login, account.Password)

	err := row.Scan(&accId.Id, &accId.Login, &accId.Password, &accId.Role)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return &accId, err
	}

	return &accId, nil
}
