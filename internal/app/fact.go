package app

import "github.com/SashaMelva/buffer_data_to_database/internal/entity"

func (a *App) CreateFacts(facts *entity.Facts) error {

	for i := range facts.Facts {
		a.buf.Add(facts.Facts[i])
		a.log.Debug("put fact: ", facts.Facts[i])
	}

	return nil
}

func (a *App) SaveFact(fact *entity.Fact) (bool, error) {
	err := a.Storage.AddFact(fact)

	if err != nil {
		return false, err
	}

	return true, nil
}
