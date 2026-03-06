# 📖 Примеры использования Page Generator

---

## 🚀 Быстрый старт

### 1. Базовый пример структуры

```go
package models

import "github.com/BekkkEvrika/page_generator/inputs"

type Product struct {
    ID    int    `pg:"-" gorm:"primaryKey" json:"id"`
    Name  string `pgType:"text-field" pgText:"Название" pgContainer:"main" json:"name"`
    Price float64 `pgType:"number-view" pgText:"Цена" pgContainer:"main" json:"price"`
}

// Минимальная реализация IModel
func (p *Product) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "main", Title: "Основная информация", Inputs: []inputs.Input{}},
    }
}
```

---

## 💼 Реальный пример: Система управления сотрудниками

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

## 🎓 Пример: Система управления курсами

```go
package models

import (
    "time"
    pg "github.com/BekkkEvrika/page_generator"
    "github.com/BekkkEvrika/page_generator/inputs"
)

type Course struct {
    ID          int       `pg:"-" gorm:"primaryKey" json:"id"`
    Title       string    `pgType:"text-field" pgText:"Название курса" pgContainer:"info" json:"title"`
    Description string    `pgType:"text-view" pgText:"Описание" pgContainer:"info" json:"description"`
    Category    string    `pgType:"combo-box" pgText:"Категория" pgContainer:"info" json:"category"`
    Instructor  string    `pgType:"combo-box" pgText:"Инструктор" pgContainer:"info" json:"instructor"`
    Level       string    `pgType:"combo-box" pgText:"Уровень" pgContainer:"details" json:"level"`
    Duration    int       `pgType:"number-view" pgText:"Длительность (часов)" pgContainer:"details" json:"duration" pgMin:"1" pgMax:"1000"`
    Price       float64   `pgType:"number-view" pgText:"Цена" pgContainer:"details" json:"price" pgMin:"0"`
    MaxStudents int       `pgType:"number-view" pgText:"Макс. студентов" pgContainer:"details" json:"max_students" pgMin:"1"`
    StartDate   time.Time `pgType:"date-time" pgText:"Дата начала" pgContainer:"schedule" json:"start_date"`
    EndDate     time.Time `pgType:"date-time" pgText:"Дата окончания" pgContainer:"schedule" json:"end_date"`
    IsActive    bool      `pgType:"check-box" pgText:"Активный" pgContainer:"status" json:"is_active"`
    IsFeatured  bool      `pgType:"check-box" pgText:"Рекомендуемый" pgContainer:"status" json:"is_featured"`
    Syllabus    string    `pgType:"file-uploader" pgText:"Программа курса" pgFileSource:"/api/upload" pgContainer:"details" json:"syllabus"`
}

func (c *Course) GetContainers() []inputs.Container {
    return []inputs.Container{
        {Key: "info", Title: "Основная информация", Inputs: []inputs.Input{}},
        {Key: "details", Title: "Детали курса", Inputs: []inputs.Input{}},
        {Key: "schedule", Title: "Расписание", Inputs: []inputs.Input{}},
        {Key: "status", Title: "Статус", Inputs: []inputs.Input{}},
    }
}

func (c *Course) GetComboItems(params *pg.QueryParams, md map[string]interface{}) map[string]inputs.ComboItems {
    return map[string]inputs.ComboItems{
        "category": {  // соответствует json:"category"
            {ID: "programming", Text: "Программирование"},
            {ID: "design", Text: "Дизайн"},
            {ID: "business", Text: "Бизнес"},
            {ID: "marketing", Text: "Маркетинг"},
        },
        "level": {     // соответствует json:"level"
            {ID: "beginner", Text: "Начинающий"},
            {ID: "intermediate", Text: "Средний"},
            {ID: "advanced", Text: "Продвинутый"},
            {ID: "expert", Text: "Эксперт"},
        },
    }
}

// Остальные методы интерфейсов...
func (c *Course) Create(params *pg.QueryParams) error { return nil }
func (c *Course) Update(params *pg.QueryParams) error { return nil }
func (c *Course) Delete(params *pg.QueryParams) error { return nil }
func (c *Course) GetDefault(params *pg.QueryParams, md map[string]interface{}) map[string]string {
    return map[string]string{"is_active": "true"}  // соответствует json:"is_active"
}
func (c *Course) GetFileExtensions() map[string][]string {
    return map[string][]string{"syllabus": {"pdf", "doc", "docx"}}  // соответствует json:"syllabus"
}
func (c *Course) GetList(params *pg.QueryParams) error { return nil }
func (c *Course) Filter(obj interface{}, params *pg.QueryParams) error { return nil }
func (c *Course) GetCount(params *pg.QueryParams) (int, error) { return 0, nil }
func (c *Course) GetCompleteNodes() map[string][]string { return map[string][]string{} }
func (c *Course) GetContextActions() []inputs.Action { return []inputs.Action{} }
func (c *Course) GetIndexes() []inputs.Index { return []inputs.Index{} }
func (c *Course) GetExports() inputs.Export { return inputs.Export{Word: true, Excel: true} }
func (c *Course) GetQueryParams() map[string]string { return map[string]string{} }
func (c *Course) GetMetaData() map[string]pg.MetaData { return map[string]pg.MetaData{} }
func (c *Course) GetClearNodes() map[string][]string { return map[string][]string{} }
func (c *Course) GetEditPage() inputs.LoadAction {
    return inputs.LoadAction{Source: "/api/courses/edit", Action: "dialog", Text: "Редактировать"}
}
```

