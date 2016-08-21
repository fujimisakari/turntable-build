package model

import (
	"github.com/gocraft/dbr"
)

type (
	TeamMap map[string]interface{}
	Teams []Team

	Team  struct {
		ID   int64  `db:"id"`
		Name string `db:"name"`
	}
)

func (t Team) GetModelSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Team item",
		"description": "",
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

func (t *Team) ToMap() map[string]interface{} {
	mapData := map[string]interface{}{
		"id":   t.ID,
		"name": t.Name,
	}
	return mapData
}

func (t *Team) LoadByID(tx *dbr.Tx, id int64) error {
	return tx.Select("*").
		From("team").
		Where("id = ?", id).
		LoadStruct(t)
}

func (t *Team) LoadByIDs(tx *dbr.Tx, ids []int64) error {
	return tx.Select("*").
		From("team").
		Where("id IN ?", ids).
		LoadStruct(t)
}

func (ts *Teams) LoadAll(tx *dbr.Tx) error {
	return tx.Select("*").
		From("Team").
		LoadStruct(ts)
}

func (ts *Teams) ToMapList() []TeamMap {
	mapList := make([]TeamMap, len(*ts))
	for i, t := range *ts {
		mapList[i] = t.ToMap()
	}
	return mapList
}
