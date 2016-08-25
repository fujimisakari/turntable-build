/*

This code has been created automatically

*/
package domain

import (
	"github.com/fujimisakari/turntable-build/model"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

func GetTournament(tx *dbr.Tx, id int64) (*model.Tournament, error) {
	tournament := new(model.Tournament)
	if err := tournament.LoadByID(tx, id); err != nil {
		return nil, errors.Errorf("Tournament does not exists: %d", id)
	}
	return tournament, nil
}

func GetTournamentAll(tx *dbr.Tx) (*model.Tournaments, error) {
	tournaments := new(model.Tournaments)
	if err := tournaments.LoadAll(tx); err != nil {
		return nil, errors.New("Tournaments does not exists")
	}
	return tournaments, nil
}
