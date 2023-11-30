package domain

import (
	"strings"
	"time"
)

type Patient struct {
	Id            int64      `json:"id"`
	Name          string     `json:"name"`
	Surname       string     `json:"surname"`
	Address       string     `json:"address"`
	Dni           string     `json:"dni"`
	DischargeDate CustomDate `json:"discharge_date"`
}

type CustomDate struct {
	time.Time
}

func (t *CustomDate) UnmarshalJSON(b []byte) (err error) {
	dateString := strings.Trim(string(b), `"`)
	date, err := time.Parse(time.DateOnly, dateString)
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
