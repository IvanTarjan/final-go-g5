package turn

import (
	"log"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type ServiceTurn interface {
	Create(turn domain.Turn) (domain.Turn, error)
	GetAll() ([]domain.Turn, error)
	GetByID(id int64) (domain.Turn, error)
	GetByPatientDni(patientDni string) (domain.Turn, error)
	Update(turn domain.Turn, id int64) (domain.Turn, error)
	Delete(id int64) error
	Patch(turn domain.Turn, id int64) (domain.Turn, error)
}

type service struct {
	repository RepositoryTurn
}

func NewServiceTurn(repository RepositoryTurn) ServiceTurn {
	return &service{repository: repository}
}

// Create implements ServiceTurn.
func (s *service) Create(turn domain.Turn) (domain.Turn, error) {
	turn, err := s.repository.Create(turn)
	if err != nil {
		log.Println("[ServiceTurn][Create] error creating turn", err)
		return domain.Turn{}, err
	}

	return turn, nil
}

// GetAll implements ServiceTurn.
func (s *service) GetAll() ([]domain.Turn, error) {
	listTurns, err := s.repository.GetAll()
	if err != nil {
		log.Println("[ServiceTurns][GetAll] error getting all turns", err)
		return []domain.Turn{}, err
	}

	return listTurns, nil
}

// GetByID implements ServiceTurn.
func (s *service) GetByID(id int64) (domain.Turn, error) {
	turn, err := s.repository.GetByID(id)
	if err != nil {
		log.Println("[ServiceTurns][GetByID] error getting turn with ID: ", id, "\n", err)
		return domain.Turn{}, err
	}

	return turn, nil
}

func (s *service) GetByPatientDni(patientDni string) (domain.Turn, error) {
	turn, err := s.repository.GetByPatientDni(patientDni)
	if err != nil {
		log.Println("[ServiceTurns][GetByPatientDNI] error getting turn with PATIENT DNI: ", patientDni, "\n", err)
		return domain.Turn{}, err
	}

	return turn, nil
}

// Patch implements ServiceTurn.
func (s *service) Patch(turn domain.Turn, id int64) (domain.Turn, error) {
	turn, err := s.repository.Patch(turn, id)
	if err != nil {
		log.Println("[ServiceTurns][Patch] error patching turn with ID: ", id, "\n", err)
		return domain.Turn{}, err
	}

	return turn, nil
}

// Update implements ServiceTurn.
func (s *service) Update(turn domain.Turn, id int64) (domain.Turn, error) {
	turn, err := s.repository.Update(turn, id)
	if err != nil {
		log.Println("[ServiceTurns][Update] error updating turn with ID: ", id, "\n", err)
		return domain.Turn{}, err
	}

	return turn, nil
}

// Delete implements ServiceTurn.
func (s *service) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		log.Println("[ServiceTurns][Delete] error deleting turn with ID: ", id, "\n", err)
		return err
	}

	return nil
}
