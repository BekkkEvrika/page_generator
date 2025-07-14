package page_generator

import (
	"strings"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(JavaToGoTimeFormat(globalDateFormat), string(text))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func JavaToGoTimeFormat(javaFmt string) string {
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
