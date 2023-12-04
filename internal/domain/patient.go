package domain

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Patient struct {
	Id            int64      `json:"id"`
	Name          string     `json:"name"`
	LastName      string     `json:"last_name"`
	Address       string     `json:"address"`
	Dni           string     `json:"dni"`
	DischargeDate CustomDate `json:"discharge_date"`
}

type CustomDate struct {
	time.Time
}

func (t *CustomDate) UnmarshalJSON(b []byte) error {
	dateString := string(b)
	date, err := time.Parse(`"`+time.DateOnly+`"`, dateString)
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}

func (t *CustomDate) MarshalJSON() ([]byte, error) {
	dateString := time.Time.Format(t.Time, time.DateOnly)
	return []byte(`"` + dateString + `"`), nil
}

//** Revisar nuevos metodos abajo **

func (c CustomDate) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.Time.Format(time.DateOnly), nil
}

func (c *CustomDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		c.Time = v
	case nil:
		return nil
	default:
		return fmt.Errorf("Scan was not supported, storing driver.Value type %T into type *CustomDate. It's possible the database driver doesn't know how to convert the type to 'CustomDate'", value)
	}
	return nil
}
