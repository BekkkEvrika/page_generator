package page_generator

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/BekkkEvrika/page_generator/inputs"
)

type UIModel struct {
	model          interface{}
	fieldSize      int
	fieldTypes     []*FieldType
	container      IModel
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

func (model *UIModel) setModel(obj interface{}) error {
	if err := model.getFieldsModel(obj); err != nil {
		return err
	}
	model.model = obj
	if val, ok := model.model.(IModel); !ok {
		return fmt.Errorf("model must implement IModel interface")
	} else {
		model.container = val
	}
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

func (model *UIModel) getUpdatePage(params *QueryParams, md map[string]interface{}) *Page {
	p := Page{}
	p.Form = &inputs.Form{}
	p.Form.Containers = model.container.GetContainers()
	for ind := 0; ind < model.fieldSize; ind++ {
		ft := model.fieldTypes[ind]
		if ft.pgEdit && ft.pg != "-" {
			inp, err := ft.makeInput()
			if err == nil {
				if ft.getGormPrimaryKey() {
					inp.ReadOnly = true
				}
				if model.combo != nil {
					if items, ok := model.combo.GetComboItems(params, md)[inp.Name]; ok {
						inp.Items = items
					}
				}
				if model.def != nil {
					inp.DefaultValue = model.def.GetDefault(params, md)[inp.Name]
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

				// Добавляем input в контейнер только если контейнер существует в форме (с учетом вложенности)
				if ft.pgContainer != "" {
					container := inputs.GetContainerByKeyInSlice(p.Form.Containers, ft.pgContainer)
					if container != nil {
						container.Inputs = append(container.Inputs, *inp)
					}
				}
			}
		}
	}
	return &p
}

func (model *UIModel) getFilterPage(params *QueryParams, md map[string]interface{}) *Page {
	p := Page{}
	p.Form = &inputs.Form{}
	p.Form.Containers = model.container.GetContainers()
	for ind := 0; ind < model.fieldSize; ind++ {
		ft := model.fieldTypes[ind]
		if ft.pg != "-" {
			inp, err := ft.makeInput()
			if err == nil && inp != nil {
				if model.def != nil {
					inp.DefaultValue = model.def.GetDefault(params, md)[inp.Name]
				}
				if model.combo != nil {
					if items, ok := model.combo.GetComboItems(params, md)[inp.Name]; ok {
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

				// Добавляем input в контейнер только если контейнер существует в форме (с учетом вложенности)
				if ft.pgContainer != "" {
					container := inputs.GetContainerByKeyInSlice(p.Form.Containers, ft.pgContainer)
					if container != nil {
						container.Inputs = append(container.Inputs, *inp)
					}
				}
			}
		}
	}
	return &p
}

func (model *UIModel) getCreatePage(params *QueryParams, md map[string]interface{}) *Page {
	p := Page{}
	p.Form = &inputs.Form{}
	fmt.Println("ssdsdg")
	p.Form.Containers = model.container.GetContainers()
	for ind := 0; ind < model.fieldSize; ind++ {
		ft := model.fieldTypes[ind]
		if !ft.getGormAutoInc() && ft.pg != "-" {
			inp, err := ft.makeInput()
			if err == nil && inp != nil {
				if model.def != nil {
					inp.DefaultValue = model.def.GetDefault(params, md)[inp.Name]
				}
				if model.combo != nil {
					if items, ok := model.combo.GetComboItems(params, md)[inp.Name]; ok {
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

				// Добавляем input в контейнер только если контейнер существует в форме (с учетом вложенности)
				if ft.pgContainer != "" {
					container := inputs.GetContainerByKeyInSlice(p.Form.Containers, ft.pgContainer)
					if container != nil {
						container.Inputs = append(container.Inputs, *inp)
					}
				}
			}
		}
	}
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
		ft.pgContainer = field.Tag.Get(pgContainer)
		model.fieldTypes = append(model.fieldTypes, &ft)
	}
	return nil
}
