package page_generator

import "github.com/BekkkEvrika/page_generator/inputs"

type Page struct {
	Form      *inputs.FormExported `json:"form"`
	Menu      *inputs.Menu         `json:"menu"`
	IsMenu    bool                 `json:"isMenu"`
	DataTable *inputs.ExpDataTable `json:"dataTable"`
}

func (p *Page) Init() {
	p.Form = &inputs.FormExported{}
	p.Menu = &inputs.Menu{}
	p.DataTable = &inputs.ExpDataTable{}
}
