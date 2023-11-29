package patient

import (
	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type RepositoryPatients interface {
	Create(patient domain.Patient) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	GetByID(id int64) (domain.Patient, error)
	Update(patient domain.Patient, id int64) (domain.Patient, error)
	Delete(id int64) error
	Patch(patient domain.Patient, id int64) (domain.Patient, error)
}
