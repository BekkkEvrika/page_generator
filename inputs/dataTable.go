package inputs

type ExpDataTable struct {
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

func (et *ExpDataTable) urlCorrect(kid, acKid string) {
	if kid != acKid {
		if et.DefaultUrl != "" {
			et.DefaultUrl = "/reverse/" + kid + et.DefaultUrl
		}
		if et.ItemsCount != "" {
			et.ItemsCount = "/reverse/" + kid + et.ItemsCount
		}
		if et.Delete.Source != "" {
			et.Delete.Source = "/reverse/" + kid + et.Delete.Source
		}
		if et.Edit.Source != "" {
			et.Edit.Source = "/reverse/" + kid + et.Edit.Source
		}
		if et.Add.Source != "" {
			et.Add.Source = "/reverse/" + kid + et.Add.Source
		}
		if et.ActionKeys.SuccessSource != "" {
			et.ActionKeys.SuccessSource = "/reverse/" + kid + et.ActionKeys.SuccessSource
			et.ActionKeys.CancelSource = "/reverse/" + kid + et.ActionKeys.CancelSource
		}
		for i := 0; i < len(et.Header); i++ {
			if et.Header[i].TableSubmit.Source != "" {
				et.Header[i].TableSubmit.Source = "/reverse/" + kid + et.Header[i].TableSubmit.Source
			}
		}
		for i := 0; i < len(et.Context); i++ {
			if et.Context[i].Source != "" {
				et.Context[i].Source = "/reverse/" + kid + et.Context[i].Source
			}
		}

	}
}
