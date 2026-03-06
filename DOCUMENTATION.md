# 📚 Документация Page Generator

Полная документация по использованию библиотеки Page Generator для автоматической генерации веб-интерфейсов.

---

## 📋 Оглавление

1. [Введение](#введение)
2. [Интерфейсы](#интерфейсы)
   - [Интерфейсы моделей](#интерфейсы-моделей)
   - [Интерфейсы таблиц и списков](#интерфейсы-таблиц-и-списков)
3. [Теги структур](#теги-структур)
4. [Типы Input'ов](#типы-inputов)
5. [Контейнеры](#контейнеры)
6. [Примеры](#примеры)
7. [API Reference](#api-reference)

---

## 🎯 Введение

Page Generator - это мощная Go-библиотека для автоматической генерации веб-интерфейсов (форм, таблиц, страниц) на основе структур данных и тегов. Библиотека интегрируется с фреймворком Gin и Keycloak для авторизации.

### Основные возможности

- 🔧 **Автоматическая генерация** форм на основе структур Go
- 📊 **Генерация таблиц** с сортировкой и фильтрацией
- 🏗️ **Контейнерная система** для группировки полей
- 🔐 **Интеграция с Keycloak** для авторизации
- 🎨 **Настраиваемые типы полей** (текст, числа, даты, файлы и т.д.)
- 🌐 **REST API** готовые маршруты
- 📱 **JSON API** для фронтенда

### Важные особенности

- **IModel** - обязательный интерфейс для всех моделей
- **Другие интерфейсы** - делают соответствующие функции доступными через UI
- **Map-интерфейсы** - используют JSON теги полей как ключи
- **Разделение ответственности** - разные интерфейсы для форм, таблиц и списков

---

## 🔌 Интерфейсы

Для расширения функциональности библиотеки вы можете реализовать следующие интерфейсы. Интерфейсы разделены на категории в зависимости от их назначения.

### Интерфейсы моделей (для форм)

Эти интерфейсы используются для работы с формами создания/редактирования записей.

#### 1. **IModel** - Основной интерфейс (ОБЯЗАТЕЛЬНЫЙ)

Предоставляет контейнеры для структуры формы.

```go
type IModel interface {
    GetContainers() []inputs.Container
}
```

**Метод:**
- `GetContainers()` - возвращает список контейнеров для формы

**Пример реализации:**
```go
type User struct {
    ID    int    `pg:"-" gorm:"primaryKey"`
    Name  string `pgType:"text-field" pgContainer:"personal"`
    Email string `pgType:"text-field" pgContainer:"contact"`
}

func (u *User) GetContainers() []inputs.Container {
    return []inputs.Container{
        {
            Key:   "personal",
            Title: "Личные данные",
            Inputs: []inputs.Input{},
        },
        {
            Key:   "contact",
            Title: "Контактная информация",
            Inputs: []inputs.Input{},
        },
    }
}
```

---

#### 2. **ICreate** - Создание записи

Интерфейс для обработки операции создания новой записи. **Делает функцию создания доступной через UI.**

```go
type ICreate interface {
    Create(params *QueryParams) error
}
```

**Метод:**
- `Create(params)` - создает новую запись из параметров запроса

**Пример реализации:**
```go
func (u *User) Create(params *QueryParams) error {
    // Логика создания пользователя
    // db.Create(u)
    return nil
}
```

---

#### 3. **IUpdate** - Обновление записи

Интерфейс для обработки операции обновления записи. **Делает функцию обновления доступной через UI.**

```go
type IUpdate interface {
    Update(params *QueryParams) error
}
```

**Метод:**
- `Update(params)` - обновляет запись из параметров запроса

**Пример реализации:**
```go
func (u *User) Update(params *QueryParams) error {
    // Логика обновления пользователя
    // db.Save(u)
    return nil
}
```

---

#### 4. **IDelete** - Удаление записи

Интерфейс для обработки операции удаления записи. **Делает функцию удаления доступной через UI.**

```go
type IDelete interface {
    Delete(params *QueryParams) error
}
```

**Метод:**
- `Delete(params)` - удаляет запись из параметров запроса

**Пример реализации:**
```go
func (u *User) Delete(params *QueryParams) error {
    // Логика удаления пользователя
    // db.Delete(u)
    return nil
}
```

---

#### 5. **IDefault** - Значения по умолчанию

Интерфейс для предоставления значений по умолчанию для полей. **Ключи должны соответствовать JSON тегам полей.**

```go
type IDefault interface {
    GetDefault(params *QueryParams, mp map[string]interface{}) map[string]string
}
```

**Метод:**
- `GetDefault(params, metadata)` - возвращает map с значениями по умолчанию для каждого поля

**Пример реализации:**
```go
func (u *User) GetDefault(params *QueryParams, md map[string]interface{}) map[string]string {
    return map[string]string{
        "name": "John Doe",        // соответствует json:"name"
        "email": "john@example.com", // соответствует json:"email"
    }
}
```

---

#### 6. **IComboBox** - Списки выбора (выпадающие меню)

Интерфейс для предоставления элементов для выпадающих меню и комбо-боксов. **Ключи должны соответствовать JSON тегам полей.**

```go
type IComboBox interface {
    GetComboItems(params *QueryParams, mp map[string]interface{}) map[string]inputs.ComboItems
}
```

**Метод:**
- `GetComboItems(params, metadata)` - возвращает map с элементами для каждого поля

**Структура ComboItems:**
```go
type ComboItem struct {
    ID   interface{} `json:"id"`
    Text interface{} `json:"text"`
}
type ComboItems []ComboItem
```

**Пример реализации:**
```go
func (u *User) GetComboItems(params *QueryParams, md map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{
        "status": {  // соответствует json:"status"
            {ID: 1, Text: "Активный"},
            {ID: 2, Text: "Неактивный"},
        },
        "role": {    // соответствует json:"role"
            {ID: "admin", Text: "Администратор"},
            {ID: "user", Text: "Пользователь"},
        },
    }
}
```

---

#### 7. **ICompleteNodes** - Автодополнение

Интерфейс для предоставления элементов автодополнения (autocomplete). **Ключи должны соответствовать JSON тегам полей.**

```go
type ICompleteNodes interface {
    GetCompleteNodes() map[string][]string
}
```

**Метод:**
- `GetCompleteNodes()` - возвращает map со списками для автодополнения

**Пример реализации:**
```go
func (u *User) GetCompleteNodes() map[string][]string {
    return map[string][]string{
        "city": {"Moscow", "Saint Petersburg", "Kazan"},  // соответствует json:"city"
        "country": {"Russia", "Kazakhstan", "Belarus"},   // соответствует json:"country"
    }
}
```

---

#### 8. **IFileExtensions** - Расширения файлов

Интерфейс для указания допустимых расширений файлов для загрузки. **Ключи должны соответствовать JSON тегам полей.**

```go
type IFileExtensions interface {
    GetFileExtensions() map[string][]string
}
```

**Метод:**
- `GetFileExtensions()` - возвращает map со списками расширений для каждого поля

**Пример реализации:**
```go
func (u *User) GetFileExtensions() map[string][]string {
    return map[string][]string{
        "avatar": {"jpg", "png", "gif"},      // соответствует json:"avatar"
        "document": {"pdf", "doc", "docx"},   // соответствует json:"document"
    }
}
```

---

#### 9. **IMetaData** - Метаданные поля

Интерфейс для предоставления метаданных полей (ключ и значение для поиска). **Ключи должны соответствовать JSON тегам полей.**

```go
type IMetaData interface {
    GetMetaData() map[string]MetaData
}
```

**Структура MetaData:**
```go
type MetaData struct {
    MetaKey  string
    MetaData string
}
```

**Метод:**
- `GetMetaData()` - возвращает map с метаданными для каждого поля

**Пример реализации:**
```go
func (u *User) GetMetaData() map[string]MetaData {
    return map[string]MetaData{
        "department": {  // соответствует json:"department"
            MetaKey: "id",
            MetaData: "name",
        },
    }
}
```

---

#### 10. **IClearNodes** - Очистка связанных полей

Интерфейс для указания полей, которые нужно очистить при изменении другого поля. **Ключи должны соответствовать JSON тегам полей.**

```go
type IClearNodes interface {
    GetClearNodes() map[string][]string
}
```

**Метод:**
- `GetClearNodes()` - возвращает map с полями для очистки

**Пример реализации:**
```go
func (u *User) GetClearNodes() map[string][]string {
    return map[string][]string{
        "country": {"city", "region"},  // соответствует json:"country"
    }
}
```

---

### Интерфейсы таблиц и списков

Эти интерфейсы используются для работы с таблицами и списками данных.

#### 11. **IGetList** - Получение списка записей

Интерфейс для получения списка записей с фильтрацией. **Делает функцию получения списка доступной через UI.**

```go
type IGetList interface {
    GetList(params *QueryParams) error
    Filter(obj interface{}, params *QueryParams) error
}
```

**Методы:**
- `GetList(params)` - получает список записей
- `Filter(obj, params)` - применяет фильтр к объекту

**Пример реализации:**
```go
func (u *User) GetList(params *QueryParams) error {
    // Получить список пользователей
    return nil
}

func (u *User) Filter(obj interface{}, params *QueryParams) error {
    // Применить фильтр
    return nil
}
```

---

#### 12. **IPagination** - Пагинация

Интерфейс для получения количества записей (для пагинации). **Делает пагинацию доступной через UI.**

```go
type IPagination interface {
    GetCount(params *QueryParams) (int, error)
}
```

**Метод:**
- `GetCount(params)` - возвращает общее количество записей

**Пример реализации:**
```go
func (u *User) GetCount(params *QueryParams) (int, error) {
    // var count int64
    // db.Model(&User{}).Count(&count)
    return int(count), nil
}
```

---

#### 13. **IQueryParams** - Параметры запроса

Интерфейс для предоставления параметров запроса по умолчанию.

```go
type IQueryParams interface {
    GetDefaultQueryParams() map[string]string
}
```

**Метод:**
- `GetDefaultQueryParams()` - возвращает map с параметрами по умолчанию

**Пример реализации:**
```go
func (u *User) GetDefaultQueryParams() map[string]string {
    return map[string]string{
        "status": "active",
        "limit": "20",
    }
}
```

---

#### 14. **IContext** - Контекстные действия

Интерфейс для предоставления контекстного меню действий. **Делает контекстное меню доступным через UI.**

```go
type IContext interface {
    GetContextActions() []inputs.Action
}
```

**Метод:**
- `GetContextActions()` - возвращает список действий контекстного меню

**Пример реализации:**
```go
func (u *User) GetContextActions() []inputs.Action {
    return []inputs.Action{
        {Type: "edit", Source: "/api/users/edit", Method: "GET"},
        {Type: "delete", Source: "/api/users/delete", Method: "DELETE"},
    }
}
```

---

#### 15. **IIndexes** - Индексы

Интерфейс для определения индексов (функции подсчета, суммирования и т.д.). **Делает итоговые функции доступными через UI.**

```go
type IIndexes interface {
    GetIndexes() []inputs.Index
}
```

**Метод:**
- `GetIndexes()` - возвращает список индексов

**Пример реализации:**
```go
func (u *User) GetIndexes() []inputs.Index {
    return []inputs.Index{
        {Function: "count", Column: "id", ColumnIdentity: "id"},
        {Function: "sum", Column: "salary", ColumnIdentity: "salary"},
    }
}
```

---

#### 16. **IExports** - Экспорт данных

Интерфейс для указания доступных форматов экспорта. **Делает экспорт доступным через UI.**

```go
type IExports interface {
    GetExports() inputs.Export
}
```

**Метод:**
- `GetExports()` - возвращает доступные форматы экспорта

**Пример реализации:**
```go
func (u *User) GetExports() inputs.Export {
    return inputs.Export{
        Word: true,   // экспорт в Word
        Excel: true,  // экспорт в Excel
        PDF: true,    // экспорт в PDF
    }
}
```

---

#### 17. **IEditData** - Данные редактирования

Интерфейс для предоставления действия редактирования (открыть страницу/диалог редактирования).

```go
type IEditData interface {
    GetEditPage() inputs.LoadAction
}
```

**Метод:**
- `GetEditPage()` - возвращает действие редактирования

**Пример реализации:**
```go
func (u *User) GetEditPage() inputs.LoadAction {
    return inputs.LoadAction{
        Source: "/api/users/edit",
        Action: "dialog",  // или "tab"
        Text: "Редактировать",
    }
}
```

---

## 🏷️ Теги структур

### Основные теги для полей структур:

#### 1. **pg** - Исключение поля из формы

```go
field int `pg:"-"`  // поле будет исключено из формы
```

**Важно:** Тег `pg` используется ТОЛЬКО для исключения полей из UI. Единственное допустимое значение - `"-"`.

**Пример:**
```go
type User struct {
    ID        int       `pg:"-"`                    // исключено из формы
    Name      string    `pgType:"text-field"`       // отображается в форме
    CreatedAt time.Time `pg:"-"`                    // исключено из формы
}
```

---

#### 2. **pgType** - Тип input'а (ОСНОВНОЙ ТЕГ)

```go
field string `pgType:"text-field"`
```

**Основной тег для указания типа input'а.** Определяет, как поле будет отображаться в форме.

Возможные значения:
- `"text-field"` - текстовое поле
- `"combo-box"` - выпадающее меню
- `"date-time"` - выбор даты и времени
- `"number-view"` - числовое поле
- `"check-box"` - чекбокс
- `"label"` - метка (только для чтения)
- `"search-view"` - поле поиска
- `"hidden"` - скрытое поле
- `"auto-complete"` - автодополнение
- `"file-uploader"` - загрузка файла

---

#### 3. **pgText** - Отображаемый текст/метка поля

```go
field string `pgText:"Имя пользователя"`
```

Это текст, который будет отображаться как метка поля в форме.

---

#### 4. **pgReadOnly** - Поле только для чтения

```go
field string `pgReadOnly:"true"`
```

Значения:
- `"true"` - поле только для чтения
- `"false"` - поле для редактирования (по умолчанию)

---

#### 5. **pgValid** - Валидационное сообщение

```go
field string `pgValid:"Email должен быть валидным"`
```

Текст сообщения, которое будет отображаться при ошибке валидации.

---

#### 6. **pgMax** / **pgMin** - Максимальная и минимальная длина

```go
field string `pgMax:"100" pgMin:"5"`
```

Числовые значения для ограничения длины текста или границ числа.

---

#### 7. **pgEdit** - Редактируемое поле

```go
field string `pgEdit:"true"`
```

Значения:
- `"true"` - поле редактируется при редактировании записи
- `"false"` - поле не редактируется (по умолчанию для ID)

---

#### 8. **pgVisible** - Видимость поля

```go
field string `pgVisible:"true"`
```

Значения:
- `"true"` - поле видимо
- `"false"` - поле скрыто

---

#### 9. **pgTemplate** - Шаблон отображения

```go
field string `pgTemplate:"<span>{{value}}</span>"`
```

HTML-шаблон для отображения значения поля.

---

#### 10. **pgFileSource** - Источник для загрузки файлов

```go
field string `pgFileSource:"/api/upload"`
```

URL для отправки загруженного файла.

---

#### 11. **pgFileMaxSize** - Максимальный размер файла

```go
field string `pgFileMaxSize:"5242880"`
```

Размер в байтах. Пример: 5242880 = 5MB

---

#### 12. **pgFromName** - Альтернативное имя при отправке

```go
field string `pgFromName:"alternate_name"`
```

Используется, если имя параметра при отправке отличается от имени поля.

---

#### 13. **pgSearch** / **pgSearchSource** - Источник поиска

```go
field string `pgSearch:"/api/search"`
```

URL для поиска значений (для search-view типов).

---

#### 14. **pgSName** / **pgSearchObject** - Объект для поиска

```go
field string `pgSName:"search_param"`
```

Имя параметра для поиска.

---

#### 15. **pgContainer** - Контейнер формы

```go
field string `pgContainer:"personal"`
```

Ключ контейнера, в который будет добавлено это поле. Значение должно соответствовать ключу контейнера, возвращаемому методом `GetContainers()`.

---

### Теги для таблиц:

#### 1. **dtTemp** - Шаблон для таблицы

```go
field string `dtTemp:"<b>{{value}}</b>"`
```

HTML-шаблон для отображения в таблице.

---

#### 2. **dtTitle** - Заголовок столбца

```go
field string `dtTitle:"Имя"`
```

Текст, отображаемый как заголовок столбца в таблице.

---

#### 3. **dtExport** - Экспортируемое поле

```go
field string `dtExport:"true"`
```

Значения:
- `"true"` - поле экспортируется
- `"false"` - поле не экспортируется

---

#### 4. **json** - JSON тег (стандартный Go)

```go
field string `json:"field_name"`
```

Используется для сериализации/десериализации JSON. Это стандартный Go тег.

---

#### 5. **gorm** - GORM тег (если используется GORM)

```go
field string `gorm:"primaryKey"`
field string `gorm:"autoIncrement"`
```

Теги для GORM ORM (если используется).

---

## 🎯 Типы Input'ов

### 1. **text-field** - Текстовое поле

Простое однострочное текстовое поле.

```go
Name string `pgType:"text-field" pgText:"Имя"`
```

---

### 2. **text-view** - Просмотр текста

Текстовое поле только для чтения (по умолчанию ReadOnly=true).

```go
Description string `pgType:"text-view" pgText:"Описание"`
```

---

### 3. **combo-box** - Выпадающее меню

Требует реализации `IComboBox` интерфейса.

```go
Status string `pgType:"combo-box" pgText:"Статус"`
```

---

### 4. **date-time** - Выбор даты и времени

Поле для выбора даты и времени.

```go
CreatedAt time.Time `pgType:"date-time" pgText:"Дата создания"`
```

---

### 5. **number-view** - Числовое поле

Поле для ввода чисел.

```go
Age int `pgType:"number-view" pgText:"Возраст" pgMin:"0" pgMax:"150"`
```

---

### 6. **check-box** - Чекбокс

Булево значение (true/false).

```go
IsActive bool `pgType:"check-box" pgText:"Активный"`
```

---

### 7. **label** - Метка

Поле только для чтения, отображается как обычный текст.

```go
ID int `pgType:"label" pgText:"ID"`
```

---

### 8. **search-view** - Поле поиска

Поле с поиском по удаленному источнику.

```go
Department string `pgType:"search-view" pgText:"Отдел" pgSearch:"/api/departments"`
```

---

### 9. **hidden** - Скрытое поле

Поле, не отображаемое в форме, но отправляемое с данными.

```go
Token string `pgType:"hidden"`
```

---

### 10. **auto-complete** - Автодополнение

Требует реализации `ICompleteNodes` интерфейса.

```go
City string `pgType:"auto-complete" pgText:"Город"`
```

---

### 11. **file-uploader** - Загрузка файла

Поле для загрузки файла.

```go
Avatar string `pgType:"file-uploader" pgText:"Аватар" pgFileSource:"/api/upload" pgFileMaxSize:"5242880"`
```

---

## 🏗️ Контейнеры

### Что такое контейнер?

Контейнер - это логическая группа полей в форме. Контейнеры позволяют организовать поля в разделы и вкладки.

### Структура контейнера

```go
type Container struct {
    Key          string      // уникальный ключ контейнера
    Title        string      // заголовок контейнера
    Orientation  string      // "vertical" или "horizontal"
    Direction    string      // "left", "right", "both", "center"
    Padding      float64     // внутренний отступ
    Margin       float64     // внешний отступ
    Border       float64     // толщина границы
    BorderRadius float64     // закругленность границы
    BackColor    string      // цвет фона
    Inputs       []Input     // элементы в контейнере
    Childs       []Container // вложенные контейнеры
}
```

### Пример использования контейнеров

```go
type Employee struct {
    ID           int       `pg:"-" gorm:"primaryKey"`
    FirstName    string    `pgType:"text-field" pgText:"Имя" pgContainer:"personal"`
    LastName     string    `pgType:"text-field" pgText:"Фамилия" pgContainer:"personal"`
    Email        string    `pgType:"text-field" pgText:"Email" pgContainer:"contact"`
    Phone        string    `pgType:"text-field" pgText:"Телефон" pgContainer:"contact"`
    Department   string    `pgType:"combo-box" pgText:"Отдел" pgContainer:"work"`
    Position     string    `pgType:"combo-box" pgText:"Должность" pgContainer:"work"`
    Salary       float64   `pgType:"number-view" pgText:"Зарплата" pgContainer:"work"`
}

func (e *Employee) GetContainers() []inputs.Container {
    return []inputs.Container{
        {
            Key:   "personal",
            Title: "Личные данные",
            Inputs: []inputs.Input{},
        },
        {
            Key:   "contact",
            Title: "Контактная информация",
            Inputs: []inputs.Input{},
        },
        {
            Key:   "work",
            Title: "Рабочая информация",
            Inputs: []inputs.Input{},
        },
    }
}
```

### Вложенные контейнеры

```go
func (e *Employee) GetContainers() []inputs.Container {
    return []inputs.Container{
        {
            Key:   "main",
            Title: "Основное",
            Childs: []inputs.Container{
                {
                    Key:   "personal",
                    Title: "Личные данные",
                    Inputs: []inputs.Input{},
                },
                {
                    Key:   "contact",
                    Title: "Контактная информация",
                    Inputs: []inputs.Input{},
                },
            },
        },
    }
}
```

---

## 💡 Примеры

### Полный пример структуры с интерфейсами

```go
package models

import (
    "time"
    pg "github.com/BekkkEvrika/page_generator"
    "github.com/BekkkEvrika/page_generator/inputs"
)

type Employee struct {
    // Основные данные
    ID           int       `pg:"-" gorm:"primaryKey" json:"id"`
    FirstName    string    `pgType:"text-field" pgText:"Имя" pgContainer:"personal" json:"first_name" pgValid:"Имя обязательно"`
    LastName     string    `pgType:"text-field" pgText:"Фамилия" pgContainer:"personal" json:"last_name"`
    Email        string    `pgType:"text-field" pgText:"Email" pgContainer:"contact" json:"email"`
    Phone        string    `pgType:"text-field" pgText:"Телефон" pgContainer:"contact" json:"phone"`

    // Профессиональная информация
    Department   string    `pgType:"combo-box" pgText:"Отдел" pgContainer:"professional" json:"department"`
    Position     string    `pgType:"combo-box" pgText:"Должность" pgContainer:"professional" json:"position"`
    HireDate     time.Time `pgType:"date-time" pgText:"Дата приема" pgContainer:"professional" json:"hire_date"`
    Salary       float64   `pgType:"number-view" pgText:"Зарплата" pgContainer:"professional" pgMin:"0" json:"salary"`

    // Другая информация
    Photo        string    `pgType:"file-uploader" pgText:"Фото" pgFileSource:"/api/upload" pgContainer:"personal" json:"photo"`
    Status       string    `pgType:"combo-box" pgText:"Статус" pgContainer:"professional" json:"status"`
    IsActive     bool      `pgType:"check-box" pgText:"Активный" pgContainer:"professional" json:"is_active"`
    Notes        string    `pgType:"text-view" pgText:"Примечания" pgContainer:"other" json:"notes"`

    // Системные поля
    CreatedAt    time.Time `pgType:"label" pgText:"Создано" pgReadOnly:"true" json:"created_at"`
    UpdatedAt    time.Time `pgType:"label" pgText:"Обновлено" pgReadOnly:"true" json:"updated_at"`
}

// IModel - Контейнеры формы
func (e *Employee) GetContainers() []inputs.Container {
    return []inputs.Container{
        {
            Key:   "personal",
            Title: "Личные данные",
            Inputs: []inputs.Input{},
        },
        {
            Key:   "contact",
            Title: "Контактная информация",
            Inputs: []inputs.Input{},
        },
        {
            Key:   "professional",
            Title: "Профессиональная информация",
            Inputs: []inputs.Input{},
        },
        {
            Key:   "other",
            Title: "Прочее",
            Inputs: []inputs.Input{},
        },
    }
}

// ICreate - Создание записи
func (e *Employee) Create(params *pg.QueryParams) error {
    // db.Create(e)
    return nil
}

// IUpdate - Обновление записи
func (e *Employee) Update(params *pg.QueryParams) error {
    // db.Save(e)
    return nil
}

// IDelete - Удаление записи
func (e *Employee) Delete(params *pg.QueryParams) error {
    // db.Delete(e)
    return nil
}

// IDefault - Значения по умолчанию
func (e *Employee) GetDefault(params *pg.QueryParams, md map[string]interface{}) map[string]string {
    return map[string]string{
        "first_name": "John",        // соответствует json:"first_name"
        "last_name": "Doe",          // соответствует json:"last_name"
        "email": "john@example.com", // соответствует json:"email"
        "hire_date": time.Now().Format("2006-01-02"), // соответствует json:"hire_date"
        "is_active": "true",         // соответствует json:"is_active"
    }
}

// IComboBox - Элементы выпадающего меню
func (e *Employee) GetComboItems(params *pg.QueryParams, md map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{
        "department": {  // соответствует json:"department"
            {ID: "hr", Text: "Отдел кадров"},
            {ID: "it", Text: "IT"},
            {ID: "sales", Text: "Продажи"},
            {ID: "accounting", Text: "Бухгалтерия"},
        },
        "position": {    // соответствует json:"position"
            {ID: "manager", Text: "Менеджер"},
            {ID: "developer", Text: "Разработчик"},
            {ID: "analyst", Text: "Аналитик"},
            {ID: "specialist", Text: "Специалист"},
        },
        "status": {      // соответствует json:"status"
            {ID: "active", Text: "Активный"},
            {ID: "on_leave", Text: "В отпуске"},
            {ID: "inactive", Text: "Неактивный"},
        },
    }
}

// ICompleteNodes - Автодополнение
func (e *Employee) GetCompleteNodes() map[string][]string {
    return map[string][]string{
        "department": {"HR", "IT", "Sales", "Accounting"},  // соответствует json:"department"
    }
}

// IFileExtensions - Допустимые расширения файлов
func (e *Employee) GetFileExtensions() map[string][]string {
    return map[string][]string{
        "photo": {"jpg", "jpeg", "png", "gif"},  // соответствует json:"photo"
    }
}

// IGetList - Получение списка
func (e *Employee) GetList(params *pg.QueryParams) error {
    // var employees []Employee
    // db.Find(&employees)
    return nil
}

func (e *Employee) Filter(obj interface{}, params *pg.QueryParams) error {
    // Применить фильтры из параметров
    return nil
}

// IPagination - Количество записей
func (e *Employee) GetCount(params *pg.QueryParams) (int, error) {
    // var count int64
    // db.Model(&Employee{}).Count(&count)
    return 0, nil
}

// IContext - Контекстное меню
func (e *Employee) GetContextActions() []inputs.Action {
    return []inputs.Action{
        {Type: "edit", Source: "/api/employees/edit", Method: "GET", Text: "Редактировать"},
        {Type: "delete", Source: "/api/employees/delete", Method: "DELETE", Text: "Удалить"},
        {Type: "view", Source: "/api/employees/view", Method: "GET", Text: "Просмотр"},
    }
}

// IIndexes - Индексы (итоги)
func (e *Employee) GetIndexes() []inputs.Index {
    return []inputs.Index{
        {Function: "count", Column: "id", ColumnIdentity: "count", Title: "Всего сотрудников"},
        {Function: "sum", Column: "salary", ColumnIdentity: "total_salary", Title: "Сумма зарплат"},
        {Function: "avg", Column: "salary", ColumnIdentity: "avg_salary", Title: "Средняя зарплата"},
    }
}

// IExports - Экспорт
func (e *Employee) GetExports() inputs.Export {
    return inputs.Export{
        Word:  true,
        Excel: true,
        PDF:   true,
    }
}

// IQueryParams - Параметры по умолчанию
func (e *Employee) GetDefaultQueryParams() map[string]string {
    return map[string]string{
        "status": "active",
        "limit": "20",
        "offset": "0",
    }
}

// IMetaData - Метаданные для поиска
func (e *Employee) GetMetaData() map[string]pg.MetaData {
    return map[string]pg.MetaData{}
}

// IClearNodes - Очистка при изменении
func (e *Employee) GetClearNodes() map[string][]string {
    return map[string][]string{}
}

// IEditData - Редактирование
func (e *Employee) GetEditPage() inputs.LoadAction {
    return inputs.LoadAction{
        Source: "/api/employees/edit",
        Action: "dialog",
        Text: "Редактировать сотрудника",
    }
}
```

---

## 📚 API Reference

### Основные структуры

#### Container - Контейнер формы

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
- `GetContainerByKey(key string) *Container` - поиск контейнера по ключу
- `GetContainerByKeyInSlice(containers []Container, key string) *Container` - поиск в срезе

#### Input - Элемент формы

```go
type Input struct {
    Type           string     // тип (text-field, combo-box и т.д.)
    Name           string     // имя поля
    Text           string     // отображаемый текст
    ReadOnly       bool       // только для чтения
    MaxLength      int        // максимальная длина
    MinLength      int        // минимальная длина
    DefaultValue   string     // значение по умолчанию
    Items          ComboItems // элементы списка
    FileExtensions []string   // допустимые расширения
    // ... другие поля
}
```

#### Page - Главная структура страницы

```go
type Page struct {
    Form      *Form      // форма
    DataTable *DataTable // таблица данных
}
```

---

## 🔗 Ссылки и ресурсы

- [GitHub репозиторий](https://github.com/BekkkEvrika/page_generator)
- [Быстрый старт](./QUICKSTART.md)
- [Примеры](./EXAMPLES.md)
- [API Reference](./API_REFERENCE.md)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Keycloak](https://www.keycloak.org/)

---

## 📝 Лицензия

MIT License

---

**Версия документации:** 1.0  
**Последнее обновление:** 2026-03-06
