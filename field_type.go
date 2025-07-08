package page_generator

import (
	"fmt"
	"github.com/BekkkEvrika/page_generator/inputs"
	"strconv"
	"strings"
)

const (
	number    = 0
	text      = 1
	slice     = 2
	date      = 3
	structure = 4
	fTrue     = "true"
	fFalse    = "false"
)

type FieldType struct {
	Name           string
	Type           int
	JsonName       string
	Gorm           string
	PgType         string
	PgText         string
	PgReadOnly     bool
	PgValid        string
	PgFormat       string
	pgMax          int
	pgMin          int
	pgEdit         bool
	pgVisible      string
	pgTemplate     string
	pg             string
	pgFileSource   string
	pgFileExt      string
	pgFileMaxSize  int
	pgFromName     string
	pgSearchSource string
	pgSearchObject string
	pgDataType     string
}

func (f *FieldType) init() {
	f.PgText = ""
	f.PgReadOnly = false
	f.PgValid = ""
	f.PgFormat = ""
}

func (f *FieldType) makeInput() (*inputs.Input, error) {
	fn, ok := inputCreators[f.PgType]
	if !ok {
		return nil, fmt.Errorf("input type incorrect")
	}
	return fn(f), nil
}

func (f *FieldType) getFromName() string {
	if f.pgFromName != "" {
		return f.pgFromName
	}
	return f.getName()
}

func (f *FieldType) getValidation() (bool, string) {
	if f.PgValid != "" {
		return true, f.PgValid
	}
	return false, ""
}

func (f *FieldType) getName() string {
	if f.JsonName != "" {
		return strings.Split(f.JsonName, ",")[0]
	} else {
		return f.Name
	}
}

func (f *FieldType) getGormAutoInc() bool {
	parts := strings.Split(f.Gorm, ";")
	for _, part := range parts {
		if part == "autoIncrement" {
			return true
		}
	}
	return false
}

func (f *FieldType) getGormPrimaryKey() bool {
	parts := strings.Split(f.Gorm, ";")
	for _, part := range parts {
		if part == "primaryKey" {
			return true
		}
	}
	return false
}

func (f *FieldType) getGormSize() int {
	parts := strings.Split(f.Gorm, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "size:") {
			size, _ := strconv.Atoi(strings.TrimPrefix(part, "size:"))
			return size
		}
	}
	return 0
}

func (f *FieldType) getGormType() string {
	parts := strings.Split(f.Gorm, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "type:") {
			tp := strings.TrimPrefix(part, "type:")
			return tp
		}
	}
	return ""
}

func (f *FieldType) setMaxLength(text string) {
	num, _ := strconv.Atoi(text)
	f.pgMax = num
}

func (f *FieldType) setMinLength(text string) {
	num, _ := strconv.Atoi(text)
	f.pgMin = num
}

func (f *FieldType) setPgFormat(text string) {
	if text != "" {
		f.PgFormat = text
	} else {
		f.PgFormat = "YYYY-MM-DD"
	}
}

func (f *FieldType) setPg(text string) {
	f.pg = text
}

func (f *FieldType) setPgVisible(text string) {
	f.pgVisible = text
}

func (f *FieldType) setPgValid(text string) {
	f.PgValid = text
}

func (f *FieldType) setPgReadOnly(read string) error {
	if read == fTrue {
		f.PgReadOnly = true
	} else if read == fFalse {
		f.PgReadOnly = false
	} else if read == "" {
		f.PgReadOnly = false
	} else {
		return fmt.Errorf("incorrect value")
	}
	return nil
}

func (f *FieldType) setPgEdit(edit string) error {
	if edit == fTrue {
		f.pgEdit = true
	} else if edit == fFalse {
		f.pgEdit = false
	} else if edit == "" {
		f.pgEdit = false
	} else {
		return fmt.Errorf("incorrect value")
	}
	return nil
}

func (f *FieldType) setPgText(text string) {
	f.PgText = text
}

func (f *FieldType) setPgType(tp string) error {
	if tp != "" {
		if contains(types, tp) {
			f.PgType = tp
		} else {
			return fmt.Errorf(" error input type ")
		}
	} else {
		switch f.Type {
		case number:
			f.PgType = types[3]
			f.pgDataType = "number"
		case text:
			f.PgType = types[2]
			f.pgDataType = "string"
		case date:
			if f.getGormType() == "date" {
				f.PgType = types[1]
			} else {
				f.PgType = types[2]
			}
		case -1:
			return fmt.Errorf(" field type not found")
		}
	}
	return nil
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
