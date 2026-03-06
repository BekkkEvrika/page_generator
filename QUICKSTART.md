# 🚀 Быстрый старт Page Generator

Этот гайд поможет вам начать использовать Page Generator за несколько минут.

---

## 📦 Установка

```bash
go get github.com/BekkkEvrika/page_generator
```

---

## ⚡ Самый простой пример (5 минут)

### 1. Создайте структуру модели

```go
package models

import "github.com/BekkkEvrika/page_generator/inputs"

type Product struct {
    ID    int    `pg:"-" gorm:"primaryKey" json:"id"`
    Name  string `pgType:"text-field" pgText:"Название" pgContainer:"main" json:"name"`
    Price float64 `pgType:"number-view" pgText:"Цена" pgContainer:"main" json:"price"`
}

// Обязательно: реализуйте IModel
func (p *Product) GetContainers() []inputs.Container {
    return []inputs.Container{
        {
            Key:    "main",
            Title:  "Основная информация",
            Inputs: []inputs.Input{},
        },
    }
}

// Остальные методы (временно пустые)
func (p *Product) Create(params *pg.QueryParams) error { return nil }
func (p *Product) Update(params *pg.QueryParams) error { return nil }
func (p *Product) Delete(params *pg.QueryParams) error { return nil }
func (p *Product) GetDefault(params *pg.QueryParams, md map[string]interface{}) map[string]string {
    return map[string]string{}
}
func (p *Product) GetComboItems(params *pg.QueryParams, md map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{}
}
func (p *Product) GetCompleteNodes() map[string][]string { return map[string][]string{} }
func (p *Product) GetFileExtensions() map[string][]string { return map[string][]string{} }
func (p *Product) GetList(params *pg.QueryParams) error { return nil }
func (p *Product) Filter(obj interface{}, params *pg.QueryParams) error { return nil }
func (p *Product) GetCount(params *pg.QueryParams) (int, error) { return 0, nil }
func (p *Product) GetContextActions() []inputs.Action { return []inputs.Action{} }
func (p *Product) GetIndexes() []inputs.Index { return []inputs.Index{} }
func (p *Product) GetExports() inputs.Export { return inputs.Export{} }
func (p *Product) GetDefaultQueryParams() map[string]string { return map[string]string{} }
func (p *Product) GetMetaData() map[string]pg.MetaData { return map[string]pg.MetaData{} }
func (p *Product) GetClearNodes() map[string][]string { return map[string][]string{} }
func (p *Product) GetEditPage() inputs.LoadAction {
    return inputs.LoadAction{Source: "/api/products/edit", Action: "dialog"}
}
```

### 2. Инициализируйте в main.go

```go
package main

import (
    "github.com/gin-gonic/gin"
    pg "github.com/BekkkEvrika/page_generator"
    "yourapp/models"
)

func main() {
    router := gin.Default()

    // Инициализация Page Generator
    err := pg.SetDefinitions(func() error {
        // Здесь может быть код для регистрации моделей
        return nil
    }, pg.PageSetting{
        Service:    "my-app",
        DateFormat: "2006-01-02",
        PageSize:   20,
    })

    if err != nil {
        panic(err)
    }

    // Регистрация маршрутов
    err = pg.GetModelsRoutes(router)
    if err != nil {
        panic(err)
    }

    router.Run(":8080")
}
```

### 3. Готово!

Теперь у вас есть:
- ✅ Форма создания товара
- ✅ Форма редактирования товара
- ✅ Таблица товаров
- ✅ Фильтрация товаров

---

## 🎯 Следующие шаги

### 1. Добавьте больше полей

```go
type Product struct {
    ID          int    `pg:"-" gorm:"primaryKey" json:"id"`
    Name        string `pgType:"text-field" pgText:"Название" pgContainer:"info" json:"name"`
    Description string `pgType:"text-view" pgText:"Описание" pgContainer:"info" json:"description"`
    Category    string `pgType:"combo-box" pgText:"Категория" pgContainer:"category" json:"category"`
    Price       float64 `pgType:"number-view" pgText:"Цена" pgContainer:"pricing" json:"price"`
    Discount    float64 `pgType:"number-view" pgText:"Скидка (%)" pgContainer:"pricing" json:"discount"`
    InStock     bool   `pgType:"check-box" pgText:"В наличии" pgContainer:"info" json:"in_stock"`
}
```

### 2. Добавьте контейнеры

```go
func (p *Product) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "info", Title: "Основная информация", Inputs: []inputs.Input{}},
        {Key: "category", Title: "Категория", Inputs: []inputs.Input{}},
        {Key: "pricing", Title: "Цены", Inputs: []inputs.Input{}},
    }
}
```

### 3. Реализуйте CRUD операции

```go
func (p *Product) Create(params *pg.QueryParams) error {
    // Сохранить в БД
    // db.Create(p)
    return nil
}

func (p *Product) Update(params *pg.QueryParams) error {
    // Обновить в БД
    // db.Save(p)
    return nil
}

func (p *Product) Delete(params *pg.QueryParams) error {
    // Удалить из БД
    // db.Delete(p)
    return nil
}
```

### 4. Добавьте выпадающие меню (IComboBox)

```go
func (p *Product) GetComboItems(params *pg.QueryParams, md map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{
        "category": {  // соответствует json:"category"
            {ID: 1, Text: "Электроника"},
            {ID: 2, Text: "Книги"},
            {ID: 3, Text: "Одежда"},
        },
    }
}
```

### 5. Добавьте значения по умолчанию (IDefault)

```go
func (p *Product) GetDefault(params *pg.QueryParams, md map[string]interface{}) map[string]string {
    return map[string]string{
        "in_stock": "true",  // соответствует json:"in_stock"
        "discount": "0",     // соответствует json:"discount"
    }
}
```

