package inputs

type Input struct {
	Type           string     `json:"type"` //combo-box,date-time,text-view,number-view,check-box,label,search-view,text-field, hidden, auto-complete,file-uploader,button
	Name           string     `json:"name,omitempty"`
	FromName       string     `json:"fromName,omitempty"`
	ReadOnly       bool       `json:"readOnly,omitempty"`
	Text           string     `json:"text,omitempty"`
	MaxLength      int        `json:"maxLength,omitempty"`
	MinLength      int        `json:"minLength,omitempty"`
	IsDefault      bool       `json:"isDefault,omitempty"`
	MetaData       string     `json:"metaData,omitempty"` //search-view uchun malumot boradi qimat
	MetaKey        string     `json:"metaKey,omitempty"`  //search-view uchun malumot boradi kalit
	ValidMessage   string     `json:"validMessage,omitempty"`
	Format         string     `json:"format,omitempty"`
	Items          ComboItems `json:"items,omitempty"`
	ClearNodes     []string   `json:"clearNodes,omitempty"`
	CompleteNodes  []string   `json:"completeNodes,omitempty"`
	InfoSource     string     `json:"infoSource,omitempty"` // <id,text>
	ItemsSource    string     `json:"itemsSource,omitempty"`
	SearchName     string     `json:"searchObject,omitempty"`
	DefaultValue   string     `json:"defaultValue,omitempty"`
	Search         string     `json:"searchSource,omitempty"`
	DataType       string     `json:"dataType,omitempty"` //number,string,bool : default string
	Visible        string     `json:"visible,omitempty"`
	Template       string     `json:"template,omitempty"` // {} ба мегирим параметроя
	FileSource     string     `json:"fileSource,omitempty"`
	FileExtensions []string   `json:"fileExtensions,omitempty"`
	FileMaxSize    int        `json:"fileMaxSize,omitempty"` // byte кати
}

type ComboItem struct {
	ID   interface{} `json:"id"`
	Text interface{} `json:"text"`
}

type ComboItems []ComboItem

func (i Input) GetName() string {
	return i.Text
}
