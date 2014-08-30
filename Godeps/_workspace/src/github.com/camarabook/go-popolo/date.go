package popolo

import "time"

const DateSpec = "2006-01-02"

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(d.Format(`"` + DateSpec + `"`)), nil
}

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return d.Time.MarshalJSON()
}
