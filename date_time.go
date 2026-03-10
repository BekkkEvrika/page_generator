package page_generator

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

func (d *DateTime) UnmarshalText(text []byte) error {
	t, err := time.Parse(javaToGoTimeFormat(globalDateFormat+" "+globalTimeFormat), string(text))
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte(`null`), nil
	}
	return []byte(`"` + t.Format(javaToGoTimeFormat(globalDateFormat+" "+globalTimeFormat)) + `"`), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "null" || str == "" {
		*d = DateTime(time.Time{})
		return nil
	}
	t, err := time.Parse(javaToGoTimeFormat(globalDateFormat+" "+globalTimeFormat), str)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) Value() (driver.Value, error) {
	return time.Time(d).Format(javaToGoTimeFormat(globalDateFormat + " " + globalTimeFormat)), nil
}

func (d *DateTime) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("DateTime: cannot scan type %T into DateTime ", value)
	}
	*d = DateTime(t)
	return nil
}

func javaToGoTimeFormat(javaFmt string) string {
	replacer := strings.NewReplacer(
		"yyyy", "2006",
		"yy", "06",
		"MM", "01",
		"dd", "02",
		"HH", "15",
		"hh", "03",
		"mm", "04",
		"ss", "05",
		"SSS", ".000",
		"a", "PM",
		"Z", "-0700",
		"XXX", "-07:00",
		"XX", "-0700",
		"X", "-07",
	)
	return replacer.Replace(javaFmt)
}
