package suim

type FormListItem struct {
	Key  string
	Text string
}

type FormField struct {
	Field       string `json:"field"`
	Label       string `json:"label"`
	Hint        string `json:"hint"`
	Placeholder string `json:"placeHolder"`
	Kind        string `json:"kind"`
	Disable     bool   `json:"disable"`
	Required    bool   `json:"required"`
	MinLength   int    `json:"minLength"`
	MaxLength   int    `json:"maxLength"`
	ReadOnly    bool   `json:"readOnly"`

	UseList bool           `json:"useList"`
	Items   []FormListItem `json:"items"`

	UseLookup          bool     `json:"useLookup"`
	LookupURL          string   `json:"lookupURL"`
	LookupKey          string   `json:"lookupKey"`
	LookupLabelFields  []string `json:"lookupLabelFields"`
	LookupSearchFields []string `json:"lookupSearchFields"`

	ShowTitle  bool `json:"showTitle"`
	ShowHint   bool `json:"showHint"`
	ShowDetail bool `json:"showDetail"`
	FixTitle   bool `json:"fixTitle"`
	FixDetail  bool `json:"fixDetail"`

	Section string `json:"section"`
	Row     int    `json:"row"`
	Col     int    `json:"col"`

	LabelField string `json:"labelField"`
}

type FormSection struct {
	Title     string        `json:"title"`
	Name      string        `json:"name"`
	ShowTitle bool          `json:"showTitle"`
	Rows      [][]FormField `json:"rows"`
	AutoCol   int           `json:"autoCol"`
}

type FormSetting struct {
	IDField          string `json:"idField"`
	Title            string `json:"title"`
	ShowTitle        bool   `json:"showTitle"`
	InitialMode      string `json:"initialMode"`
	HideButtons      bool   `json:"hideButtons"`
	HideEditButton   bool   `json:"hideEditButton"`
	HideSubmitButton bool   `json:"hideSubmitButton"`
	HideCancelButton bool   `json:"hideCancelButton"`
	SubmitText       string `json:"submitText"`
	AutoCol          int    `json:"autoCol"`
}

type FormConfig struct {
	Setting  FormSetting   `json:"setting"`
	Sections []FormSection `json:"sections"`
}
