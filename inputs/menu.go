package inputs

type Menu struct {
	Name        string       `json:"name"`
	MenuColumns []MenuColumn `json:"menuColumns"`
	Access      []string     `json:"access"`
}

func (m Menu) GetAccess() []string {
	return m.Access
}

func (m Menu) GetName() string {
	return m.Name
}

type MenuColumn struct {
	Items []Action `json:"items"`
}
