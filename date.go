package page_generator

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(javaToGoTimeFormat(globalDateFormat), string(text))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte(`null`), nil
	}
	return []byte(`"` + t.Format(javaToGoTimeFormat(globalDateFormat)) + `"`), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "null" || str == "" {
		*d = Date(time.Time{})
		return nil
	}
	t, err := time.Parse(javaToGoTimeFormat(globalDateFormat), str)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d).Format(javaToGoTimeFormat(globalDateFormat)), nil
}

func (d *Date) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("DateTime: cannot scan type %T into DateTime ", value)
	}
	*d = Date(t)
	return nil
}
