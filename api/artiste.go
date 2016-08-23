package api

import (
	"strconv"

	_ "github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"

	dm_art "turntable-build/domain/artiste"
	dm_team "turntable-build/domain/team"
)

func GetArtiste() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
		tx := c.Get("Tx").(*dbr.Tx)

		artiste, err := dm_art.GetArtiste(tx, id)
		if err != nil {
			return err
		}
		team, err := dm_team.GetTeam(tx, artiste.TeamID)
		if err != nil {
			return err
		}

		context := map[string]interface{}{
			"artiste": artiste.ToMap(),
			"team":    team.ToMap(),
		}
		c.Set("context", context)
		return nil
	}
}

func GetArtisteAll() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		artistes, err := dm_art.GetArtisteAll(tx)
		if err != nil {
			return err
		}

		context := map[string]interface{}{
			"artistes": artistes.ToMapList(),
		}
		c.Set("context", context)
		return nil
	}
}
