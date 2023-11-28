package domain

type Dentist struct {
	Id       int64   `json:"Id"`
	Name     string  `json:"Name"`
	LastName float64 `json:"LastName"`
	License  int     `json:"License"`
}
