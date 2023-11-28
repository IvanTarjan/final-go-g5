package dentists

import "github.com/IvanTarjan/final-go-g5/internal/domain"

type Service struct {
	repository Repository
}

func NewServiceDentist(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Post(domain.Dentist) (domain.Dentist, error) {
	panic("ª")
}

func (s *Service) GetByID(id int64) (domain.Dentist, error) {
	panic("ª")
}

func (s *Service) Put(dentist domain.Dentist, id int64) (domain.Dentist, error) {
	panic("ª")
}

func (s *Service) Patch([]string) (domain.Dentist, error) {
	panic("ª")
}

func (s *Service) Delete(id int64) error {
	panic("ª")
}
