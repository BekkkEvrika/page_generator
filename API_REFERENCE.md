# 📚 API Reference - Page Generator

Полный справочник всех функций, методов и структур библиотеки.

---

## 📦 Основной пакет `page_generator`

### Функции инициализации

#### `SetDefinitions(initFunc InitFunction, setting PageSetting) error`

Инициализирует библиотеку с параметрами.

**Параметры:**
- `initFunc` - callback функция для регистрации моделей
- `setting` - настройки сервиса

**Возвращает:** ошибку если инициализация не удалась

**Пример:**
```go
err := pg.SetDefinitions(func() error {
    // регистрация моделей
    return nil
}, pg.PageSetting{
    Service: "my-service",
    DateFormat: "2006-01-02",
    PageSize: 20,
})
```

---

#### `GetModelsRoutes(engine *gin.Engine) error`

Регистрирует все маршруты для моделей в Gin.

**Параметры:**
- `engine` - экземпляр Gin Engine

**Возвращает:** ошибку если регистрация не удалась

**Пример:**
```go
err := pg.GetModelsRoutes(ginEngine)
```

---

#### `GetModelsRoutesGroup(routerGroup *gin.RouterGroup) error`

Регистрирует маршруты в группе маршрутов Gin.

**Параметры:**
- `routerGroup` - группа маршрутов Gin

**Возвращает:** ошибку если регистрация не удалась

**Пример:**
```go
api := router.Group("/api")
err := pg.GetModelsRoutesGroup(api)
```

---

### Структуры

#### `PageSetting`

Настройки для инициализации Page Generator.

```go
type PageSetting struct {
    Service          string                // имя сервиса для генерации URL
    DateFormat       string                // формат даты (例: "2006-01-02")
    PageSize         int                   // размер страницы по умолчанию
    KeyCloakSettings *KeyCloakSettings     // настройки Keycloak (опционально)
}
```

---

#### `KeyCloakSettings`

Настройки для интеграции с Keycloak.

```go
type KeyCloakSettings struct {
    BaseURL    string  // базовый URL Keycloak
    Realm      string  // область Keycloak
    ClientId   string  // ID клиента
    ClientUUID string  // UUID клиента
    Secret     string  // секрет клиента
}
```

---

#### `PageModel`

Основной класс для управления страницей данных.

```go
type PageModel struct {
    // приватные поля...
}
```

**Методы:**

##### `SetModel(obj interface{}, columns int) error`

Устанавливает модель для редактирования/создания.

```go
pageModel.SetModel(&User{}, 2) // 2 колонки в форме
```

---

##### `SetTableModel(obj interface{}) error`

Устанавливает модель для отображения таблицы.

```go
pageModel.SetTableModel(&User{})
```

---

##### `SetListModel(obj interface{}) error`

Устанавливает модель для получения списка.

```go
pageModel.SetListModel(&User{})
```

---

##### `SetFilterModel(obj interface{}, columns int) error`

Устанавливает модель для фильтрации.

```go
pageModel.SetFilterModel(&User{}, 2)
```

---

#### `QueryParams`

Параметры запроса.

```go
type QueryParams struct {
    // Содержит параметры GET/POST запроса
}
```

---

#### `MetaData`

Метаданные поля для поиска.

```go
type MetaData struct {
    MetaKey  string  // ключ для поиска (например: "id")
    MetaData string  // поле для отображения (например: "name")
}
```

---

#### `UIModel`

Модель пользовательского интерфейса.

```go
type UIModel struct {
    // приватные поля...
}
```

**Методы:**

##### `setModel(obj interface{}) error`

Устанавливает модель и инициализирует контейнеры.

---

##### `getUpdatePage(params *QueryParams, md map[string]interface{}) *Page`

Получает страницу редактирования.

---

##### `getCreatePage(params *QueryParams, md map[string]interface{}) *Page`

Получает страницу создания.

---

##### `getFilterPage(params *QueryParams, md map[string]interface{}) *Page`

Получает страницу фильтрации.

---

##### `getFieldsModel(obj interface{}) error`

Парсит структуру и извлекает поля.

---

#### `FieldType`

Информация о типе поля.

