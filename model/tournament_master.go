/*

This code has been created automatically

*/
package model

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type (
	TournamentMap map[string]interface{}
	Tournaments   []Tournament

	Tournament struct {
		ID   int64  `db:"id"`
		Name string `db:"name"`
	}
)

func (t Tournament) ModelSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Tournament item",
		"description": "Tournament info api",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]string{
				"type": "integer",
			},
			"name": map[string]string{
				"type": "string",
			},
		},
	}
	return schema
}

func (t *Tournament) LoadByID(tx *dbr.Tx, id int64) error {
	return tx.Select("*").
		From("tournament").
		Where("id = ?", id).
		LoadStruct(t)
}

func (t *Tournaments) LoadByIDs(tx *dbr.Tx, ids []int64) error {
	return tx.Select("*").
		From("tournament").
		Where("id IN ?", ids).
		LoadStruct(t)
}

func (ts *Tournaments) LoadAll(tx *dbr.Tx) error {
	return tx.Select("*").
		From("tournament").
		LoadStruct(ts)
}

func (t *Tournament) ToMap() map[string]interface{} {
	mapData := map[string]interface{}{
		"id":   t.ID,
		"name": t.Name,
	}
	return mapData
}

func (ts *Tournaments) ToMapList() []TournamentMap {
	mapList := make([]TournamentMap, len(*ts))
	for i, t := range *ts {
		mapList[i] = t.ToMap()
	}
	return mapList
}
