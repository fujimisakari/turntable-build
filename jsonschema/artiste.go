package jsonschema

import "github.com/fujimisakari/turntable-build/model"

type (
	GetArtisteAPISchema struct {}
	GetArtisteAllAPISchema struct {}
)

func (s GetArtisteAPISchema)GetRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Artiste Request",
		"description": "",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]string{
				"type": "integer",
			},
		},
		"required": []string{
			"id",
		},
		"additionalProperties": false,
	}
	return schema
}

func (s GetArtisteAPISchema)GetResponseSchema() map[string]interface{} {
	artiste := new(model.Artiste)
	team := new(model.Team)
	schema := map[string]interface{}{
		"title":       "Artiste Response",
		"description": "",
		"type":        "object",
		"properties": map[string]interface{}{
			"artiste": map[string]interface{}{
				"type":  "object",
				"items": artiste.ModelSchema(),
			},
			"team": map[string]interface{}{
				"type":  "object",
				"items": team.ModelSchema(),
			},
		},
		"required": []string{
			"artiste", "team",
		},
		"additionalProperties": false,
	}
	return schema
}

func (s GetArtisteAllAPISchema)GetRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Artiste Request",
		"description": "",
		"type":        "object",
		"properties": map[string]interface{}{},
		"required": []string{},
		"additionalProperties": false,
	}
	return schema
}

func (s GetArtisteAllAPISchema)GetResponseSchema() map[string]interface{} {
	artiste := new(model.Artiste)
	schema := map[string]interface{}{
		"title":       "ArtisteAll Response",
		"description": "",
		"type":        "object",
		"properties": map[string]interface{}{
			"artistes": map[string]interface{}{
				"type":  "array",
				"items": artiste.ModelSchema(),
			},
		},
		"required": []string{
			"artistes",
		},
		"additionalProperties": false,
	}
	return schema
}
