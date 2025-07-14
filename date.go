package page_generator

import "time"

type Date time.Time

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(globalDateFormat, string(text))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}
