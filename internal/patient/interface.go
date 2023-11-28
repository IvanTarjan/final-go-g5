package patient

import "github.com/IvanTarjan/final-go-g5/internal/domain"

type RepositoryInterface interface {
	Post(domain.Patient) (domain.Patient, error)
	GetById(id int64) (domain.Patient, error)
	Put(domain.Patient) (domain.Patient, error)
	Patch(map[string]string) (domain.Patient, error)
	DeleteById(id int64) error
}