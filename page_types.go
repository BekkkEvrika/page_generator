package page_generator

import "testPager/page_generator/inputs"

type MetaData struct {
	MetaKey  string
	MetaData string
}

type IContext interface {
	GetContextActions() []inputs.Action
}

type IIndexes interface {
	GetIndexes() []inputs.Index
}

type IExports interface {
	GetExports() inputs.Export
}

type IGetList interface {
	GetList() error
	Filter(obj interface{}) error
}

type ICreate interface {
	Create() error
}

type IUpdate interface {
	Update() error
}

type IDelete interface {
	Delete() error
}

type IDefault interface {
	GetDefault() map[string]string
}

type IComboBox interface {
	GetComboItems() map[string]inputs.ComboItems
}

type ICompleteNodes interface {
	GetCompleteNodes() map[string][]string
}

type IFileExtensions interface {
	GetFileExtensions() map[string][]string
}

type IMetaData interface {
	GetMetaData() map[string]MetaData
}

type IClearNodes interface {
	GetClearNodes() map[string][]string
}
