package domain

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Turn struct {
	Id        int64          `json:"id"`
	DentistId int64          `json:"dentist_id"`
	PatientId int64          `json:"patient_id"`
	DateTime  CustomDateTime `json:"date_time"`
	Details   string         `json:"details"`
}

//Para cambiar
// type Turn struct {
// 	Id          int64      `json:"id"`
// 	PatientDni   string      `json:"patient_id"`
// 	DentistLicense   string      `json:"dentist_id"`
// 	DateTime    CustomDate `json:"date_time"`
// 	Details string     `json:"details"`
// }

type CustomDateTime struct {
	time.Time
}

const dateTimeFormat = "2006-01-02 15:04:05"

func (cdt *CustomDateTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		return nil
	}

	t, err := time.Parse(`"`+dateTimeFormat+`"`, s)
	if err != nil {
		return err
	}

	cdt.Time = t
	return nil
}

func (t CustomDateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(`"` + t.Time.Format(dateTimeFormat) + `"`), nil
}

// ** Revisar nuevos metodos abajo **

func (cdt CustomDateTime) Value() (driver.Value, error) {
	if cdt.IsZero() {
		return nil, nil
	}

	return cdt.Time.Format(time.DateTime), nil
}

func (cdt *CustomDateTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		cdt.Time = v
	case nil:
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *CustomDateTime. It's possible the database driver doesn't know how to convert the type to 'CustomDate'", value)
	}
	return nil
}
