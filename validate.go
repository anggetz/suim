package suim

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var GoCustomValidator string

func Validate(obj interface{}) error {
	objMeta, fields, err := ObjToFields(obj)
	if err != nil {
		return fmt.Errorf("fail reading meta data. %s", err.Error())
	}

	errorTexts := []string{}
	for _, f := range fields {
		fieldErr := validateField(obj, f)
		if fieldErr != nil {
			errorTexts = append(errorTexts, fmt.Sprintf("%s: %s", f.Field, fieldErr.Error()))
		}
	}

	if len(errorTexts) > 0 {
		return errors.New(strings.Join(errorTexts, " | "))
	}

	if objMeta.GoCustomValidator != "" {
		rv := reflect.ValueOf(obj)
		if _, hasFn := rv.Type().MethodByName(objMeta.GoCustomValidator); hasFn {
			mtd := rv.MethodByName(objMeta.GoCustomValidator)
			outs := mtd.Call([]reflect.Value{rv})
			if len(outs) > 0 {
				if outs[0].Type().String() == "error" {
					if rvErr := outs[0]; !rvErr.IsZero() {
						return fmt.Errorf("custom validator error. %s", rvErr.Interface().(error))
					}
				}
			}
		}
	}

	return nil
}

func validateField(obj interface{}, fm Field) error {
	v, has := getValue(obj, fm.Field)
	if has {
		/*
			rvMain := reflect.ValueOf(v)
			isPtr := rvMain.Kind() == reflect.Ptr
			rvElem := rvMain
			if isPtr {
				rvElem = rvMain.Elem()
			}
		*/

		switch fm.DataType {
		case "string":
			objStr := v.String()
			if objStr == "" && fm.Form.Required {
				return errors.New("could not be nil or empty")
			}

			if fm.Form.MinLength > 0 {
				if len(objStr) < fm.Form.MinLength {
					return fmt.Errorf("min length is %d", fm.Form.MinLength)
				}
			}

			if fm.Form.MaxLength > 0 && len(objStr) > fm.Form.MaxLength {
				return fmt.Errorf("max length is %d", fm.Form.MaxLength)
			}

			// useList and not using lookup
			if fm.Form.UseList && fm.Form.LookupUrl == "" {
				found := false
			itemLoop:
				for _, item := range fm.Form.Items {
					if item.Text == objStr {
						found = true
						break itemLoop
					}
				}

				if !found {
					return fmt.Errorf("")
				}
			}

		default:
			//-- do nothing
		}
	}
	return nil
}

func getValue(obj interface{}, name string) (reflect.Value, bool) {
	rv := reflect.Indirect(reflect.ValueOf(obj))

	if rv.Kind() == reflect.Struct {
		f := rv.FieldByName(name)
		return f, true
	} else if rv.Kind() == reflect.Map {
		return rv.MapIndex(reflect.ValueOf(name)), true
	}

	return reflect.Value{}, false
}
