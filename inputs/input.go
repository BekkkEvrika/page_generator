package inputs

type Input struct {
	Type           string     `json:"type"` //combo-box,date-time,text-view,number-view,check-box,label,search-view,text-field, hidden, auto-complete,file-uploader
	Name           string     `json:"name"`
	FromName       string     `json:"fromName"`
	ReadOnly       bool       `json:"readOnly"`
	Text           string     `json:"text"`
	MaxLength      int        `json:"maxLength"`
	MinLength      int        `json:"minLength"`
	IsDefault      bool       `json:"isDefault"`
	MetaData       string     `json:"metaData"` //search-view uchun malumot boradi qimat
	MetaKey        string     `json:"metaKey"`  //search-view uchun malumot boradi kalit
	ValidMessage   string     `json:"validMessage"`
	Format         string     `json:"format"`
	Items          ComboItems `json:"items"`
	ClearNodes     []string   `json:"clearNodes"`
	CompleteNodes  []string   `json:"completeNodes"`
	InfoSource     string     `json:"infoSource"` // <id,text>
	ItemsSource    string     `json:"itemsSource,omitempty"`
	SearchName     string     `json:"searchObject"`
	DefaultValue   string     `json:"defaultValue"`
	Search         string     `json:"searchSource"`
	DataType       string     `json:"dataType"` //number,string,bool : default string
	Visible        string     `json:"visible"`
	Template       string     `json:"template"` // {} ба мегирим параметроя
	FileSource     string     `json:"fileSource"`
	FileExtensions []string   `json:"fileExtensions"`
	FileMaxSize    int        `json:"fileMaxSize"` // byte кати
	Access         []string   `json:"access,omitempty"`
}

type ComboItem struct {
	ID   interface{} `json:"id"`
	Text interface{} `json:"text"`
}

type ComboItems []ComboItem

func (i Input) GetName() string {
	return i.Text
}

func (i Input) GetAccess() []string {
	return i.Access
}
