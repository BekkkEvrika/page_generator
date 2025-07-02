package inputs

type ActionKeys struct {
	Key1 string            `json:"key1"`
	Key2 string            `json:"key2"`
	Keys map[string]string `json:"keys"`
	//PrintKeys     map[string]string `json:"printKeys"`
	SuccessSource string   `json:"successSource"`
	CancelSource  string   `json:"cancelSource"`
	Access        []string `json:"access,omitempty"`
}

func (a ActionKeys) GetName() string {
	return "Success/Cancel"
}

func (a ActionKeys) GetAccess() []string {
	return a.Access
}
