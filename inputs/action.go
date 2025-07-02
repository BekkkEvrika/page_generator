package inputs

type Action struct {
	Type   string `json:"type"` // delete, load,loadDialog,loadHtml
	Source string `json:"source"`
	Method string `json:"method"`
	Text   string `json:"text"`
}

func (a Action) GetName() string {
	return a.Text
}
