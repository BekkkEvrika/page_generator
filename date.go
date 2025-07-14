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

func (d Date) Value() (driver.Value, error) {
	return time.Time(d).Format(javaToGoTimeFormat(globalDateFormat)), nil
}

func (d *Date) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("Date: cannot scan type %T into Date ", value)
	}
	*d = Date(t)
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
