/*

This code has been created automatically

*/
package model

import (
    _ "github.com/Sirupsen/logrus"
    "github.com/gocraft/dbr"
)

type (
    {{ .UpperName }}Map map[string]interface{}
    {{ .UpperName }}s   []{{.UpperName}}

    {{ .UpperName }}    struct {
    {{ range $i, $p := .ModelProperties }}{{ $p.Name }}  {{ $p.Type }} `db:"{{ $p.DBColumn }}"`
    {{ end }}
    }
)

func ({{ .Self }} {{ .UpperName }}) ModelSchema() map[string]interface{} {
    schema := map[string]interface{}{
        "title":       "{{ .UpperName }} item",
        "description": "{{ .Description }}",
        "type":        "object",
        "properties": map[string]interface{}{
            {{ range $i, $p := .JsonSchemaProperties }}"{{ $p.Name }}": map[string]interface{}{
            "type": "{{ $p.Type }}",
            "example": "{{ $p.Example }}",
            },
            {{ end }}
        },
    }
    return schema
}

func ({{ .Self }} *{{ .UpperName }}) LoadByID(tx *dbr.Tx, id int64) error {
    return tx.Select("*").
        From("{{ .LowerName }}").
        Where("id = ?", id).
        LoadStruct({{ .Self }})
}

func ({{ .Self }} *{{ .UpperName }}s) LoadByIDs(tx *dbr.Tx, ids []int64) error {
    return tx.Select("*").
        From("{{ .LowerName }}").
        Where("id IN ?", ids).
        LoadStruct({{ .Self }})
}

func ({{ .Self }}s *{{ .UpperName }}s) LoadAll(tx *dbr.Tx) error {
    return tx.Select("*").
        From("{{ .LowerName }}").
        LoadStruct({{ .Self }}s)
}

func ({{ .Self }} *{{ .UpperName }}) ToMap() map[string]interface{} {
    mapData := map[string]interface{}{
        {{ range $i, $p := .MapSchemaProperties }}"{{ $p.JsonName }}": {{ $p.Self }}.{{ $p.ModelName }},
        {{ end }}
    }
    return mapData
}

func ({{ .Self }}s *{{ .UpperName }}s) ToMapList() []{{ .UpperName }}Map {
    mapList := make([]{{ .UpperName }}Map, len(*{{ .Self }}s))
    for i, {{ .Self }} := range *{{ .Self }}s {
        mapList[i] = {{ .Self }}.ToMap()
    }
    return mapList
}
