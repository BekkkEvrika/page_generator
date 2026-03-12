# page_generator

Go-библиотека для автоматической генерации структуры страниц (форм и таблиц) для фронтенда на основе Go-структур с тегами. Интегрируется с [Gin](https://github.com/gin-gonic/gin).

## Установка

```bash
go get github.com/BekkkEvrika/page_generator
```

## Быстрый старт

```go
package main

import (
    "log"
    "github.com/gin-gonic/gin"
    pg "github.com/BekkkEvrika/page_generator"
    "github.com/BekkkEvrika/page_generator/inputs"
)

// 1. Модель с тегами
type UserModel struct {
    ID   int    `json:"id"   gorm:"primaryKey;autoIncrement" pgType:"number" pgText:"ID"  pgContainer:"main"`
    Name string `json:"name" gorm:"size:100"                 pgType:"text"   pgText:"Имя" pgContainer:"main" pgEdit:"true" pgPlaceholder:"Введите имя"`
}

// 2. Обязательный интерфейс IModel
func (u *UserModel) GetContainers() *[]inputs.Container {
    return &[]inputs.Container{
        {Key: "main", Direction: "vertical", Title: "Основное"},
    }
}

// 3. Опционально: создание
func (u *UserModel) Create(params *pg.QueryParams) error {
    // сохранить в БД
    return nil
}

func (u *UserModel) GetPageSettings() *pg.PageSettings {
    return &pg.PageSettings{FormId: "user-form", Version: "1.0", Title: "Пользователь"}
}

func (u *UserModel) GetActions() *[]inputs.FormAction { return nil }

func main() {
    r := gin.Default()

    err := pg.SetDefinitions(func() error {
        pm := &pg.PageModel{}
        pm.SetModel(&UserModel{})
        pg.AddPageModel("users", pm)
        return nil
    }, pg.PageSetting{
        Service:    "api",
        DateFormat: "DD.MM.YYYY",
        TimeFormat: "HH:mm",
        PageSize:   20,
    })
    if err != nil {
        log.Fatal(err)
    }
    pg.GetModelsRoutes(r)
    r.Run(":8080")
}
```

Подробная документация — [DOCUMENTATION.md](DOCUMENTATION.md)
