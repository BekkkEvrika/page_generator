package page_generator

import (
	"fmt"
	"github.com/BekkkEvrika/page_generator/inputs"
	"net/url"
	"reflect"
	"strings"
)

const (
	pg             = "pg"
	pgType         = "pgType"
	pgText         = "pgText"
	pgReadOnly     = "pgReadOnly"
	pgValid        = "pgValid"
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
	DeleteAction     = "delete"
	LoadTabAction    = "load"
	LoadDialogAction = "loadDialog"
	LoadHTML         = "loadHtml"
)

type PageModel struct {
	model            *UIModel
	listModel        interface{}
	tableModel       interface{}
	filterModel      *UIModel
	headerFieldTypes []*HeaderField
	getList          IGetList
	context          IContext
	indexes          IIndexes
	exports          IExports
	queryParams      IQueryParams
	pagination       IPagination
	pageListUrl      string
	defaultUrl       string
	addUrl           string
	editUrl          string
	deleteUrl        string
	countUrl         string
	listType         reflect.Type
	modelType        reflect.Type
	filterType       reflect.Type
}

func (pm *PageModel) getOnlyTable() *Page {
	page := &Page{}
	if pm.filterModel != nil {
		page = pm.filterModel.getFilterPage()
	}
	if pm.getList != nil {
		page.DataTable = &inputs.ExpDataTable{}
		page.DataTable.DefaultUrl = "/" + serviceName + pm.defaultUrl
		if pm.queryParams != nil {
			query, _ := pm.setQueryParams(page.DataTable.DefaultUrl)
			page.DataTable.DefaultUrl = query
		}
		if pm.pagination != nil {
			page.DataTable.PageSize = pageSize
			page.DataTable.ItemsCount = "/" + serviceName + pm.countUrl
		}
		for in, val := range pm.headerFieldTypes {
			h := inputs.TableHeader{
				Key:          val.getName(),
				Title:        val.Title,
				Order:        in + 1,
				Template:     val.Template,
				IsExportable: val.Export,
			}
			page.DataTable.Header = append(page.DataTable.Header, h)
		}
	}
	return page
}

func (pm *PageModel) getDataPage() *Page {
	page := &Page{}
	if pm.filterModel != nil {
		page = pm.filterModel.getFilterPage()
	}
	if pm.getList != nil {
		page.DataTable = &inputs.ExpDataTable{}
		page.DataTable.DefaultUrl = "/" + serviceName + pm.defaultUrl
		if pm.queryParams != nil {
			query, _ := pm.setQueryParams(page.DataTable.DefaultUrl)
			page.DataTable.DefaultUrl = query
		}
		if pm.pagination != nil {
			page.DataTable.PageSize = pageSize
			page.DataTable.ItemsCount = "/" + serviceName + pm.countUrl
		}
		for in, val := range pm.headerFieldTypes {
			h := inputs.TableHeader{
				Key:          val.getName(),
				Title:        val.Title,
				Order:        in + 1,
				Template:     val.Template,
				IsExportable: val.Export,
			}
			page.DataTable.Header = append(page.DataTable.Header, h)
		}
	}
	if pm.model.create != nil {
		page.DataTable.Add = inputs.LoadAction{
			Source: "/" + serviceName + pm.addUrl,
			Action: loadDialog,
			Text:   "Добавить",
		}
	}
	if pm.model.update != nil {
		page.DataTable.Edit = inputs.LoadAction{
			Source: "/" + serviceName + pm.editUrl,
			Action: loadDialog,
			Text:   "Изменить",
		}
	}
	if pm.model.delete != nil {
		page.DataTable.Delete = inputs.Action{
			Type:   DeleteAction,
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
	return page
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
	if val, ok := obj.(IPagination); ok {
		pm.pagination = val
	}
	pm.listModel = obj
	pm.listType = reflect.TypeOf(obj)
	return nil
}

func (pm *PageModel) SetFilterModel(obj interface{}, columns int) error {
	pm.filterModel = &UIModel{}
	if err := pm.filterModel.setModel(obj, columns); err != nil {
		return err
	}
	pm.filterType = reflect.TypeOf(obj)
	return nil
}

func (pm *PageModel) SetTableModel(obj interface{}) error {
	if err := pm.getFieldsTable(obj); err != nil {
		return err
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
	if val, ok := obj.(IQueryParams); ok {
		pm.queryParams = val
	}
	pm.tableModel = obj
	return nil
}

func (pm *PageModel) SetModel(obj interface{}, columns int) error {
	pm.model = &UIModel{}
	if err := pm.model.setModel(obj, columns); err != nil {
		return err
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

func (pm *PageModel) setQueryParams(defUrl string) (string, error) {
	u, err := url.Parse(defUrl)
	if err != nil {
		return defUrl, nil // или можно вернуть "", если критично
	}
	original := u.RawQuery

	// Собираем новые параметры вручную
	var parts []string
	for k, v := range pm.queryParams.GetDefaultQueryParams() {
		parts = append(parts, k+"="+v)
	}

	// Объединяем со старой query-строкой, если она была
	if original != "" {
		u.RawQuery = original + "&" + strings.Join(parts, "&")
	} else {
		u.RawQuery = strings.Join(parts, "&")
	}
	return u.String(), nil
}

func checkType(t reflect.Type) int {
	timeType := reflect.TypeOf(Date{})
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
	case reflect.Bool:
		return boolean
	default:
		return -1
	}
}
