# page_generator — Документация

Библиотека для автоматической генерации структуры страниц (форм и таблиц) для фронтенда на основе Go-структур с тегами. Интегрируется с [Gin](https://github.com/gin-gonic/gin).

---

## Содержание

1. [Теги полей модели](#теги-полей-модели)
2. [Типы полей pgType](#типы-полей-pgtype)
3. [Интерфейсы](#интерфейсы)
   - [Обязательный IModel](#обязательный-imodel)
   - [Для форм](#для-форм)
   - [Для таблиц и списков](#для-таблиц-и-списков)
4. [Структуры данных](#структуры-данных)
5. [Инициализация](#инициализация)
6. [Регистрация моделей](#регистрация-моделей)
7. [Маршруты](#маршруты)

---

## Теги полей модели

| Тег              | Описание                                                                    | Значения                     |
|------------------|-----------------------------------------------------------------------------|------------------------------|
| `pg`             | Исключить поле из UI полностью                                              | `"-"`                        |
| `pgType`         | Тип элемента формы                                                          | см. [Типы полей](#типы-полей-pgtype) |
| `pgText`         | Подпись поля (`label`)                                                      | любая строка                 |
| `pgPlaceholder`  | Placeholder поля. Если не задан — используется значение `pgText`            | любая строка                 |
| `pgEdit`         | Включить поле в форму редактирования                                        | `"true"` / `"false"`         |
| `pgReadOnly`     | Поле только для чтения                                                      | `"true"` / `"false"`         |
| `pgFromName`     | json-имя поля-источника значения (для связанных полей)                      | json-имя поля                |
| `pgSearch`       | URL источника данных для поля типа `search`                                 | URL строка                   |
| `pgSName`        | Имя поля в объекте результата поиска                                        | строка                       |
| `pgContainer`    | Ключ контейнера (`Container.Key`), в который помещается поле                | строка — ключ контейнера     |
| `pgVariant`      | Вариант отображения кнопки (`button`)                                       | `"primary"` / `"secondary"` / `"destructive"` / `"outline"` / `"ghost"` / `"link"` |
| `pgAction`       | ID действия формы (`FormAction.ID`), которое вызывается при клике на кнопку | строка — ID действия         |

> Тег `pg:"-"` полностью исключает поле из всех страниц UI. Это единственное допустимое значение тега `pg`.

Теги для **табличной модели** (`SetTableModel`):

| Тег        | Описание                        |
|------------|---------------------------------|
| `dtTitle`  | Заголовок колонки таблицы       |
| `dtExport` | Включить в экспорт (`"true"`)   |
| `dtTemp`   | Шаблон отображения ячейки       |

---

## Типы полей `pgType`

| Значение   | Описание                       | Автоматически назначается для  |
|------------|--------------------------------|-------------------------------|
| `select`   | Выпадающий список              | —                             |
| `date`     | Поле даты                      | `DateTime` с gorm `type:date` |
| `datetime` | Поле даты и времени            | `DateTime`                    |
| `text`     | Текстовое поле (однострочное)  | `string` с gorm `size` ≤ 60   |
| `number`   | Числовое поле                  | числовые типы Go              |
| `checkbox` | Флажок                         | `bool`                        |
| `label`    | Метка (только отображение)     | —                             |
| `search`   | Поле поиска со справочником    | —                             |
| `textarea` | Текстовое поле (многострочное) | `string` с gorm `size` > 60   |
| `hidden`   | Скрытое поле                   | —                             |
| `file`     | Загрузка файла                 | —                             |
| `button`   | Кнопка                         | —                             |

---

## Интерфейсы

### Обязательный IModel

Каждая модель формы **обязана** реализовывать `IModel`:

```go
type IModel interface {
    GetContainers()   *[]inputs.Container
    GetPageSettings() *PageSettings
    GetActions()      *[]inputs.FormAction
}
```

| Метод              | Описание                                                       |
|--------------------|----------------------------------------------------------------|
| `GetContainers()`  | Возвращает структуру контейнеров формы                         |
| `GetPageSettings()`| Возвращает настройки страницы (FormId, Version, Title и др.)   |
| `GetActions()`     | Возвращает список действий формы. Если не нужно — вернуть `nil`|

Пример:
```go
func (m *MyModel) GetContainers() *[]inputs.Container {
    return &[]inputs.Container{
        {Key: "main",   Direction: "vertical",   Title: "Основное"},
        {Key: "extra",  Direction: "horizontal",  Title: "Дополнительно",
            Containers: []inputs.Container{
                {Key: "nested", Direction: "vertical"},
            },
        },
    }
}

func (m *MyModel) GetPageSettings() *PageSettings {
    return &PageSettings{
        FormId:  "my-form",
        Version: "1.0",
        Title:   "Моя форма",
        Card:    true,
    }
}

func (m *MyModel) GetActions() *[]inputs.FormAction {
    return &[]inputs.FormAction{
        {
            ID:      "save-action",
            Trigger: inputs.Click,
            Config:  &inputs.FormActionConfig{Type: inputs.APICall, URL: "/api/save", Method: "POST"},
        },
    }
}
```

> `Key` контейнера должен совпадать со значением тега `pgContainer` у полей модели.  
> Контейнеры поддерживают вложенность через поле `Containers []Container`.

---

### Для форм

Все интерфейсы **опциональны**. Если модель реализует интерфейс — функциональность автоматически включается.

> В методах, возвращающих `map`, **ключ** — это значение тега `json` соответствующего поля модели.

---

#### `ICreate`
Активирует страницу создания и маршруты `POST /create/*`.
```go
type ICreate interface {
    Create(params *QueryParams) error
}
```

#### `IUpdate`
Активирует страницу редактирования и маршрут `PUT /update/data`.
```go
type IUpdate interface {
    Update(params *QueryParams) error
}
```

#### `IDelete`
Активирует маршрут `DELETE /delete/data`.
```go
type IDelete interface {
    Delete(params *QueryParams) error
}
```

#### `IDefault`
Значения по умолчанию для полей формы.
```go
type IDefault interface {
    GetDefault(params *QueryParams, mp map[string]interface{}) map[string]string
}
```
Пример:
```go
func (m *MyModel) GetDefault(params *QueryParams, mp map[string]interface{}) map[string]string {
    return map[string]string{"status": "active"}
}
```

#### `IComboBox`
Элементы для полей типа `select`.
```go
type IComboBox interface {
    GetComboItems(params *QueryParams, mp map[string]interface{}) map[string]inputs.ComboItems
}
```
Пример:
```go
func (m *MyModel) GetComboItems(params *QueryParams, mp map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{
        "status": {{ID: "active", Text: "Активен"}, {ID: "inactive", Text: "Неактивен"}},
    }
}
```

#### `IFormValidation`
Правила валидации полей.
```go
type IFormValidation interface {
    GetFormValidation() map[string]inputs.FieldValidation
}
```
Пример:
```go
func (m *MyModel) GetFormValidation() map[string]inputs.FieldValidation {
    min, max := 3, 100
    return map[string]inputs.FieldValidation{
        "name": {MinLength: &min, MaxLength: &max, Message: "От 3 до 100 символов"},
    }
}
```

#### `IFormVisibility`
Правила видимости полей.
```go
type IFormVisibility interface {
    GetFormVisibility() map[string][]inputs.Rule
}
```
Пример:
```go
func (m *MyModel) GetFormVisibility() map[string][]inputs.Rule {
    return map[string][]inputs.Rule{
        "passport": {{Field: "type", Operator: inputs.OpEq, Value: "person"}},
    }
}
```

#### `IFieldActions`
Действия над полями при выполнении условий.
```go
type IFieldActions interface {
    GetFieldActions() map[string][]inputs.FieldAction
}
```
Пример:
```go
func (m *MyModel) GetFieldActions() map[string][]inputs.FieldAction {
    return map[string][]inputs.FieldAction{
        "city": {{
            When:         inputs.Rule{Field: "country", Operator: inputs.OpEq, Value: "UZ"},
            Action:       "clear",
            TargetFields: []string{"district"},
        }},
    }
}
```
Допустимые значения `Action`: `"clear"` | `"setRequired"` | `"setOptional"` | `"show"` | `"hide"` | `"setValue"`

#### `IFormActions`
Действия формы (вызываются по событию — `init`, `change`, `click`).
```go
type IFormActions interface {
    GetFormActions() *[]inputs.FormAction
}
```
Пример:
```go
func (m *MyModel) GetFormActions() *[]inputs.FormAction {
    return &[]inputs.FormAction{
        {
            ID:      "load-data",
            Trigger: inputs.Init,
            Config:  &inputs.FormActionConfig{Type: inputs.APICall, URL: "/api/data", Method: "GET"},
        },
    }
}
```

> `IFormActions` имеет приоритет над `GetActions()` из `IModel` — если модель реализовывает `IFormActions`, используется его результат.

#### `IFileConfig`
Конфигурация загрузки файлов для полей типа `file`.
```go
type IFileConfig interface {
    GetFileConfig() map[string]inputs.FileConfig
}
```
Пример:
```go
func (m *MyModel) GetFileConfig() map[string]inputs.FileConfig {
    return map[string]inputs.FileConfig{
        "avatar": {Accept: "image/*", MaxSizeBytes: 5 * 1024 * 1024, MaxFiles: 1, UploadURL: "/api/upload"},
    }
}
```

#### `IMetaData`
Мета-данные для полей типа `search`.
```go
type IMetaData interface {
    GetMetaData() map[string]MetaData
}
type MetaData struct {
    MetaKey  string
    MetaData string
}
```

#### `IEditData`
Переопределяет поведение кнопки "Изменить" в таблице.
```go
type IEditData interface {
    GetEditPage() inputs.LoadAction
}
```

---

### Для таблиц и списков

#### `IGetList`
Обязателен для списковой модели. Загружает данные.
```go
type IGetList interface {
    GetList(params *QueryParams) error
}
```

#### `IPagination`
Добавляет пагинацию.
```go
type IPagination interface {
    GetCount(params *QueryParams) (int, error)
}
```

#### `IFilter`
Фильтрация данных таблицы.
```go
type IFilter interface {
    Filter(obj interface{}, params *QueryParams) error
}
```

#### `IContext`
Контекстные действия строк таблицы.
```go
type IContext interface {
    GetContextActions() []inputs.Action
}
```

#### `IIndexes`
Итоговые строки таблицы (count, sum, avg).
```go
type IIndexes interface {
    GetIndexes() []inputs.Index
}
```

#### `IExports`
Кнопки экспорта (Excel, Word, PDF).
```go
type IExports interface {
    GetExports() inputs.Export
}
```

#### `IQueryParams`
Дефолтные query-параметры URL источника данных.
```go
type IQueryParams interface {
    GetDefaultQueryParams() map[string]string
}
```

---

## Структуры данных

### `Page`
```go
type Page struct {
    FormId      string               `json:"formId"`
    Version     string               `json:"version"`
    Title       string               `json:"title"`
    Description string               `json:"description"`
    Card        bool                 `json:"card,omitempty"`
    Form        *inputs.Form         `json:"form"`
    FormActions *[]inputs.FormAction `json:"formActions,omitempty"`
    DataTable   *inputs.DataTable    `json:"dataTable"`
}
```

### `PageSettings`
```go
type PageSettings struct {
    FormId      string `json:"formId"`
    Version     string `json:"version"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Card        bool   `json:"card,omitempty"`
}
```

### `Form`
```go
type Form struct {
    Containers *[]Container `json:"containers"`
}
```

### `Container`
```go
type Container struct {
    Key            string      `json:"id"`
    Direction      string      `json:"direction"`            // "horizontal" | "vertical"
    Gap            int         `json:"gap"`
    Card           bool        `json:"card,omitempty"`
    Align          string      `json:"align"`                // "start"|"center"|"end"|"between"|"stretch"
    GridColumns    int         `json:"gridColumns,omitempty"`
    Title          string      `json:"title"`
    Fields         []Input     `json:"fields,omitempty"`
    Containers     []Container `json:"containers,omitempty"` // вложенные контейнеры
    VisibilityRule *Rule       `json:"visibilityRule,omitempty"`
}
```

### `Input`
```go
type Input struct {
    Id              string           `json:"id"`
    Type            string           `json:"type"`
    Name            string           `json:"name,omitempty"`
    Label           string           `json:"label,omitempty"`
    ActionID        string           `json:"actionId,omitempty"`
    Variant         string           `json:"variant,omitempty"`    // для button: "primary"|"secondary"|"destructive"|"outline"|"ghost"|"link"
    FromName        string           `json:"fromName,omitempty"`
    ReadOnly        bool             `json:"readOnly,omitempty"`
    Placeholder     string           `json:"placeholder,omitempty"`
    Validation      *FieldValidation `json:"validation,omitempty"`
    Options         ComboItems       `json:"options,omitempty"`
    VisibilityRules []Rule           `json:"visibilityRules,omitempty"`
    FieldActions    []FieldAction    `json:"fieldActions,omitempty"`
    FileConfig      *FileConfig      `json:"fileConfig,omitempty"`
    DefaultValue    string           `json:"defaultValue,omitempty"`
    Format          string           `json:"format,omitempty"`
    ColSpan         int              `json:"colSpan,omitempty"`
    Hint            string           `json:"hint,omitempty"`
    DataType        string           `json:"dataType,omitempty"`
    Search          string           `json:"searchSource,omitempty"`
    SearchName      string           `json:"searchObject,omitempty"`
    MetaData        string           `json:"metaData,omitempty"`
    MetaKey         string           `json:"metaKey,omitempty"`
}
```

### `FormAction`
```go
type FormAction struct {
    ID      string            `json:"id"`
    Trigger FormActionTrigger `json:"trigger"` // "init" | "change" | "click"
    Config  *FormActionConfig `json:"config,omitempty"`
}

type FormActionConfig struct {
    Type           FormActionType `json:"type"`    // "apiCall" | "changeApiCall" | "calculate"
    URL            string         `json:"url,omitempty"`
    Method         string         `json:"method,omitempty"` // GET, POST, PUT, DELETE
    Formula        string         `json:"formula,omitempty"`
    SuccessMessage string         `json:"successMessage,omitempty"`
}
```

### `Rule`
```go
type Rule struct {
    Field    string       `json:"field"`
    Operator RuleOperator `json:"operator"` // "eq"|"neq"|"gt"|"gte"|"lt"|"lte"|"in"|"notIn"|"empty"|"notEmpty"|"contains"
    Value    interface{}  `json:"value,omitempty"`
    ValueRef string       `json:"valueRef,omitempty"`
    Combine  string       `json:"combine,omitempty"` // "and" | "or"
}
```

### `FieldValidation`
```go
type FieldValidation struct {
    Min            *int   `json:"min,omitempty"`
    Max            *int   `json:"max,omitempty"`
    MinLength      *int   `json:"minLength,omitempty"`
    MaxLength      *int   `json:"maxLength,omitempty"`
    Pattern        string `json:"pattern,omitempty"`
    PatternMessage string `json:"patternMessage,omitempty"`
    Message        string `json:"message,omitempty"`
}
```

### `FileConfig`
```go
type FileConfig struct {
    Accept          string   `json:"accept,omitempty"`
    MaxSizeBytes    int64    `json:"maxSizeBytes,omitempty"`
    MaxFiles        int      `json:"maxFiles,omitempty"`
    UploadURL       string   `json:"uploadUrl,omitempty"`
    AcceptMimeTypes []string `json:"acceptMimeTypes,omitempty"`
}
```

### `ComboItem`
```go
type ComboItem struct {
    ID   interface{} `json:"value"`
    Text interface{} `json:"label"`
}
```

---

## Инициализация

```go
err := page_generator.SetDefinitions(func() error {
    // регистрация моделей
    return nil
}, page_generator.PageSetting{
    Service:    "my-service",  // префикс всех маршрутов
    DateFormat: "DD.MM.YYYY",
    TimeFormat: "HH:mm",
    PageSize:   20,
})
```

---

## Регистрация моделей

```go
pm := &page_generator.PageModel{}

// Обязательно: модель формы — должна реализовывать IModel
pm.SetModel(&MyModel{})

// Опционально: табличная модель (определяет заголовки колонок)
pm.SetTableModel(&MyTableModel{})

// Опционально: списковая модель (должна реализовывать IGetList)
pm.SetListModel(&MyListModel{})

// Опционально: модель фильтра (должна реализовывать IModel)
pm.SetFilterModel(&MyFilterModel{})

page_generator.AddPageModel("users", pm)
```

---

## Маршруты

```go
r := gin.Default()
page_generator.GetModelsRoutes(r)
r.Run(":8080")
```

Автоматически регистрируемые маршруты для ключа `"users"`:

| Метод    | Маршрут               | Условие          | Описание                             |
|----------|-----------------------|------------------|--------------------------------------|
| `GET`    | `/users/list/page`    | `SetListModel`   | Структура страницы списка            |
| `GET`    | `/users/list/table`   | `SetListModel`   | Только структура таблицы             |
| `GET`    | `/users/list/data`    | `SetListModel`   | Данные списка                        |
| `GET`    | `/users/list/count`   | `IPagination`    | Количество записей                   |
| `POST`   | `/users/list/filter`  | `SetFilterModel` | Данные списка с фильтром             |
| `GET`    | `/users/create/page`  | `ICreate`        | Структура формы создания             |
| `POST`   | `/users/create/page`  | `ICreate`        | Структура формы создания (с данными) |
| `POST`   | `/users/create/data`  | `ICreate`        | Сохранить новую запись               |
| `GET`    | `/users/update/page`  | `IUpdate`        | Структура формы редактирования       |
| `POST`   | `/users/update/page`  | `IUpdate`        | Структура формы редактирования       |
| `PUT`    | `/users/update/data`  | `IUpdate`        | Сохранить изменения                  |
| `DELETE` | `/users/delete/data`  | `IDelete`        | Удалить запись                       |