---

## 🏭 Пример: Вложенные контейнеры

```go
type ComplexForm struct {
    // поля...
}

func (cf *ComplexForm) GetContainers() []inputs.Container {
    return []inputs.Container{
        {
            Key: "main",
            Title: "Основное",
            Childs: []inputs.Container{
                {
                    Key:   "basic",
                    Title: "Базовые данные",
                    Inputs: []inputs.Input{},
                },
                {
                    Key:   "extended",
                    Title: "Расширенные данные",
                    Inputs: []inputs.Input{},
                },
            },
        },
    }
}

// Использование вложенных контейнеров в тегах:
// pgContainer:"basic" - будет помещено в вложенный контейнер
```

---

## 🔍 Примеры использования различных типов полей

### Текстовые поля
```go
type TextExample struct {
    ID          int    `pg:"-" gorm:"primaryKey"`
    SingleLine   string `pgType:"text-field" pgText:"Однострочное" pgMax:"100"`
    MultiLine    string `pgType:"text-view" pgText:"Многострочное"`
    Hidden       string `pgType:"hidden"`
    ReadOnly     string `pgType:"label" pgText:"Только чтение"`
}
```

### Числовые поля
```go
type NumberExample struct {
    ID          int     `pg:"-" gorm:"primaryKey"`
    Age        int     `pgType:"number-view" pgText:"Возраст" pgMin:"0" pgMax:"150"`
    Salary     float64 `pgType:"number-view" pgText:"Зарплата" pgMin:"0"`
    Quantity   int     `pgType:"number-view" pgText:"Количество" pgMin:"1"`
}
```

### Выпадающие меню и выбор
```go
type SelectExample struct {
    ID          int    `pg:"-" gorm:"primaryKey"`
    Status     string `pgType:"combo-box" pgText:"Статус"`
    Category   string `pgType:"combo-box" pgText:"Категория"`
    IsActive   bool   `pgType:"check-box" pgText:"Активен"`
}
```

### Поиск и автодополнение
```go
type SearchExample struct {
    ID          int    `pg:"-" gorm:"primaryKey"`
    City       string `pgType:"auto-complete" pgText:"Город"`
    Department string `pgType:"search-view" pgText:"Отдел" pgSearch:"/api/departments"`
    Employee   string `pgType:"search-view" pgText:"Сотрудник" pgSearch:"/api/employees"`
}
```

### Загрузка файлов
```go
type FileExample struct {
    ID          int    `pg:"-" gorm:"primaryKey"`
    Avatar      string `pgType:"file-uploader" pgText:"Аватар" pgFileSource:"/api/upload" pgFileMaxSize:"5242880"`
    Document    string `pgType:"file-uploader" pgText:"Документ" pgFileSource:"/api/upload" pgFileMaxSize:"10485760"`
    Certificate string `pgType:"file-uploader" pgText:"Сертификат" pgFileSource:"/api/upload"`
}
```

---

## 🧪 Тестирование

```go
package models

import "testing"

func TestEmployeeContainers(t *testing.T) {
    employee := &Employee{}
    containers := employee.GetContainers()
    
    if len(containers) != 4 {
        t.Errorf("Expected 4 containers, got %d", len(containers))
    }
    
    if containers[0].Key != "personal" {
        t.Errorf("Expected first container key to be 'personal', got '%s'", containers[0].Key)
    }
}

func TestEmployeeComboItems(t *testing.T) {
    employee := &Employee{}
    items := employee.GetComboItems(nil, nil)
    
    if items["department"] == nil {  // соответствует json:"department"
        t.Error("Expected department combo items")
    }
    
    if len(items["department"]) == 0 {
        t.Error("Expected non-empty department items")
    }
}
```

---

## 💾 Сохранение и загрузка данных

```go
// Загрузка данных в форму
func LoadEmployeeForm(id int) (*Employee, error) {
    employee := &Employee{}
    // db.First(employee, id)
    return employee, nil
}

// Сохранение данных из формы
func SaveEmployeeForm(employee *Employee) error {
    if err := employee.Create(nil); err != nil {
        return err
    }
    return nil
}

// Обновление данных
func UpdateEmployeeForm(id int, employee *Employee) error {
    // db.Model(employee).Where("id = ?", id).Updates(employee)
    return employee.Update(nil)
}
```

---

## 🚀 Интеграция с Gin

```go
package routes

import (
    "github.com/gin-gonic/gin"
    pg "github.com/BekkkEvrika/page_generator"
    "yourapp/models"
)

func SetupRoutes(router *gin.Engine) {
    // Инициализация Page Generator
    err := pg.SetDefinitions(func() error {
        // Регистрация моделей
        pageModel := &pg.PageModel{}
        pageModel.SetModel(&models.Employee{}, 2) // 2 колонки в форме
        pageModel.SetTableModel(&models.Employee{})
        pageModel.SetListModel(&models.Employee{})
        pageModel.SetFilterModel(&models.Employee{}, 2)
        
        return nil
    }, pg.PageSetting{
        Service: "my-service",
        DateFormat: "2006-01-02",
        PageSize: 20,
    })
    
    if err != nil {
        panic(err)
    }
    
    // Получение маршрутов
    pg.GetModelsRoutes(router)
}
```

---

## 📝 Лицензия

MIT License

