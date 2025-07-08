package page_generator

import "github.com/BekkkEvrika/page_generator/inputs"

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
	GetList(claims MapClaims) error
	Filter(obj interface{}) error
}

type ICreate interface {
	Create(claims MapClaims) error
}

type IUpdate interface {
	Update(claims MapClaims) error
}

type IDelete interface {
	Delete(claims MapClaims) error
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
