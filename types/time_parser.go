package types

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

// Ensure CustomTime implements the driver.Valuer and sql.Scanner interfaces
var _ driver.Valuer = (*CustomTime)(nil)
var _ sql.Scanner = (*CustomTime)(nil)

const CustomTimeFormat = "2006-01-02T15:04:05.999999-0700"

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // Remove quotes
	t, err := time.Parse(CustomTimeFormat, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// Value implements the driver.Valuer interface.
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

// Scan implements the sql.Scanner interface.
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}

	ct.Time, _ = value.(time.Time)
	return nil
}

func (ct CustomTime) String() string {
	return ct.Time.Format(CustomTimeFormat)
}
