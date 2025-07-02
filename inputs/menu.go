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

func (m Menu) Generate(access []string, token string, kid, acKid string) *Menu {
	menu := Menu{}
	for _, vls := range m.MenuColumns {
		clm := MenuColumn{}
		for _, val := range vls.Items {
			if !existsAccess(access, val.Access) {
				continue
			}
			val.Access = nil
			if kid != acKid {
				if val.Source != "" {
					val.Source = "/reverse/" + kid + val.Source
				}
			}
			clm.Items = append(clm.Items, val)
		}
		menu.MenuColumns = append(menu.MenuColumns, clm)
	}
	return &menu
}
