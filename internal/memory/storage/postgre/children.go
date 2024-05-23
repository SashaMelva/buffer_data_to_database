package storage

import (
	"strings"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) CreateChildren(idParent, children int) error {
	var childrenId int
	s.log.Debug(idParent, children)
	query := `insert into children(user_parent_id, user_id) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, idParent, children)
	err := result.Scan(&childrenId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetChildrens(idParent int) (*entity.ChilgrenLists, error) {
	var children entity.ChilgrenLists
	query := `SELECT user_id, first_name, middle_name, last_name, account_id  FROM public.children
		INNER JOIN users ON children.user_id = users.account_id
		WHERE user_parent_id = $1`
	rows, err := s.ConnectionDB.Query(query, idParent)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ch := entity.Chilgren{}

		if err := rows.Scan(
			&ch.IdUser,
			&ch.FirstName,
			&ch.MiddelName,
			&ch.LastName,
			&ch.AccountId,
		); err != nil {
			return nil, err
		}

		children.Chilgrens = append(children.Chilgrens, &ch)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &children, nil
}

func (s *Storage) DeleteChildren(childrenId int) error {
	query := `delete from children where user_id = $1`
	_, err := s.ConnectionDB.Exec(query, childrenId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetFiltersByChaild(idParent int) (*entity.Fileters, error) {
	var children entity.Fileters
	query := `SELECT id, name from filters Where account_id = $1`
	rows, err := s.ConnectionDB.Query(query, idParent)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ch := entity.Fileter{}

		if err := rows.Scan(
			&ch.Id,
			&ch.Name,
		); err != nil {
			return nil, err
		}

		children.Filters = append(children.Filters, &ch)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &children, nil
}

func (s *Storage) UpdateFiltersByChaild(filters *entity.FilterGeners) error {
	str := strings.Join(filters.IdGeners, ",")
	query := `update users set gener_ids=$1  where account_id=$6`
	_, err := s.ConnectionDB.Exec(query, str, filters.IdChildren)

	if err != nil {
		return err
	}

	return nil
}
