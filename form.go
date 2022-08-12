package suim

type FormListItem struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

type FormField struct {
	Field          string `json:"field"`
	Label          string `json:"label"`
	Hint           string `json:"hint"`
	Hide           bool   `json:"hide"`
	Placeholder    string `json:"placeHolder"`
	Kind           string `json:"kind"`
	Disable        bool   `json:"disable"`
	Required       bool   `json:"required"`
	Multiple       bool   `json:"multiple"`
	MultiRow       int    `json:"multiRow"`
	MinLength      int    `json:"minLength"`
	MaxLength      int    `json:"maxLength"`
	ReadOnly       bool   `json:"readOnly"`
	ReadOnlyOnEdit bool   `json:"readOnlyOnEdit"`
	ReadOnlyOnNew  bool   `json:"readOnlyOnNew"`

	UseList  bool           `json:"useList"`
	AllowAdd bool           `json:"allowAdd"`
	Items    []FormListItem `json:"items"`

	UseLookup     bool     `json:"useLookup"`
	LookupUrl     string   `json:"lookupUrl"`
	LookupKey     string   `json:"lookupKey"`
	LookupLabels  []string `json:"lookupLabels"`
	LookupSearchs []string `json:"lookupSearchs"`

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
