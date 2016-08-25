/*

This code has been created automatically

*/
package domain

import (
	"github.com/fujimisakari/turntable-build/model"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

func GetTeam(tx *dbr.Tx, id int64) (*model.Team, error) {
	team := new(model.Team)
	if err := team.LoadByID(tx, id); err != nil {
		return nil, errors.Errorf("Team does not exists: %d", id)
	}
	return team, nil
}

func GetTeamAll(tx *dbr.Tx) (*model.Teams, error) {
	teams := new(model.Teams)
	if err := teams.LoadAll(tx); err != nil {
		return nil, errors.New("Teams does not exists")
	}
	return teams, nil
}
