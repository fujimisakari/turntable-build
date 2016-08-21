package jsonschema

func GetSchemaMapper() map[string]interface{} {
	mapper := map[string]interface{}{
		"/api/artiste":     new(GetArtisteAllAPISchema),
		"/api/artiste/:id": new(GetArtisteAPISchema),
	}
	return mapper
}
