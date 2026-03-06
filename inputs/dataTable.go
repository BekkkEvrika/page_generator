package inputs

type DataTable struct {
	Title         string        `json:"title"`
	Header        []TableHeader `json:"header"`
	KeyColumn     string        `json:"keyColumn"`
	PageSize      int           `json:"pageSize"`
	ItemsCount    string        `json:"itemsCount"`
	Delete        Action        `json:"delete"`
	Edit          LoadAction    `json:"edit"`
	Add           LoadAction    `json:"add"`
	Context       []Action      `json:"context"`
	DefaultUrl    string        `json:"default_url"`
	Indexes       []Index       `json:"indexes"`
	Type          string        `json:"type"` //with-action, no-action ==""
	Exports       Export        `json:"exports"`
	Top           string        `json:"top"`    // html <h4>Title</h4> image src=base64();
	Bottom        string        `json:"bottom"` // html
	ActionKeys    ActionKeys    `json:"actionKeys"`
	HelperButtons []HelpButton  `json:"helperButtons"`
}

type Export struct {
	Word  bool `json:"word"`
	Excel bool `json:"excel"`
	PDF   bool `json:"pdf"`
}
