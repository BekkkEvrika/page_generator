package inputs

type LoadAction struct {
	Source string   `json:"source"`
	Action string   `json:"action"` //tab,dialog
	Text   string   `json:"text"`
	Access []string `json:"access,omitempty"`
}

func (l LoadAction) GetName() string {
	return l.Text
}

func (l LoadAction) GetAccess() []string {
	return l.Access
}
