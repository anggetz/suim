package suim

type GridField struct {
	Field      string `json:"field"`
	Label      string `json:"label"`
	Halign     string `json:"halign"`
	Valign     string `json:"valign"`
	LabelField string `json:"labelField"`
	Length     int    `json:"length"`
	Width      string `json:"width"`
	Pos        int    `json:"pos"`
}

type GridSetting struct {
	IDField        string   `json:"idField"`
	KeywordFields  []string `json:"keywordFields"`
	SortableFields []string `json:"sortable"`
}

type GridConfig struct {
	Setting GridSetting `json:"setting"`
	Fields  []GridField `json:"fields"`
}
