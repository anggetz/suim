package suim

type FormListItem struct {
	Key  string
	Text string
}

type FormField struct {
	Field       string
	Label       string
	Hint        string
	Placeholder string
	Control     string
	Disable     bool
	Required    bool
	MinLength   int
	MaxLength   int

	UseList bool
	Items   []FormListItem

	UseLookup bool
	LookupURL string
	LookupKey string
	LookupTxt string

	ShowTitle  bool
	ShowHint   bool
	ShowDetail bool
	FixTitle   bool
	FixDetail  bool

	Section string
	Row     int
	Col     int
}

type FormSection struct {
	Title     string
	Name      string
	ShowTitle bool
	Rows      [][]FormField
	AutoCol   int
}

type FormSetting struct {
	Title            string
	ShowTitle        bool
	InitialMode      string
	ShowButtons      bool
	ShowEditButton   bool
	ShowSubmitButton bool
	ShowCancelButton bool
	SubmitText       string
	AutoCol          int
}

type FormConfig struct {
	Setting  FormSetting
	Sections []FormSection
}
