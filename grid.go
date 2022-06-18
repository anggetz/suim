package suim

type GridField struct {
	Field      string
	Label      string
	Halign     string
	Valign     string
	LabelField string
	Length     int
	Width      string
	Pos        int
}

type GridSetting struct {
	IDField string
}

type GridConfig struct {
	Setting GridSetting
	Fields  []GridField
}
