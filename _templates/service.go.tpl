/*

This code has been created automatically

*/
package domain

import (
    "github.com/fujimisakari/turntable-build/model"
    "github.com/gocraft/dbr"
    "github.com/pkg/errors"
)

func Get{{ .UpperName }}(tx *dbr.Tx, id int64) (*model.{{ .UpperName }}, error) {
    {{ .LowerName }} := new(model.{{ .UpperName }})
    if err := {{ .LowerName }}.LoadByID(tx, id); err != nil {
        return nil, errors.Errorf("{{ .UpperName }} does not exists: %d", id)
    }
    return {{ .LowerName }}, nil
}

func Get{{ .UpperName }}All(tx *dbr.Tx) (*model.{{ .UpperName }}s, error) {
    {{ .LowerName }}s := new(model.{{ .UpperName }}s)
    if err := {{ .LowerName }}s.LoadAll(tx); err != nil {
        return nil, errors.New("{{ .UpperName }}s does not exists")
    }
    return {{ .LowerName }}s, nil
}
