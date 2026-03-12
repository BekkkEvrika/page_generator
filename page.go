package page_generator

import "github.com/BekkkEvrika/page_generator/inputs"

type Page struct {
	FormId      string               `json:"formId"`
	Version     string               `json:"version"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Form        *inputs.Form         `json:"form"`
	Card        bool                 `json:"card,omitempty"`
	FormActions *[]inputs.FormAction `json:"formActions,omitempty"`
	DataTable   *inputs.DataTable    `json:"dataTable"`
}

func (p *Page) SetSettings(settings *PageSettings) {
	if settings != nil {
		p.FormId = settings.FormId
		p.Version = settings.Version
		p.Title = settings.Title
		p.Description = settings.Description
		p.Card = settings.Card
	}
}

type PageSettings struct {
	FormId      string `json:"formId"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Card        bool   `json:"card,omitempty"`
}

func (p *Page) Init() {
	p.Form = &inputs.Form{}
	p.DataTable = &inputs.DataTable{}
}
