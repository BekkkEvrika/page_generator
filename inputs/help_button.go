package inputs

type HelpButton struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Access []string `json:"access,omitempty"`
}

func (h HelpButton) GetName() string {
	return h.Name
}

func (h HelpButton) GetAccess() []string {
	return h.Access
}
