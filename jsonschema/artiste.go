package jsonschema

import "github.com/fujimisakari/turntable-build/model"

type (
	GetArtisteAPISchema struct {}
	GetArtisteAllAPISchema struct {}
)

func (s GetArtisteAPISchema)GetRequestSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"title":       "Artiste Request",
		"description": "Artiste Request API description",
		"method":      "GET",
		"type":        "object",
		"properties": map[string]interface{}{
			"id": map[string]string{
				"type": "integer",
				"example": "1",
				"description": "Artiste ID",
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
		"description": "Artiste Response API description",
		"type":        "object",
		"properties": map[string]interface{}{
			"artiste": artiste.ModelSchema(),
			"team": team.ModelSchema(),
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
		"title":       "ArtisteAll Request",
		"description": "ArtisteAll Request API description",
		"method":      "GET",
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
		"description": "ArtisteAll Response API description",
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
