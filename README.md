# Page Generator

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Page Generator - это мощная Go-библиотека для автоматической генерации веб-интерфейсов (форм, таблиц, страниц) на основе структур данных и тегов. Библиотека интегрируется с фреймворком Gin и Keycloak для авторизации.

## 🚀 Быстрый старт

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
    pg.GetModelsRoutes(router)

    router.Run(":8080")
}
```

## 📚 Документация

- **[Быстрый старт](./QUICKSTART.md)** - начало работы за 5 минут
- **[Полная документация](./DOCUMENTATION.md)** - все интерфейсы и теги
- **[Примеры](./EXAMPLES.md)** - реальные примеры использования
- **[API Reference](./API_REFERENCE.md)** - справочник всех функций

## ✨ Особенности

- 🔧 **Автоматическая генерация** форм на основе структур Go
- 📊 **Генерация таблиц** с сортировкой и фильтрацией
- 🏗️ **Контейнерная система** для группировки полей
- 🔐 **Интеграция с Keycloak** для авторизации
- 🎨 **Настраиваемые типы полей** (текст, числа, даты, файлы и т.д.)
- 🌐 **REST API** готовые маршруты
- 📱 **JSON API** для фронтенда

## 🏗️ Архитектура

```
Page Generator
│
├── UIModel (управление моделью UI)
│   ├── getCreatePage() → форма создания
│   ├── getUpdatePage() → форма редактирования
│   ├── getFilterPage() → форма фильтрации
│   └── getFieldsModel() → парсинг полей
│
├── Container (контейнеры)
│   ├── GetContainerByKey() → поиск контейнера
│   └── Inputs[] → элементы формы
│
├── Input (элементы формы)
│   ├── text-field, combo-box, date-time...
│   └── валидация, значения по умолчанию
│
└── Интерфейсы
    ├── IModel → контейнеры
    ├── ICreate, IUpdate, IDelete → CRUD
    ├── IComboBox, IDefault → данные
    └── IFileExtensions, IMetaData → расширения
```

## 🏷️ Теги структур

### Основные теги

```go
type User struct {
    ID        int       `pg:"-" gorm:"primaryKey"`                    // исключить из формы
    Name      string    `pgType:"text-field" pgText:"Имя"`           // текстовое поле
    Email     string    `pgType:"text-field" pgContainer:"contact"`  // в контейнере contact
    Age       int       `pgType:"number-view" pgMin:"0" pgMax:"150"` // числовое поле с валидацией
    Status    string    `pgType:"combo-box"`                         // выпадающее меню
    Photo     string    `pgType:"file-uploader"`                     // загрузка файла
    IsActive  bool      `pgType:"check-box"`                         // чекбокс
    CreatedAt time.Time `pgType:"date-time" pgReadOnly:"true"`       // дата (только чтение)
}
```

### Типы полей

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

## 🔌 Интерфейсы

### Обязательные

```go
type IModel interface {
    GetContainers() []inputs.Container
}
```

### CRUD операции

```go
type ICreate interface { Create(params *QueryParams) error }
type IUpdate interface { Update(params *QueryParams) error }
type IDelete interface { Delete(params *QueryParams) error }
```

### Данные и настройки

```go
type IComboBox interface {
    GetComboItems(params *QueryParams, md map[string]interface{}) map[string]inputs.ComboItems
}

type IDefault interface {
    GetDefault(params *QueryParams, md map[string]interface{}) map[string]string
}
```

## 📦 Установка

```bash
go get github.com/BekkkEvrika/page_generator
```

## 🧪 Тестирование

```bash
go test ./...
```

## 📄 Лицензия

MIT License - см. файл [LICENSE](LICENSE)

## 🤝 Вклад в проект

1. Fork проект
2. Создайте feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit изменения (`git commit -m 'Add some AmazingFeature'`)
4. Push в branch (`git push origin feature/AmazingFeature`)
5. Откройте Pull Request

## 📞 Поддержка

- 📖 [Документация](./DOCUMENTATION.md)
- 💻 [Примеры](./EXAMPLES.md)
- 🐛 [Issues](https://github.com/BekkkEvrika/page_generator/issues)
- 📧 Email: support@example.com

---

**⭐ Если проект полезен, поставьте звезду!**

---

**Версия:** 1.0.0  
**Последнее обновление:** 2026-03-06
