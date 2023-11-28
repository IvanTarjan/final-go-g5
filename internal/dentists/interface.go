package dentists

import "github.com/IvanTarjan/final-go-g5/internal/domain"

type Repository interface {
	Post(dentist domain.Dentist) (domain.Dentist, error)
	GetByID(id int64) (domain.Dentist, error)
	Put(dentist domain.Dentist, id int64) (domain.Dentist, error)
	Patch(id int64, properties map[string]interface{}) (domain.Dentist, error)
	Delete(id int64) error
}
