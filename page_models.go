package page_generator

import (
	"fmt"
	"github.com/BekkkEvrika/page_generator/inputs"
	"math"
	"reflect"
	"strconv"
	"time"
)

const (
	pg             = "pg"
	pgType         = "pgType"
	pgText         = "pgText"
	pgReadOnly     = "pgReadOnly"
	pgValid        = "pgValid"
	pgFormat       = "pgFormat"
	pgMaxLength    = "pgMax"
	pgMinLength    = "pgMin"
	pgEdit         = "pgEdit"
	pgVisible      = "pgVisible"
	pgTemplate     = "pgTemp"
	pgFileSource   = "pgFileSource"
	pgFileMaxSize  = "pgFileMaxSize"
	pgFromName     = "pgFromName"
	pgSearchSource = "pgSearch"
	pgSearchObject = "pgSName"
)

const (
	dtTemp   = "dtTemp"
	dtTitle  = "dtTitle"
	dtExport = "dtExport"
)

const (
	loadDialog = "dialog"
	loadTab    = "tab"
)

const (
	deleteAction     = "delete"
	loadTabAction    = "load"
	loadDialogAction = "loadDialog"
	loadHTML         = "loadHtml"
)

type PageModel struct {
	model            interface{}
	listModel        interface{}
	tableModel       interface{}
	filterModel      interface{}
	modelColSize     int
	modelFieldTypes  []*FieldType
	headerFieldTypes []*HeaderField

	getList        IGetList
	create         ICreate
	update         IUpdate
	delete         IDelete
	def            IDefault
	combo          IComboBox
	completeNodes  ICompleteNodes
	fileExtensions IFileExtensions
	meta           IMetaData
	clearNodes     IClearNodes
	context        IContext
	indexes        IIndexes
	exports        IExports

	pageListUrl string
	filterUrl   string
	defaultUrl  string
	addUrl      string
	editUrl     string
	deleteUrl   string
	createUrl   string
	updateUrl   string

	listType   reflect.Type
	modelType  reflect.Type
	filterType reflect.Type
}

func (pm *PageModel) getDataPage() *Page {
	page := Page{}
	if pm.filterModel != nil {
		page.Form = &inputs.FormExported{}
	}
	if pm.getList != nil {
		page.DataTable = &inputs.ExpDataTable{}
		page.DataTable.DefaultUrl = "/" + serviceName + pm.defaultUrl
		for in, val := range pm.headerFieldTypes {
			h := inputs.TableHeader{
				Key:          val.Name,
				Title:        val.Title,
				Order:        in + 1,
				Template:     val.Template,
				IsExportable: val.Export,
			}
			page.DataTable.Header = append(page.DataTable.Header, h)
		}
	}
	if pm.create != nil {
		page.DataTable.Add = inputs.LoadAction{
			Source: "/" + serviceName + pm.addUrl,
			Action: loadDialog,
			Text:   "Добавить",
		}
	}
	if pm.update != nil {
		page.DataTable.Edit = inputs.LoadAction{
			Source: "/" + serviceName + pm.editUrl,
			Action: loadDialog,
			Text:   "Изменить",
		}
	}
	if pm.delete != nil {
		page.DataTable.Delete = inputs.Action{
			Type:   deleteAction,
			Source: "/" + serviceName + pm.deleteUrl,
			Method: "DELETE",
			Text:   "Удалить",
		}
	}
	if pm.context != nil {
		page.DataTable.Context = pm.context.GetContextActions()
	}
	if pm.indexes != nil {
		page.DataTable.Indexes = pm.indexes.GetIndexes()
	}
	if pm.exports != nil {
		page.DataTable.Exports = pm.exports.GetExports()
	}
	return &page
}

