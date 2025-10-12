package page_generator

import "github.com/BekkkEvrika/page_generator/inputs"

type MetaData struct {
	MetaKey  string
	MetaData string
}

type IQueryParams interface {
	GetDefaultQueryParams() map[string]string
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
	GetList(params *QueryParams) error
	Filter(obj interface{}, params *QueryParams) error
}

type IPagination interface {
	GetCount(params *QueryParams) (int, error)
}

type ICreate interface {
	Create(params *QueryParams) error
}

type IUpdate interface {
	Update(params *QueryParams) error
}

type IDelete interface {
	Delete(params *QueryParams) error
}

type IDefault interface {
	GetDefault(params *QueryParams) map[string]string
}

type IComboBox interface {
	GetComboItems(params *QueryParams) map[string]inputs.ComboItems
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

type IEditData interface {
	GetEditPage() inputs.LoadAction
}