```go
type FieldType struct {
    Name           string  // имя поля
    Type           int     // внутренний тип
    JsonName       string  // имя в JSON
    Gorm           string  // GORM теги
    PgType         string  // тип для Page Generator
    PgText         string  // отображаемый текст
    PgReadOnly     bool    // только для чтения
    PgValid        string  // сообщение валидации
    pgMax          int     // максимальное значение
    pgMin          int     // минимальное значение
    pgEdit         bool    // редактируемое
    pgVisible      string  // видимость
    pgTemplate     string  // шаблон
    pgContainer    string  // ключ контейнера
    // ... другие поля
}
```

---

## 📦 Пакет `inputs`

### Структуры

#### `Container`

Контейнер для группировки элементов формы.

```go
type Container struct {
    Key          string      // уникальный ключ
    Title        string      // заголовок контейнера
    Orientation  string      // "vertical" или "horizontal"
    Direction    string      // "left", "right", "both", "center"
    Padding      float64     // внутренний отступ
    Margin       float64     // внешний отступ
    Border       float64     // толщина границы
    BorderRadius float64     // закругленность
    BackColor    string      // цвет фона
    Inputs       []Input     // элементы
    Childs       []Container // вложенные контейнеры
}
```

**Методы:**

##### `GetContainerByKey(key string) *Container`

Поиск контейнера по ключу рекурсивно.

```go
container := rootContainer.GetContainerByKey("personal")
```

---

##### `GetContainerByKeyInSlice(containers []Container, key string) *Container`

Поиск контейнера в срезе контейнеров.

```go
container := inputs.GetContainerByKeyInSlice(form.Containers, "personal")
```

---

#### `Input`

Элемент формы.

```go
type Input struct {
    Type           string     // тип (text-field, combo-box и т.д.)
    Name           string     // имя поля
    FromName       string     // альтернативное имя
    ReadOnly       bool       // только для чтения
    Text           string     // отображаемый текст
    MaxLength      int        // максимальная длина
    MinLength      int        // минимальная длина
    DefaultValue   string     // значение по умолчанию
    Items          ComboItems // элементы списка
    FileExtensions []string   // допустимые расширения
    // ... другие поля
}
```

---

#### `ComboItem`

Элемент для выпадающего меню.

```go
type ComboItem struct {
    ID   interface{} // идентификатор
    Text interface{} // отображаемый текст
}
```

---

#### `ComboItems`

Срез элементов для выпадающего меню.

```go
type ComboItems []ComboItem
```

---

#### `Form`

Структура формы.

```go
type Form struct {
    Name       string      // имя формы
    Containers []Container // контейнеры формы
}
```

---

#### `Page`

Главная структура страницы.

```go
type Page struct {
    Form      *Form      // форма
    DataTable *DataTable // таблица данных
}
```

---

#### `DataTable`

Таблица данных.

```go
type DataTable struct {
    Title      string        // название таблицы
    Header     []TableHeader // заголовки столбцов
    KeyColumn  string        // основной ключевой столбец
    PageSize   int           // размер страницы
    ItemsCount string        // URL для получения количества
    Delete     Action        // действие удаления
    Edit       LoadAction    // действие редактирования
    Add        LoadAction    // действие добавления
    Context    []Action      // контекстные действия
    DefaultUrl string        // URL по умолчанию
    Indexes    []Index       // индексы
    Type       string        // тип таблицы
    Exports    Export        // экспорт
    // ... другие поля
}
```

---

#### `TableHeader`

Заголовок столбца таблицы.

```go
type TableHeader struct {
    Key          string       // ключ столбца
    Title        string       // название
    IsExportable bool         // экспортируемый
    Type         string       // тип
    Element      TableElement // элемент
    TableSubmit  TableSubmit  // отправка
    Order        int          // порядок
    Template     string       // шаблон
    Access       []string     // права доступа
}
```

---

#### `Action`

Действие (кнопка, ссылка).

```go
type Action struct {
    Type       string // тип действия
    Source     string // источник (URL)
    Method     string // HTTP метод
    Text       string // отображаемый текст
    Message    string // сообщение
    LastAction string // последующее действие
}
```

---

#### `LoadAction`

Действие загрузки страницы/диалога.

```go
type LoadAction struct {
    Source string // источник (URL)
    Action string // "dialog" или "tab"
    Text   string // отображаемый текст
    Access []string // права доступа
}
```

---

#### `Export`

Форматы экспорта.

```go
type Export struct {
    Word  bool // экспорт в Word
    Excel bool // экспорт в Excel
    PDF   bool // экспорт в PDF
}
```

---

#### `Index`

Индекс (итоговая функция).

```go
type Index struct {
    Function      string // функция (count, sum, avg и т.д.)
    Column        string // столбец
    ColumnIdentity string // идентификатор
    Title         string // название
}
```

