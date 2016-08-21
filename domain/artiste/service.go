package domain

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"

	"turntable-build/model"
)

func GetArtiste(tx *dbr.Tx, id int64) *model.Artiste {
	artiste := new(model.Artiste)
	if err := artiste.LoadByID(tx, id); err != nil {
		logrus.Debug(err)
		// return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
	}
	return artiste
}

func GetArtisteAll(tx *dbr.Tx) *model.Artistes {
	artistes := new(model.Artistes)
	if err := artistes.LoadAll(tx); err != nil {
		logrus.Debug(err)
		// return echo.NewHTTPError(fasthttp.StatusNotFound, "Member does not exists.")
	}
	return artistes
}
