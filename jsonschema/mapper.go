package jsonschema

var artist = map[string]interface{}{
	"artiste":     new(GetArtisteAllAPISchema),
	"artiste/:id": new(GetArtisteAPISchema),
}

func GetSchemaMapper() map[string]interface{} {
    allSchemaMapper := make(map[string]interface{})
	allSchemaList := []map[string]interface{}{artist}

	for _, schem := range allSchemaList {
        for k, v := range schem {
            allSchemaMapper[k] = v
        }
	}
	return allSchemaMapper
}

func GetSchemaMapperForDocument() []map[string]interface{} {
	dSchemaMapperList := []map[string]interface{}{
		map[string]interface{}{"artiste": artist},
	}
	return dSchemaMapperList
}
