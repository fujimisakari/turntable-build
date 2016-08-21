package model

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type (
	ArtisteMap map[string]interface{}
	Artistes   []Artiste

	Artiste    struct {
		ID     int64  `db:"id"`
		Name   string `db:"name"`
		TeamID int64  `db:"team_id"`
	}
)

func (t Artiste) GetModelSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Artiste item",
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

func (t *Artiste) LoadByID(tx *dbr.Tx, id int64) error {
	return tx.Select("*").
		From("artiste").
		Where("id = ?", id).
		LoadStruct(t)
}

func (t *Artistes) LoadByIDs(tx *dbr.Tx, ids []int64) error {
	return tx.Select("*").
		From("artiste").
		Where("id IN ?", ids).
		LoadStruct(t)
}

func (ts *Artistes) LoadAll(tx *dbr.Tx) error {
	return tx.Select("*").
		From("artiste").
		LoadStruct(ts)
}

func (t *Artiste) ToMap() map[string]interface{} {
	mapData := map[string]interface{}{
		"id":   t.ID,
		"name": t.Name,
	}
	return mapData
}

func (ts *Artistes) ToMapList() []ArtisteMap {
	mapList := make([]ArtisteMap, len(*ts))
	for i, t := range *ts {
		mapList[i] = t.ToMap()
	}
	return mapList
}