---

## 📚 Понимание контейнеров

### Что такое контейнер?

Контейнер - это контейнер для группировки связанных полей в форме.

```go
// Определите контейнеры
func (p *Product) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "main", Title: "Основное"},
        {Key: "advanced", Title: "Расширенное"},
    }
}

// Используйте контейнеры в тегах
type Product struct {
    Name string `pgContainer:"main"`     // будет в контейнере "main"
    SKU  string `pgContainer:"advanced"` // будет в контейнере "advanced"
}
```

### Вложенные контейнеры

```go
func (e *Employee) GetContainers() []inputs.Container {
    return []inputs.Container{
        {
            Key: "main",
            Title: "Основное",
            Childs: []inputs.Container{
                {Key: "basic", Title: "Базовые данные"},
                {Key: "extended", Title: "Расширенные данные"},
            },
        },
    }
}
```

---

## 🔍 Отладка

### Проверка структуры

```go
// Убедитесь, что структура реализует IModel
var _ pg.IModel = (*Product)(nil)

// Убедитесь, что структура реализует ICreate
var _ pg.ICreate = (*Product)(nil)

// И так далее для остальных интерфейсов...
```

### Логирование

```go
// Добавьте логирование в методы
func (p *Product) Create(params *pg.QueryParams) error {
    log.Printf("Creating product: %+v\n", p)
    // ... логика создания
    return nil
}
```

---

## ⚠️ Частые ошибки

### Ошибка 1: Забыли реализовать IModel

```go
// ❌ ОШИБКА
type Product struct {
    ID   int
    Name string
}
// Нет GetContainers()!

// ✅ ИСПРАВИТЬ
func (p *Product) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "main", Title: "Основное", Inputs: []inputs.Input{}},
    }
}
```

### Ошибка 2: Неправильное имя контейнера

```go
// ❌ ОШИБКА
func (p *Product) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "main", Title: "Основное", Inputs: []inputs.Input{}},
    }
}

type Product struct {
    Name string `pgContainer:"other"` // контейнер "other" не существует!
}

// ✅ ИСПРАВИТЬ
type Product struct {
    Name string `pgContainer:"main"` // используйте существующий контейнер
}
```

### Ошибка 3: Неправильные ключи в map-интерфейсах

```go
// ❌ ОШИБКА
func (p *Product) GetComboItems(params *pg.QueryParams, md map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{
        "Category": {  // не соответствует json тегу!
            {ID: 1, Text: "Электроника"},
        },
    }
}

// ✅ ИСПРАВИТЬ
func (p *Product) GetComboItems(params *pg.QueryParams, md map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{
        "category": {  // соответствует json:"category"
            {ID: 1, Text: "Электроника"},
        },
    }
}
```

### Ошибка 4: Забыли читать из params

```go
// ❌ ОШИБКА
func (p *Product) Update(params *pg.QueryParams) error {
    // params не используются!
    db.Save(p)
    return nil
}

// ✅ ИСПРАВИТЬ
func (p *Product) Update(params *pg.QueryParams) error {
    // Используйте params для получения данных из формы
    // params содержит данные, отправленные пользователем
    db.Save(p)
    return nil
}
```

---

## 💡 Советы и трюки

### 1. Группировка полей по контейнерам

```go
func (p *Product) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "basic", Title: "Базовые данные"},
        {Key: "details", Title: "Детали"},
        {Key: "seo", Title: "SEO"},
    }
}
```

### 2. Условная видимость полей

```go
// Используйте pgVisible для скрытия полей
type Product struct {
    Price       float64 `pgType:"number-view"`
    SpecialPrice float64 `pgType:"number-view" pgVisible:"false"` // скрыто по умолчанию
}
```

### 3. Только для чтения поля

```go
// Используйте pgReadOnly для полей, которые нельзя редактировать
type Product struct {
    ID        int       `pgType:"label" pgReadOnly:"true"`
    CreatedAt time.Time `pgType:"date-time" pgReadOnly:"true"`
}
```

### 4. Валидация полей

```go
type Product struct {
    Name string `pgType:"text-field" pgValid:"Название обязательно" pgMax:"100"`
    Price float64 `pgType:"number-view" pgMin:"0" pgMax:"1000000"`
}
```

---

## 🚀 Развертывание

### Docker

```dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
```

### Environment переменные

```go
// Используйте переменные окружения
config := pg.PageSetting{
    Service:    os.Getenv("SERVICE_NAME"),
    DateFormat: os.Getenv("DATE_FORMAT"),
    PageSize:   parseInt(os.Getenv("PAGE_SIZE")),
}
```

---

## 🎓 Темы для дальнейшего изучения

1. **Интеграция с БД** - GORM, MongoDB
2. **Авторизация** - Keycloak, JWT
3. **Экспорт данных** - Excel, PDF, Word
4. **Кастомизация UI** - CSS, JavaScript
5. **Производительность** - кэширование, оптимизация

---

## ✅ Чек-лист для начала работы

- [ ] Установлена библиотека
- [ ] Создана структура модели
- [ ] Реализован интерфейс IModel
- [ ] Реализованы методы CRUD (ICreate, IUpdate, IDelete)
- [ ] Инициализирована библиотека в main.go
- [ ] Зарегистрированы маршруты
- [ ] Проверены контейнеры
- [ ] Проверены типы полей
- [ ] Работает создание записи
- [ ] Работает редактирование записи
- [ ] Работает удаление записи
- [ ] Работает таблица с данными

---

**Готовы начать? Начните с примера выше и постепенно добавляйте функциональность! 🚀**

---

**Версия:** 1.0  
**Последнее обновление:** 2026-03-06
