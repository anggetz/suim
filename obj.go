package suim

import (
	"math"
	"reflect"

	"github.com/sebarcode/codekit"
)

/*
type SuimModel interface {
	HandleChange(fieldName string, v1, v2, vOld interface{})
}
*/

type ObjMeta struct {
	Grid               GridSetting
	Form               FormSetting
	GoCustomValidator  string
	HandleChangeFields []string
}

var (
	sectionGroups = map[string][]FormSectionGroup{}
)

func autoFormSections(obj interface{}) ([]FormSectionGroup, error) {
	v := reflect.Indirect(reflect.ValueOf(obj))
	typeString := v.Type().String()

	groups, has := sectionGroups[typeString]
	if has {
		return groups, nil
	}

	mt, fields, err := ObjToFields(obj)
	if err != nil {
		return nil, err
	}

	res := []FormSection{}
	lastSection := ""
	sectionCount := 0
	sectionNames := []string{}
	for _, f := range fields {
		if lastSection != f.Form.Section && !codekit.HasMember(sectionNames, f.Form.Section) {
			res = append(res, FormSection{Title: f.Form.Section, Name: f.Form.Section, AutoCol: 1})
			sectionNames = append(sectionNames, f.Form.Section)
			lastSection = f.Form.Section
			sectionCount++
		}
		if f.Form.Section != "" {
			if f.Form.SectionShowTitle {
				for idx, s := range res {
					if s.Title == f.Form.Section {
						s.ShowTitle = true
						res[idx] = s
					}
				}
			}

			if f.Form.SectionAutoCol > 1 {
				for idx, s := range res {
					if s.Title == f.Form.Section {
						s.AutoCol = f.Form.SectionAutoCol
						res[idx] = s
					}
				}
			}

			if f.Form.SectionWidth != "" {
				for idx, s := range res {
					if s.Title == f.Form.Section {
						s.Width = f.Form.SectionWidth
						res[idx] = s
					}
				}
			}
		}
	}

	sizePerBlock := math.Ceil(float64(len(res)) / float64(mt.Form.SectionSize))

	currentGroup := FormSectionGroup{}
	for _, sect := range res {
		currentGroup.Sections = append(currentGroup.Sections, sect)
		if len(currentGroup.Sections) == int(sizePerBlock) {
			groups = append(groups, currentGroup)
			currentGroup = FormSectionGroup{}
		}
	}
	if len(currentGroup.Sections) > 0 {
		groups = append(groups, currentGroup)
		currentGroup = FormSectionGroup{}
	}

	sectionGroups[typeString] = groups
	return groups, nil
}
