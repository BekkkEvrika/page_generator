package page_generator

import (
	"fmt"
	"github.com/BekkkEvrika/page_generator/inputs"
	"math"
	"reflect"
	"strconv"
)

type UIModel struct {
	model          interface{}
	columnSize     int
	fieldSize      int
	fieldTypes     []*FieldType
	create         ICreate
	update         IUpdate
	delete         IDelete
	def            IDefault
	combo          IComboBox
	completeNodes  ICompleteNodes
	fileExtensions IFileExtensions
	meta           IMetaData
	clearNodes     IClearNodes
	createUrl      string
	updateUrl      string
	filterUrl      string
}

func (model *UIModel) setModel(obj interface{}, columns int) error {
	if err := model.getFieldsModel(obj); err != nil {
		return err
	}
	model.model = obj
	model.columnSize = columns
	if val, ok := model.model.(ICreate); ok {
		model.create = val
	}
	if val, ok := model.model.(IUpdate); ok {
		model.update = val
	}
	if val, ok := model.model.(IDelete); ok {
		model.delete = val
	}
	if val, ok := model.model.(IDefault); ok {
		model.def = val
	}
	if val, ok := model.model.(IComboBox); ok {
		model.combo = val
	}
	if val, ok := model.model.(ICompleteNodes); ok {
		model.completeNodes = val
	}
	if val, ok := model.model.(IFileExtensions); ok {
		model.fileExtensions = val
	}
	if val, ok := model.model.(IMetaData); ok {
		model.meta = val
	}
	if val, ok := model.model.(IClearNodes); ok {
		model.clearNodes = val
	}
	return nil
}

func (model *UIModel) getUpdatePage() *Page {
	p := Page{}
	p.Form = &inputs.FormExported{}
	colLen := int(math.Ceil(float64(model.fieldSize / model.columnSize)))
	indCol := 0
	column := inputs.Column{}
	for ind := 0; ind < model.fieldSize; ind++ {
		indCol++
		ft := model.fieldTypes[ind]
		if ft.pgEdit && ft.pg != "-" {
			inp, err := ft.makeInput()
			if err == nil {
				if ft.getGormPrimaryKey() {
					inp.ReadOnly = true
				}
				if model.combo != nil {
					if items, ok := model.combo.GetComboItems()[inp.Name]; ok {
						inp.Items = items
					}
				}
				if model.def != nil {
					inp.DefaultValue = model.def.GetDefault()[inp.Name]
				}
				if model.completeNodes != nil {
					if items, ok := model.completeNodes.GetCompleteNodes()[inp.Name]; ok {
						inp.CompleteNodes = items
					}
				}
				if model.fileExtensions != nil {
					if items, ok := model.fileExtensions.GetFileExtensions()[inp.Name]; ok {
						inp.FileExtensions = items
					}
				}
				if model.fileExtensions != nil {
					if items, ok := model.fileExtensions.GetFileExtensions()[inp.Name]; ok {
						inp.FileExtensions = items
					}
				}
				if model.meta != nil {
					if meta, ok := model.meta.GetMetaData()[inp.Name]; ok {
						inp.MetaKey = meta.MetaKey
						inp.MetaData = meta.MetaData
					}
				}
				if model.clearNodes != nil {
					if items, ok := model.clearNodes.GetClearNodes()[inp.Name]; ok {
						inp.ClearNodes = items
					}
				}
				column.Inputs = append(column.Inputs, *inp)
			}
		}
		if indCol == colLen {
			p.Form.Columns = append(p.Form.Columns, column)
			column = inputs.Column{}
			indCol = 0
		}
	}
	p.Form.Submit.Text = "Сохранить"
	p.Form.Submit.Source = "/" + serviceName + model.updateUrl
	p.Form.Submit.Method = "PUT"
	p.Form.Submit.SuccessMessage = "Успешно сохранено!"
	p.Form.Submit.ConfirmMessage = "Действительно хотите совершить эту операцию?"
	p.Form.Submit.LastAction = "success-message,close"
	return &p
}

func (model *UIModel) getFilterPage() *Page {
	p := Page{}
	p.Form = &inputs.FormExported{}
	colLen := int(math.Ceil(float64(model.fieldSize) / float64(model.columnSize)))
	indCol := 0
	column := inputs.Column{}
	for ind := 0; ind < model.fieldSize; ind++ {
		ft := model.fieldTypes[ind]
		if ft.pg != "-" {
			indCol++
			inp, err := ft.makeInput()
			if err == nil && inp != nil {
				if model.def != nil {
					inp.DefaultValue = model.def.GetDefault()[inp.Name]
				}
				if model.combo != nil {
					if items, ok := model.combo.GetComboItems()[inp.Name]; ok {
						inp.Items = items
					}
				}
				if model.completeNodes != nil {
					if items, ok := model.completeNodes.GetCompleteNodes()[inp.Name]; ok {
						inp.CompleteNodes = items
					}
				}
				if model.fileExtensions != nil {
					if items, ok := model.fileExtensions.GetFileExtensions()[inp.Name]; ok {
						inp.FileExtensions = items
					}
				}
				if model.meta != nil {
					if meta, ok := model.meta.GetMetaData()[inp.Name]; ok {
						inp.MetaKey = meta.MetaKey
						inp.MetaData = meta.MetaData
					}
				}
				if model.clearNodes != nil {
					if items, ok := model.clearNodes.GetClearNodes()[inp.Name]; ok {
						inp.ClearNodes = items
					}
				}
				column.Inputs = append(column.Inputs, *inp)
			}
		}
		if indCol == colLen {
			p.Form.Columns = append(p.Form.Columns, column)
			column = inputs.Column{}
			indCol = 0
		}
	}
	if len(column.Inputs) > 0 {
		p.Form.Columns = append(p.Form.Columns, column)
	}
	p.Form.Submit.Text = "Найти"
	p.Form.Submit.Source = "/" + serviceName + model.filterUrl
	p.Form.Submit.Method = "POST"
	p.Form.Submit.Type = "loader"
	return &p
}

