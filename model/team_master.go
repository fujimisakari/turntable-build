/*

This code has been created automatically

*/
package model

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type (
	TeamMap map[string]interface{}
	Teams   []Team

	Team struct {
		ID   int64  `db:"id"`
		Name string `db:"name"`
	}
)

func (t Team) ModelSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Team item",
		"description": "Team info api",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":    "integer",
				"example": "1",
			},
			"name": map[string]interface{}{
				"type":    "string",
				"example": "Invisible skrach piklz",
			},
		},
	}
	return schema
}

func (t *Team) LoadByID(tx *dbr.Tx, id int64) error {
	return tx.Select("*").
		From("team").
		Where("id = ?", id).
		LoadStruct(t)
}

func (t *Teams) LoadByIDs(tx *dbr.Tx, ids []int64) error {
	return tx.Select("*").
		From("team").
		Where("id IN ?", ids).
		LoadStruct(t)
}

func (ts *Teams) LoadAll(tx *dbr.Tx) error {
	return tx.Select("*").
		From("team").
		LoadStruct(ts)
}

func (t *Team) ToMap() map[string]interface{} {
	mapData := map[string]interface{}{
		"id":   t.ID,
		"name": t.Name,
	}
	return mapData
}

func (ts *Teams) ToMapList() []TeamMap {
	mapList := make([]TeamMap, len(*ts))
	for i, t := range *ts {
		mapList[i] = t.ToMap()
	}
	return mapList
}
