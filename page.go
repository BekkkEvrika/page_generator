package page_generator

import "github.com/BekkkEvrika/page_generator/inputs"

type Page struct {
	FormId      string            `json:"formId"`
	Version     string            `json:"version"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Form        *inputs.Form      `json:"form"`
	DataTable   *inputs.DataTable `json:"dataTable"`
}

func (p *Page) Init() {
	p.Form = &inputs.Form{}
	p.DataTable = &inputs.DataTable{}
}
