package suim

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/eaciit/toolkit"
)

func ObjToFields(obj interface{}) (*ObjMeta, []Field, error) {
	fields := []Field{}
	meta := new(ObjMeta)

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, []Field{}, errors.New("object should be a struct or pointer of a struct")
	}
	t := v.Type()
	fieldNum := v.NumField()

	gs := GridSetting{}

	fs := FormSetting{
		ShowTitle:        false,
		ShowButtons:      true,
		ShowEditButton:   true,
		ShowSubmitButton: true,
		ShowCancelButton: true,
		InitialMode:      "edit",
		SubmitText:       "Save",
		AutoCol:          1,
	}

	for i := 0; i < fieldNum; i++ {
		ft := t.Field(i)
		tag := ft.Tag
		alias := tag.Get(toolkit.TagName())
		if alias == "-" {
			continue
		}

		field, e := toField(ft)
		if e != nil {
			return nil, []Field{}, fmt.Errorf("%s: %s", alias, e)
		}
		fields = append(fields, field)

		SetIfStruct(&fs, "Title", fs.Title == "", TagValue(tag, "obj_title", t.Name()))
		SetIfStruct(&fs, "ShowTitle", TagExist(tag, "form_show_title"), TagValue(tag, "form_show_title", "") == "1")
		SetIfStruct(&fs, "ShowButtons", TagExist(tag, "form_show_buttons"), TagValue(tag, "form_show_buttons", "") == "1")
		SetIfStruct(&fs, "ShowEditButtons", TagExist(tag, "form_show_edit_buttons"), TagValue(tag, "form_show_edit_buttons", "") == "1")
		SetIfStruct(&fs, "ShowSubmitButtons", TagExist(tag, "form_show_submi_buttons"), TagValue(tag, "form_show_submit_buttons", "") == "1")
		SetIfStruct(&fs, "ShowCancelButtons", TagExist(tag, "form_show_cancel_buttons"), TagValue(tag, "form_show_cancel_buttons", "") == "1")
		SetIfStruct(&fs, "InitialMode", TagExist(tag, "form_initial_mode"), TagValue(tag, "form_initial_mode", "edit"))
		SetIfStruct(&fs, "SubmitText", TagExist(tag, "form_submit_text"), TagValue(tag, "form_submit_text", "Save"))
		SetIfStruct(&fs, "AutoCol", TagExist(tag, "form_auto_col"), DefInt(TagValue(tag, "form_auto_col", "1"), 1))
	}

	meta.Grid = gs
	meta.Form = fs

	return meta, fields, nil
}

func toField(rt reflect.StructField) (Field, error) {
	f := Field{}
	f.Field = rt.Name
	f.DataType = rt.Type.Name()

	tag := rt.Tag

	f.GridElement = TagValue(tag, "grid", "show")
	f.FormElement = TagValue(tag, "form", "show")

	if f.FormElement == "show" {
		form := FormField{}
		form.Field = TagValue(tag, toolkit.TagName(), rt.Name)
		pos := strings.Split(TagValue(tag, "form_pos", ","), ",")
		rowStr := DefTxt(pos[0], "0")
		colStr := "0"
		if len(pos) > 1 {
			colStr = DefTxt(pos[1], "0")
		}
		form.Row, _ = strconv.Atoi(rowStr)
		form.Col, _ = strconv.Atoi(colStr)
		form.Section = TagValue(tag, "form_section", "General")
		form.Control = TagValue(tag, "form_control", "")
		if form.Control == "" {
			switch f.DataType {
			case "int":
				form.Control = "number"
			case "time.Time", "*time.Time":
				form.Control = "date"
			case "bool":
				form.Control = "checkbox"
			default:
				form.Control = "text"
			}
		}
		form.Disable = TagExist(tag, "form_disable")
		form.FixDetail = TagExist(tag, "form_fix_detail")
		form.FixTitle = TagExist(tag, "form_fix_title")
		form.Hint = TagValue(tag, "form_hint", "")
		items := strings.Split(TagValue(tag, "form_items", ""), "|")
		form.Items = []FormListItem{}
		for _, item := range items {
			parts := strings.Split(item, ":")
			if parts[0] == "" {
				continue
			}
			if len(parts) > 1 {
				form.Items = append(form.Items, FormListItem{Key: parts[0], Text: parts[1]})
			} else if len(parts) == 1 {
				form.Items = append(form.Items, FormListItem{Key: parts[0], Text: parts[0]})
			}
		}
		form.Label = TagValue(tag, "form_label", TagValue(tag, "label", Label(rt.Name, "")))
		form.UseList = len(form.Items) > 0 || TagExist(tag, "form_use_list")
		if form.UseList && len(items) == 0 {
			lookups := strings.Split(TagValue(tag, "form_lookup", ""), "|")
			if len(lookups) < 2 {
				return f, errors.New("lookup should contains at least 2 elements: url and fieldof key")
			}
			form.LookupKey = lookups[1]
			form.LookupTxt = lookups[1]
			if len(lookups) > 2 {
				form.LookupTxt = lookups[2]
			}
		}
		form.Placeholder = TagValue(tag, "form_placeholder", form.Label)
		lengths := strings.Split(TagValue(tag, "form_length", "0,999"), ",")
		form.MinLength = DefInt(DefSliceItem(lengths, 0, "0"), 0)
		form.MaxLength = DefInt(DefSliceItem(lengths, 1, "999"), 999)
		form.Required = TagExist(tag, "form_required")
		form.ShowDetail = TagExist(tag, "form_show_detail")
		form.ShowHint = TagExist(tag, "form_show_hint")
		form.ShowTitle = TagExist(tag, "form_show_title")
		f.Form = form
	}

	if f.GridElement != "hide" {
		grid := GridField{}
		grid.Field = f.Form.Field
		grid.Halign = TagValue(tag, "grid_halign", "start")
		grid.Valign = TagValue(tag, "grid_valign", "start")
		grid.Label = TagValue(tag, "grid_label", TagValue(tag, "label", Label(rt.Name, "")))
		grid.LabelField = TagValue(tag, "grid_label_field", "")
		grid.Length = DefInt(TagValue(tag, "grid_length", "0"), 0)
		grid.Pos = DefInt(TagValue(tag, "grid_pos", "0"), 0)
		grid.Width = TagValue(tag, "width", "")
		f.Grid = grid
	}

	return f, nil
}
