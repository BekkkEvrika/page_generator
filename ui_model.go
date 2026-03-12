package page_generator

import (
	"fmt"
	"reflect"

	"github.com/BekkkEvrika/page_generator/inputs"
)

type UIModel struct {
	model        interface{}
	fieldSize    int
	fieldTypes   []*FieldType
	container    IModel
	create       ICreate
	update       IUpdate
	delete       IDelete
	def          IDefault
	combo        IComboBox
	formActions  IFormActions
	validation   IFormValidation
	visibility   IFormVisibility
	fieldActions IFieldActions
	fileConfig   IFileConfig
	meta         IMetaData
	createUrl    string
	updateUrl    string
	filterUrl    string
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
	if val, ok := model.model.(IFormActions); ok {
		model.formActions = val
	}

	if val, ok := model.model.(IMetaData); ok {
		model.meta = val
	}
	if val, ok := model.model.(IFormValidation); ok {
		model.validation = val
	}
	if val, ok := model.model.(IFormVisibility); ok {
		model.visibility = val
	}
	if val, ok := model.model.(IFieldActions); ok {
		model.fieldActions = val
	}
	if val, ok := model.model.(IFileConfig); ok {
		model.fileConfig = val
	}
	return nil
}

// enrichInput заполняет поля inp данными из интерфейсов (combo, default, validation и т.д.)
func (model *UIModel) enrichInput(inp *inputs.Input, params *QueryParams, md map[string]interface{}) {
	if model.combo != nil {
		if items, ok := model.combo.GetComboItems(params, md)[inp.Name]; ok {
			inp.Options = items
		}
	}
	if model.def != nil {
		inp.DefaultValue = model.def.GetDefault(params, md)[inp.Name]
	}
	if model.validation != nil {
		if valid, ok := model.validation.GetFormValidation()[inp.Name]; ok {
			inp.Validation = &valid
		}
	}
	if model.visibility != nil {
		if valid, ok := model.visibility.GetFormVisibility()[inp.Name]; ok {
			inp.VisibilityRules = valid
		}
	}
	if model.fieldActions != nil {
		if valid, ok := model.fieldActions.GetFieldActions()[inp.Name]; ok {
			inp.FieldActions = valid
		}
	}
	if model.fileConfig != nil {
		if valid, ok := model.fileConfig.GetFileConfig()[inp.Name]; ok {
			inp.FileConfig = &valid
		}
	}
	if model.meta != nil {
		if meta, ok := model.meta.GetMetaData()[inp.Name]; ok {
			inp.MetaKey = meta.MetaKey
			inp.MetaData = meta.MetaData
		}
	}
}

// buildPage строит страницу с формой, фильтруя поля через accept(ft)
func (model *UIModel) buildPage(params *QueryParams, md map[string]interface{}, accept func(ft *FieldType) bool, transform func(ft *FieldType, inp *inputs.Input)) *Page {
	p := Page{}
	p.Form = &inputs.Form{}
	p.SetSettings(model.container.GetPageSettings())
	p.Form.Containers = model.container.GetContainers()
	p.FormActions = model.container.GetActions()
	for ind := 0; ind < model.fieldSize; ind++ {
		ft := model.fieldTypes[ind]
		if !accept(ft) {
			continue
		}
		inp, err := ft.makeInput()
		if err != nil || inp == nil {
			continue
		}
		inp.Id = inp.Name
		inp.Label = ft.PgText
		inp.Placeholder = ft.getPlaceholder()
		inp.ActionID = ft.pgAction
		model.enrichInput(inp, params, md)
		if transform != nil {
			transform(ft, inp)
		}
		if ft.pgContainer != "" {
			if container := inputs.GetContainerByKeyInSlice(*p.Form.Containers, ft.pgContainer); container != nil {
				container.Fields = append(container.Fields, *inp)
			}
		}
	}
	if model.formActions != nil {
		p.FormActions = model.formActions.GetFormActions()
	}
	return &p
}

func (model *UIModel) getUpdatePage(params *QueryParams, md map[string]interface{}) *Page {
	return model.buildPage(params, md, func(ft *FieldType) bool {
		return ft.pgEdit && ft.pg != "-"
	}, func(ft *FieldType, inp *inputs.Input) {
		if ft.getGormPrimaryKey() {
			inp.ReadOnly = true
		}
	})
}

func (model *UIModel) getFilterPage(params *QueryParams, md map[string]interface{}) *Page {
	return model.buildPage(params, md, func(ft *FieldType) bool {
		return ft.pg != "-"
	}, nil)
}

func (model *UIModel) getCreatePage(params *QueryParams, md map[string]interface{}) *Page {
	return model.buildPage(params, md, func(ft *FieldType) bool {
		return !ft.getGormAutoInc() && ft.pg != "-"
	}, nil)
}

func (model *UIModel) getFieldsModel(obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
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
		ft.pgSearchSource = field.Tag.Get(pgSearchSource)
		ft.pgSearchObject = field.Tag.Get(pgSearchObject)
		ft.pgFromName = field.Tag.Get(pgFromName)
		ft.pgVariant = field.Tag.Get(pgVariant)
		ft.pgAction = field.Tag.Get(pgAction)
		if err := ft.setPgType(field.Tag.Get(pgType)); err != nil {
			return err
		}
		ft.setPgText(field.Tag.Get(pgText))
		ft.setPgPlaceholder(field.Tag.Get(pgPlaceholder))
		if err := ft.setPgReadOnly(field.Tag.Get(pgReadOnly)); err != nil {
			return err
		}
		if err := ft.setPgEdit(field.Tag.Get(pgEdit)); err != nil {
			return err
		}
		ft.pgContainer = field.Tag.Get(pgContainer)
		model.fieldTypes = append(model.fieldTypes, &ft)
	}
	return nil
}