func (pm *PageModel) getCreatePage() *Page {
	p := Page{}
	p.Form = &inputs.FormExported{}
	colLen := int(math.Ceil(float64(len(pm.modelFieldTypes) / pm.modelColSize)))
	indCol := 0
	column := inputs.Column{}
	for ind := 0; ind < len(pm.modelFieldTypes); ind++ {
		indCol++
		ft := pm.modelFieldTypes[ind]
		if !ft.getGormAutoInc() && ft.pg != "-" {
			inp, err := ft.makeInput()
			if err == nil && inp != nil {
				if pm.def != nil {
					inp.DefaultValue = pm.def.GetDefault()[inp.Name]
				}
				if pm.combo != nil {
					if items, ok := pm.combo.GetComboItems()[inp.Name]; ok {
						inp.Items = items
					}
				}
				if pm.completeNodes != nil {
					if items, ok := pm.completeNodes.GetCompleteNodes()[inp.Name]; ok {
						inp.CompleteNodes = items
					}
				}
				if pm.fileExtensions != nil {
					if items, ok := pm.fileExtensions.GetFileExtensions()[inp.Name]; ok {
						inp.FileExtensions = items
					}
				}
				if pm.meta != nil {
					if meta, ok := pm.meta.GetMetaData()[inp.Name]; ok {
						inp.MetaKey = meta.MetaKey
						inp.MetaData = meta.MetaData
					}
				}
				if pm.clearNodes != nil {
					if items, ok := pm.clearNodes.GetClearNodes()[inp.Name]; ok {
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
	p.Form.Submit.Source = "/" + serviceName + pm.createUrl
	p.Form.Submit.Method = "POST"
	p.Form.Submit.SuccessMessage = "Успешно сохранено!"
	p.Form.Submit.ConfirmMessage = "Действительно хотите совершить эту операцию?"
	p.Form.Submit.LastAction = "success-message,close"
	return &p
}

func (pm *PageModel) getUpdatePage() *Page {
	p := Page{}
	p.Form = &inputs.FormExported{}
	colLen := int(math.Ceil(float64(len(pm.modelFieldTypes) / pm.modelColSize)))
	indCol := 0
	column := inputs.Column{}
	for ind := 0; ind < len(pm.modelFieldTypes); ind++ {
		indCol++
		ft := pm.modelFieldTypes[ind]
		if ft.pgEdit && ft.pg != "-" {
			inp, err := ft.makeInput()
			if err == nil {
				if ft.getGormPrimaryKey() {
					inp.ReadOnly = true
				}
				if pm.combo != nil {
					if items, ok := pm.combo.GetComboItems()[inp.Name]; ok {
						inp.Items = items
					}
				}
				if pm.def != nil {
					inp.DefaultValue = pm.def.GetDefault()[inp.Name]
				}
				if pm.completeNodes != nil {
					if items, ok := pm.completeNodes.GetCompleteNodes()[inp.Name]; ok {
						inp.CompleteNodes = items
					}
				}
				if pm.fileExtensions != nil {
					if items, ok := pm.fileExtensions.GetFileExtensions()[inp.Name]; ok {
						inp.FileExtensions = items
					}
				}
				if pm.fileExtensions != nil {
					if items, ok := pm.fileExtensions.GetFileExtensions()[inp.Name]; ok {
						inp.FileExtensions = items
					}
				}
				if pm.meta != nil {
					if meta, ok := pm.meta.GetMetaData()[inp.Name]; ok {
						inp.MetaKey = meta.MetaKey
						inp.MetaData = meta.MetaData
					}
				}
				if pm.clearNodes != nil {
					if items, ok := pm.clearNodes.GetClearNodes()[inp.Name]; ok {
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
	p.Form.Submit.Source = "/" + serviceName + pm.createUrl
	p.Form.Submit.Method = "PUT"
	p.Form.Submit.SuccessMessage = "Успешно сохранено!"
	p.Form.Submit.ConfirmMessage = "Действительно хотите совершить эту операцию?"
	p.Form.Submit.LastAction = "success-message,close"
	return &p
}

func (pm *PageModel) SetListModel(obj interface{}) error {
	if pm.tableModel == nil {
		return fmt.Errorf("first set table model")
	}
	if val, ok := obj.(IGetList); ok {
		pm.getList = val
	} else {
		return fmt.Errorf("not list model")
	}
	if val, ok := obj.(IContext); ok {
		pm.context = val
	}
	if val, ok := obj.(IIndexes); ok {
		pm.indexes = val
	}
	if val, ok := obj.(IExports); ok {
		pm.exports = val
	}
	pm.listModel = obj
	pm.listType = reflect.TypeOf(obj)
	return nil
}

func (pm *PageModel) SetFilterModel(obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Не структура ")
	}
	pm.filterModel = obj
	pm.filterType = reflect.TypeOf(obj)
	return nil
}

func (pm *PageModel) SetTableModel(obj interface{}) error {
	if err := pm.getFieldsTable(obj); err != nil {
		return err
	}
	pm.tableModel = obj
	return nil
}

func (pm *PageModel) SetModel(obj interface{}, columns int) error {
	if err := pm.getFieldsModel(obj); err != nil {
		return err
	}
	pm.model = obj
	pm.modelColSize = columns
	if val, ok := pm.model.(ICreate); ok {
		pm.create = val
	}
	if val, ok := pm.model.(IUpdate); ok {
		pm.update = val
	}
	if val, ok := pm.model.(IDelete); ok {
		pm.delete = val
	}
	if val, ok := pm.model.(IDefault); ok {
		pm.def = val
	}
	if val, ok := pm.model.(IComboBox); ok {
		pm.combo = val
	}
	if val, ok := pm.model.(ICompleteNodes); ok {
		pm.completeNodes = val
	}
	if val, ok := pm.model.(IFileExtensions); ok {
		pm.fileExtensions = val
	}
	if val, ok := pm.model.(IMetaData); ok {
		pm.meta = val
	}
	if val, ok := pm.model.(IClearNodes); ok {
		pm.clearNodes = val
	}
	pm.modelType = reflect.TypeOf(obj)
	return nil
}

func (pm *PageModel) getFieldsTable(obj interface{}) error {
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
		hf := HeaderField{}
		hf.init()
		hf.Name = field.Name
		hf.JsonName = field.Tag.Get("json")
		hf.Template = field.Tag.Get(dtTemp)
		hf.Title = field.Tag.Get(dtTitle)
		if err := hf.setExport(field.Tag.Get(dtExport)); err != nil {
			return err
		}
		pm.headerFieldTypes = append(pm.headerFieldTypes, &hf)
	}
	return nil
}

func (pm *PageModel) getFieldsModel(obj interface{}) error {
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
		ft.setPgFormat(field.Tag.Get(pgFormat))
		ft.setMaxLength(field.Tag.Get(pgMaxLength))
		ft.setMinLength(field.Tag.Get(pgMinLength))
		ft.setPg(field.Tag.Get(pg))
		ft.setPgVisible(field.Tag.Get(pgVisible))
		pm.modelFieldTypes = append(pm.modelFieldTypes, &ft)
	}
	return nil
}

func checkType(t reflect.Type) int {
	timeType := reflect.TypeOf(time.Time{})
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return number
	case reflect.String:
		return text
	case reflect.Slice:
		return slice
	case reflect.Struct:
		if t == timeType {
			return date
		}
		return structure
	default:
		return -1
	}
}
