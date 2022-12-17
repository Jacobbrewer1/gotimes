package datetime

import (
	"github.com/Jacobbrewer1/gotimes/layouts"
	"time"
)

type DateTime time.Time

func (d *DateTime) MarshalText() ([]byte, error) {
	var utc = d.UTC()
	return []byte(utc.String()), nil
}

func (d *DateTime) UnmarshalText(text []byte) error {
	loc, err := time.LoadLocation(layouts.UtcLocation)
	if err != nil {
		return err
	}

	t, err := time.ParseInLocation(layouts.DateTimeWithNumericAndUtcZone, string(text), loc)
	if err != nil {
		return err
	}

	t = t.In(time.Now().Location())

	if d == nil {
		d = &DateTime{}
	}

	*d = DateTime(t)
	return nil
}

func (d *DateTime) Scan(src any) error {
	loc, err := time.LoadLocation(layouts.UtcLocation)
	if err != nil {
		return err
	}

	t, err := time.ParseInLocation(layouts.DateTimeWithNumericAndUtcZone, src.(string), loc)
	if err != nil {
		return err
	}

	t = t.In(time.Now().Location())

	if d == nil {
		d = &DateTime{}
	}

	*d = DateTime(t)
	return nil
}

func (d *DateTime) String() string {
	return time.Time(*d).Format(layouts.DateTimeMarshalLayoutWithZone)
}

func (d *DateTime) TimeValue() *time.Time {
	if d == nil {
		return nil
	}

	var t = time.Time(*d)
	return &t
}

func (d *DateTime) Display() string {
	return d.TimeValue().Format(layouts.DateTimeDisplay)
}

func (d *DateTime) UTC() *DateTime {
	if d == nil {
		return nil
	}

	var t = DateTime(d.TimeValue().UTC())
	return &t
}