func (model *UIModel) getCreatePage() *Page {
	p := Page{}
	p.Form = &inputs.FormExported{}
	colLen := int(math.Ceil(float64(model.fieldSize / model.columnSize)))
	indCol := 0
	column := inputs.Column{}
	for ind := 0; ind < model.fieldSize; ind++ {
		ft := model.fieldTypes[ind]
		if !ft.getGormAutoInc() && ft.pg != "-" {
			indCol++
			inp, err := ft.makeInput()
			if err == nil && inp != nil {
				if model.def != nil {
					inp.DefaultValue = model.def.GetDefault()[inp.Name]
				}
				if model.combo != nil {
					if items, ok := model.combo.GetComboItems()[inp.Name]; ok {
						inp.Items = items
					}
				}
				if model.completeNodes != nil {
					if items, ok := model.completeNodes.GetCompleteNodes()[inp.Name]; ok {
						inp.CompleteNodes = items
					}
				}
				if model.fileExtensions != nil {
					if items, ok := model.fileExtensions.GetFileExtensions()[inp.Name]; ok {
						inp.FileExtensions = items
					}
				}
				if model.meta != nil {
					if meta, ok := model.meta.GetMetaData()[inp.Name]; ok {
						inp.MetaKey = meta.MetaKey
						inp.MetaData = meta.MetaData
					}
				}
				if model.clearNodes != nil {
					if items, ok := model.clearNodes.GetClearNodes()[inp.Name]; ok {
						inp.ClearNodes = items
					}
				}
				column.Inputs = append(column.Inputs, *inp)
			}
		}
		if indCol == colLen {
			p.Form.Columns = append(p.Form.Columns, column)
			column = inputs.Column{}
			indCol = 0
		}
	}
	if len(column.Inputs) > 0 {
		p.Form.Columns = append(p.Form.Columns, column)
	}
	p.Form.Submit.Text = "Сохранить"
	p.Form.Submit.Source = "/" + serviceName + model.createUrl
	p.Form.Submit.Method = "POST"
	p.Form.Submit.SuccessMessage = "Успешно сохранено!"
	p.Form.Submit.ConfirmMessage = "Действительно хотите совершить эту операцию?"
	p.Form.Submit.LastAction = "success-message,close"
	return &p
}

func (model *UIModel) getFieldsModel(obj interface{}) error {
	val := reflect.ValueOf(obj)
	// Если указатель — разыменуем
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// Проверка: это struct?
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Не структура ")
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		ft := FieldType{Type: checkType(field.Type)}
		ft.init()
		ft.Name = field.Name
		ft.JsonName = field.Tag.Get("json")
		ft.Gorm = field.Tag.Get("gorm")
		ft.setPg(field.Tag.Get(pg))
		if ft.pg == "-" {
			continue
		}
		model.fieldSize++
		ft.pgTemplate = field.Tag.Get(pgTemplate)
		ft.pgSearchSource = field.Tag.Get(pgSearchSource)
		ft.pgSearchObject = field.Tag.Get(pgSearchObject)
		ft.pgFromName = field.Tag.Get(pgFromName)
		ft.pgFileSource = field.Tag.Get(pgFileSource)
		ft.pgFileMaxSize, _ = strconv.Atoi(field.Tag.Get(pgFileMaxSize))
		if err := ft.setPgType(field.Tag.Get(pgType)); err != nil {
			return err
		}
		ft.setPgText(field.Tag.Get(pgText))
		if err := ft.setPgReadOnly(field.Tag.Get(pgReadOnly)); err != nil {
			return err
		}
		if err := ft.setPgEdit(field.Tag.Get(pgEdit)); err != nil {
			return err
		}
		ft.setPgValid(field.Tag.Get(pgValid))
		ft.setMaxLength(field.Tag.Get(pgMaxLength))
		ft.setMinLength(field.Tag.Get(pgMinLength))
		ft.setPgVisible(field.Tag.Get(pgVisible))
		model.fieldTypes = append(model.fieldTypes, &ft)
	}
	return nil
}
