package inputs

type TableHeader struct {
	Key          string       `json:"key"`
	Title        string       `json:"title"`
	IsExportable bool         `json:"isExportable"` //true- pechat meshad false- pechat nameshad
	Type         string       `json:"type"`         //file,image,number,static,auto-complete,''
	Element      TableElement `json:"element"`
	TableSubmit  TableSubmit  `json:"tableSubmit"`
	Order        int          `json:"order"`
	Template     string       `json:"template"`
	Access       []string     `json:"access,omitempty"`
}

func (t TableHeader) GetName() string {
	return t.Title
}

func (t TableHeader) GetAccess() []string {
	return t.Access
}

type TableElement struct {
	Items []ComboItem `json:"items"`
	Where string      `json:"where"` // <key>=,<>,>,< <key>
}

type TableSubmit struct {
	Method     string `json:"method"`
	Source     string `json:"source"`
	Text       string `json:"text"`
	LastAction string `json:"lastAction"` //submitForm,successMessage
	Message    string `json:"message"`
}
