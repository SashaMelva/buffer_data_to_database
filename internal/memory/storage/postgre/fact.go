package storage

import "github.com/SashaMelva/buffer_data_to_database/internal/entity"

func (s *Storage) AddFact(fact *entity.Fact) error {
	query := `INSERT INTO fact (period_start, period_end, comment) VALUES ($1::timestamp, $2::timestamp, $3)`
	_, err := s.ConnectionDB.Exec(query, fact.PeriodStart, fact.PeriodEnd, fact.Comment)

	if err != nil {
		return err
	}

	return nil
}
