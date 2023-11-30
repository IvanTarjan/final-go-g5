package domain

import (
	"strings"
	"time"
)

type Turn struct {
	Id          int64      `json:"id"`
	PatientId   int64      `json:"patient_id"`
	DentistId   int64      `json:"dentist_id"`
	DateTime    CustomDate `json:"date_time"`
	Description string     `json:"description"`
}

//Para cambiar
// type Turn struct {
// 	Id          int64      `json:"id"`
// 	PatientDni   string      `json:"patient_id"`
// 	DentistLicense   string      `json:"dentist_id"`
// 	DateTime    CustomDate `json:"date_time"`
// 	Description string     `json:"description"`
// }

type CustomDateTime struct {
	time.Time
}

func (t *CustomDateTime) UnmarshalJSON(b []byte) (err error) {
	dateString := strings.Trim(string(b), `"`)
	date, err := time.Parse(time.DateTime, dateString)
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}

func (t *CustomDateTime) MarshalJSON() ([]byte, error) {
	dateString := time.Time.Format(t.Time, time.DateTime)
	return []byte(`"` + dateString + `"`), nil
}
