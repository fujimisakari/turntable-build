package domain

import (
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"

	"turntable-build/model"
)

func GetArtiste(tx *dbr.Tx, id int64) (*model.Artiste, error) {
	artiste := new(model.Artiste)
	if err := artiste.LoadByID(tx, id); err != nil {
		return nil, errors.Errorf("Artiste does not exists: %d", id)
	}
	return artiste, nil
}

func GetArtisteAll(tx *dbr.Tx) (*model.Artistes, error) {
	artistes := new(model.Artistes)
	if err := artistes.LoadAll(tx); err != nil {
		return nil, errors.New("Artistes does not exists")
	}
	return artistes, nil
}
