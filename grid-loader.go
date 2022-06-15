package suim

import (
	"reflect"

	"git.kanosolution.net/koloni/crowd"
)

func CreateGridConfig(obj interface{}) (*GridConfig, error) {
	name := reflect.Indirect(reflect.ValueOf(obj)).Type().Name()
	cfg, ok := gridConfigs[name]
	if ok {
		return cfg, nil
	}

	meta, fields, err := ObjToFields(obj)
	if err != nil {
		return nil, err
	}

	cfg = new(GridConfig)
	cfg.Setting = meta.Grid

	for idx, f := range fields {
		if f.Grid.Pos == 0 {
			f.Grid.Pos = 1000
		}
		fields[idx] = f
	}

	var cfgFields []GridField
	err = crowd.FromSlice(fields).Filter(func(f Field) bool {
		return f.GridElement != "hide"
	}).Map(func(f Field) GridField {
		return f.Grid
	}).Sort(func(f1, f2 GridField) bool {
		return f1.Pos < f2.Pos
	}).Collect().Run(&cfgFields)
	if err != nil {
		return nil, err
	}
	cfg.Fields = cfgFields

	return cfg, nil
}
