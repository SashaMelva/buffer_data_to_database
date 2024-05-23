package storage

import (
	"math"
	"time"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) AddHistory(videoId, children int, date time.Time) error {
	query := `insert into history (account_id, date_time, video_id) values($1, $2, $3)`
	_, err := s.ConnectionDB.Exec(query, children, date, videoId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetHistoryByCategoriesVideos(children int, date_start, date_end string) (*entity.ProcentByCategoresUser, error) {
	var categories entity.ProcentByCategoresUser
	s.log.Debug(children, date_start, date_end)
	query := `SELECT genre.name, Count(video_id) AS count_video FROM public.history
	INNER join video ON history.video_id = video.id
	INNER join genre ON  genre.id = video.gener_id
	WHERE history.date_time BETWEEN $1::timestamp AND $2::timestamp AND account_id = $3
	GROUP BY  genre.name
	Order BY 2 desc
	LIMIT 3`
	rows, err := s.ConnectionDB.Query(query, date_start, date_end, children)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ch := entity.ProcentByCategoryUser{}

		if err := rows.Scan(
			&ch.CategoryName,
			&ch.Count,
		); err != nil {
			return nil, err
		}

		categories.Category = append(categories.Category, &ch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	all := 0
	for i := range categories.Category {
		all += categories.Category[i].Count
	}
	s.log.Debug(all)

	for i := range categories.Category {
		s.log.Debug(i, float64(categories.Category[i].Count)/float64(all), all, all%categories.Category[i].Count)
		categories.Category[i].Procent = math.Round(float64(categories.Category[i].Count) / float64(all) * 100)
	}

	return &categories, nil
}

func (s *Storage) GetHistoryByResoursesVideos(children int) (*entity.ProcentByCategoresUser, error) {
	var categories entity.ProcentByCategoresUser
	s.log.Debug(children)
	query := `SELECT name_servis, Count(video_id) AS count_video FROM public.history
	INNER join video ON history.video_id = video.id
	WHERE account_id = $1
	GROUP BY  name_servis
	Order BY 2 desc
	LIMIT 3`
	rows, err := s.ConnectionDB.Query(query, children)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ch := entity.ProcentByCategoryUser{}

		if err := rows.Scan(
			&ch.CategoryName,
			&ch.Count,
		); err != nil {
			return nil, err
		}

		categories.Category = append(categories.Category, &ch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	all := 0
	for i := range categories.Category {
		all += categories.Category[i].Count
	}
	s.log.Debug(all)

	for i := range categories.Category {
		s.log.Debug(i, float64(categories.Category[i].Count)/float64(all), all, all%categories.Category[i].Count)
		categories.Category[i].Procent = math.Round(float64(categories.Category[i].Count) / float64(all) * 100)
	}

	return &categories, nil
}
