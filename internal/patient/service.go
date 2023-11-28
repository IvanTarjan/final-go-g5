package patient

import "github.com/IvanTarjan/final-go-g5/internal/domain"

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) *Service {
	return &Service{repository}
}

func (s *Service) Post(patient domain.Patient) (domain.Patient, error)

func (s *Service) GetById(id int64) (domain.Patient, error)

func (s *Service) Put(patient domain.Patient) (domain.Patient, error)

func (s *Service) Patch(attributes map[string]string) (domain.Patient, error)

func (s *Service) DeleteById(id int64) error