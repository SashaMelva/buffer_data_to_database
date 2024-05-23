package storage

import (
	"database/sql"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) SaveTokens(tokens *entity.RefreshToken) error {

	query := `UPDATE token SET token = $1 where account_id = $2`
	_, err := s.ConnectionDB.Exec(query, tokens.RefreshToken, tokens.UserId)

	if err != nil {
		return err
	}

	return nil
}
func (s *Storage) GetTokenByUser(id int) (*entity.RefreshToken, error) {
	var token entity.RefreshToken
	query := `select token, account_id from token where account_id = $1`
	row := s.ConnectionDB.QueryRow(query, id)

	err := row.Scan(
		&token.RefreshToken,
		&token.UserId,
	)

	s.log.Debug(token)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return &token, err
	}

	return &token, nil
}

func (s *Storage) CreateToken(userId int) error {
	var childrenId int
	query := `insert into token (account_id, token) values($1, '')`
	s.log.Debug(query, userId)
	_, err := s.ConnectionDB.Exec(query, userId)

	s.log.Debug(childrenId)
	if err != nil {
		return err
	}

	return nil
}
