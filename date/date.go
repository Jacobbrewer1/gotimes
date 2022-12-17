package date

import (
	"github.com/Jacobbrewer1/gotimes/layouts"
	"time"
)

type Date time.Time

func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(layouts.DateLayout, string(text))
	if err != nil {
		return err
	}

	if d == nil {
		d = &Date{}
	}

	*d = Date(t)
	return nil
}

func (d *Date) Scan(src any) error {
	t, err := time.Parse(layouts.DateTimeWithNumericAndUtcZone, src.(string))
	if err != nil {
		return err
	}

	if d == nil {
		d = &Date{}
	}

	*d = Date(t)
	return nil
}

func (d *Date) String() string {
	return time.Time(*d).Format(layouts.DateLayout)
}

func (d *Date) TimeValue() *time.Time {
	if d == nil {
		return nil
	}

	var t = time.Time(*d)
	return &t
}

func (d *Date) Display() string {
	return d.TimeValue().Format(layouts.DateDisplay)
}