---

## 🎨 Типы Input'ов

| Тип | Описание | Пример |
|-----|----------|--------|
| `text-field` | Текстовое поле | `pgType:"text-field"` |
| `text-view` | Текст (только чтение) | `pgType:"text-view"` |
| `combo-box` | Выпадающее меню | `pgType:"combo-box"` |
| `date-time` | Дата и время | `pgType:"date-time"` |
| `number-view` | Числовое поле | `pgType:"number-view"` |
| `check-box` | Чекбокс | `pgType:"check-box"` |
| `label` | Метка | `pgType:"label"` |
| `search-view` | Поле поиска | `pgType:"search-view"` |
| `hidden` | Скрытое поле | `pgType:"hidden"` |
| `auto-complete` | Автодополнение | `pgType:"auto-complete"` |
| `file-uploader` | Загрузка файла | `pgType:"file-uploader"` |

---

## 🏷️ Доступные теги

| Тег | Значение | Пример |
|-----|----------|--------|
| `pg` | Исключение поля | `pg:"-"` |
| `pgType` | Тип Input'а | `pgType:"text-field"` |
| `pgText` | Отображаемый текст | `pgText:"Имя"` |
| `pgReadOnly` | Только чтение | `pgReadOnly:"true"` |
| `pgValid` | Сообщение валидации | `pgValid:"Обязательное поле"` |
| `pgMax` | Максимальное значение | `pgMax:"100"` |
| `pgMin` | Минимальное значение | `pgMin:"0"` |
| `pgEdit` | Редактируемое | `pgEdit:"true"` |
| `pgVisible` | Видимость | `pgVisible:"true"` |
| `pgTemplate` | Шаблон отображения | `pgTemplate:"<b>{{value}}</b>"` |
| `pgFileSource` | Источник загрузки | `pgFileSource:"/api/upload"` |
| `pgFileMaxSize` | Макс размер файла | `pgFileMaxSize:"5242880"` |
| `pgFromName` | Альтернативное имя | `pgFromName:"alt_name"` |
| `pgSearch` | Источник поиска | `pgSearch:"/api/search"` |
| `pgSName` | Объект поиска | `pgSName:"search_param"` |
| `pgContainer` | Ключ контейнера | `pgContainer:"personal"` |
| `dtTemp` | Шаблон таблицы | `dtTemp:"<b>{{value}}</b>"` |
| `dtTitle` | Заголовок столбца | `dtTitle:"Имя"` |
| `dtExport` | Экспортируемое | `dtExport:"true"` |
| `json` | JSON имя | `json:"field_name"` |
| `gorm` | GORM теги | `gorm:"primaryKey"` |

---

## ⚙️ Константы

### Типы действий

```go
const (
    DeleteAction     = "delete"
    LoadTabAction    = "load"
    LoadDialogAction = "loadDialog"
    LoadHTML         = "loadHtml"
)
```

### Типы загрузки

```go
const (
    loadDialog = "dialog"
    LoadTab    = "tab"
)
```

---

## 🔄 HTTP методы

- `GET` - получение данных
- `POST` - создание данных
- `PUT` - обновление данных
- `DELETE` - удаление данных

---

## 📋 Таблица совместимости интерфейсов и типов

| Интерфейс | Тип Input | Использование |
|-----------|----------|---------------|
| `IComboBox` | combo-box, search-view | Список значений |
| `ICompleteNodes` | auto-complete | Автодополнение |
| `IFileExtensions` | file-uploader | Расширения файлов |
| `IMetaData` | search-view | Метаданные поиска |
| `IClearNodes` | combo-box | Очистка связанных полей |
| `IDefault` | все типы | Значения по умолчанию |

---

## 🚀 Быстрые старты

### Минимальная конфигурация

```go
type User struct {
    ID   int    `pg:"-" gorm:"primaryKey"`
    Name string `pgType:"text-field" pgContainer:"main"`
}

func (u *User) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "main", Title: "Основное", Inputs: []inputs.Input{}},
    }
}

// ... остальные методы ...
```

---

## 📚 Дополнительные ресурсы

- [Документация](./DOCUMENTATION.md)
- [Примеры](./EXAMPLES.md)
- [Быстрый старт](./QUICKSTART.md)
- [GitHub репозиторий](https://github.com/BekkkEvrika/page_generator)

---

**Версия:** 1.0  
**Последнее обновление:** 2026-03-06
