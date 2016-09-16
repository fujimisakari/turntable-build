/*

This code has been created automatically

*/
package model

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type (
	ArtisteMap map[string]interface{}
	Artistes   []Artiste

	Artiste struct {
		ID     int64  `db:"id"`
		Name   string `db:"name"`
		TeamID int64  `db:"team_id"`
	}
)

func (a Artiste) ModelSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Artiste item",
		"description": "Artiste info api",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":    "integer",
				"example": "1",
			},
			"name": map[string]interface{}{
				"type":    "string",
				"example": "Q-bert",
			},
		},
	}
	return schema
}

func (a *Artiste) LoadByID(tx *dbr.Tx, id int64) error {
	return tx.Select("*").
		From("artiste").
		Where("id = ?", id).
		LoadStruct(a)
}

func (a *Artistes) LoadByIDs(tx *dbr.Tx, ids []int64) error {
	return tx.Select("*").
		From("artiste").
		Where("id IN ?", ids).
		LoadStruct(a)
}

func (as *Artistes) LoadAll(tx *dbr.Tx) error {
	return tx.Select("*").
		From("artiste").
		LoadStruct(as)
}

func (a *Artiste) ToMap() map[string]interface{} {
	mapData := map[string]interface{}{
		"id":   a.ID,
		"name": a.Name,
	}
	return mapData
}

func (as *Artistes) ToMapList() []ArtisteMap {
	mapList := make([]ArtisteMap, len(*as))
	for i, a := range *as {
		mapList[i] = a.ToMap()
	}
	return mapList
}
