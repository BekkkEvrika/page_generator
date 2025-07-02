package page_generator

import "fmt"

type HeaderField struct {
	Name     string
	JsonName string
	Gorm     string
	Template string
	Title    string
	Export   bool
}

func (f *HeaderField) init() {
	f.Template = ""
}

func (f *HeaderField) setExport(exp string) error {
	if exp == fTrue {
		f.Export = true
	} else if exp == fFalse {
		f.Export = false
	} else if exp == "" {
		f.Export = false
	} else {
		return fmt.Errorf("incorrect value")
	}
	return nil
}
