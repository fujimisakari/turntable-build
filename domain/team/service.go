package domain

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"

	"turntable-build/model"
)

func GetTeam(tx *dbr.Tx, id int64) *model.Team {
	team := new(model.Team)
	if err := team.LoadByID(tx, id); err != nil {
		logrus.Debug(err)
		// return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
	}
	return team
}

func GetTeamAll(tx *dbr.Tx) *model.Teams {
	teams := new(model.Teams)
	if err := teams.LoadAll(tx); err != nil {
		logrus.Debug(err)
		// return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
	}
	return teams
}
