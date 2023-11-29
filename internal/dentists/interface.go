package dentists

import (
	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type RepositoryDentists interface {
	Create(dentist domain.Dentist) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	GetByID(id int64) (domain.Dentist, error)
	Update(dentist domain.Dentist, id int64) (domain.Dentist, error)
	Delete(id int64) error
	Patch(dentist domain.Dentist, id int64) (domain.Dentist, error)
}
