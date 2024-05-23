package storage

import (
	"database/sql"
	"strings"

	"github.com/SashaMelva/smart_filter/internal/entity"
)

func (s *Storage) GetStatusVideos() (*entity.StatusVideos, error) {
	var videos entity.StatusVideos
	query := `SELECT id, name FROM public.starus_video`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		video := entity.StatusVideo{}

		if err := rows.Scan(
			&video.Id,
			&video.Name,
		); err != nil {
			return nil, err
		}

		videos.AgeCategores = append(videos.AgeCategores, &video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &videos, nil
}

func (s *Storage) GetAllAgeCategoryVideos() (*entity.AgeCategores, error) {
	var videos entity.AgeCategores
	query := `SELECT id, name FROM public.age_categoty`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		video := entity.AgeCategory{}

		if err := rows.Scan(
			&video.Id,
			&video.Name,
		); err != nil {
			return nil, err
		}

		videos.AgeCategores = append(videos.AgeCategores, &video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &videos, nil
}

func (s *Storage) ChekVideo(url string) (bool, error) {
	var count int
	query := `SELECT COUNT(id) FROM public.video WHERE url = $1`
	rows := s.ConnectionDB.QueryRow(query, url)

	err := rows.Scan(
		&count,
	)

	if err == sql.ErrNoRows {
		return false, err
	} else if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (s *Storage) GetVideoByUrl(url string) (*entity.VideoId, error) {
	var count entity.VideoId
	query := `SELECT id, url, age_categoty_id, status_id, gener_id 
	FROM video 
	WHERE url = $1 and status_id = 5`
	rows := s.ConnectionDB.QueryRow(query, strings.TrimSpace(url))
	s.log.Debug(rows)
	err := rows.Scan(
		&count.Id,
		&count.Url,
		&count.AgeCategoryId,
		&count.StatusId,
		&count.GenerId,
	)
	s.log.Debug(count)
	s.log.Debug(err)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return &count, nil
}

func (s *Storage) AddNewVideo(url string, name string) error {
	var eventId int
	query := `insert into video(url, status_id, name) values($1, $2, $3) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, url, 1, name)
	err := result.Scan(&eventId)

	s.log.Debug(eventId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateStatusVideo(video entity.VideoIdStatus) error {
	query := `update video set status_id = $1 WHERE id = $2`
	_, err := s.ConnectionDB.Exec(query, video.Status, video.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateVideo(video entity.VideoId) error {
	s.log.Debug(video)
	argStr := strings.Join(video.Tags, ",")
	query := `update video set
	url = $1, name = $2, age_categoty_id = $3, status_id = $4, tags = $5, gener_id = $6, lenguage = $7 WHERE id = $8`
	_, err := s.ConnectionDB.Exec(query, video.Url, video.Name, video.AgeCategoryId, video.StatusId, argStr, video.GenerId, video.Lenguage, video.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetFilterAgeCategory() (*entity.FilterAgeCategores, error) {
	var videos entity.FilterAgeCategores
	query := `SELECT id, name, category_id, params_eng, params_ru FROM filter_age_category`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		video := entity.FilterAgeCategore{}

		if err := rows.Scan(
			&video.Id,
			&video.Name,
			&video.Category_id,
			&video.Params_eng,
			&video.Params_ru,
		); err != nil {
			return nil, err
		}

		videos.Filters = append(videos.Filters, &video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &videos, nil
}
