package app

import "github.com/SashaMelva/buffer_data_to_database/internal/entity"

func (a *App) CreateFact(facts entity.Facts) error {

	list, err := a.storage.GetChildrens(id)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return list, nil
}
