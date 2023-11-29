package domain

type Dentist struct {
	Id       int64  `json:"Id"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	License  string `json:"License"`
}
