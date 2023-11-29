package dentist

import (
	"log"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type ServiceDentists interface {
	Create(dentist domain.Dentist) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	GetByID(id int64) (domain.Dentist, error)
	Update(dentist domain.Dentist, id int64) (domain.Dentist, error)
	Delete(id int64) error
	Patch(dentist domain.Dentist, id int64) (domain.Dentist, error)
}

type service struct {
	repository RepositoryDentists
}

func NewServiceDentist(repository RepositoryDentists) *service {
	return &service{repository: repository}
}

// Create creates a new dentist.
func (s *service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.repository.Create(dentist)
	if err != nil {
		log.Println("[ServiceDentist][Create] error creating dentist", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// GetAll returns all dentists.
func (s *service) GetAll() ([]domain.Dentist, error) {
	listDentists, err := s.repository.GetAll()
	if err != nil {
		log.Println("[ServiceDentist][GetAll] error getting all dentist", err)
		return []domain.Dentist{}, err
	}

	return listDentists, nil
}

// GetByID returns a Dentist by ID.
func (s *service) GetByID(id int64) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(id)
	if err != nil {
		log.Println("[ServiceDentist][GetByID] error getting dentist with ID: ", id, "\n", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update updates a dentist by ID.
func (s *service) Update(dentist domain.Dentist, id int64) (domain.Dentist, error) {
	dentist, err := s.repository.Update(dentist, id)
	if err != nil {
		log.Println("[ServiceDentist][Update] error updating dentist with ID: ", id, "\n", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Delete deletes a dentist by ID.
func (s *service) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		log.Println("[ServiceDentist][Delete] error deleting dentist with ID: ", id, "\n", err)
		return err
	}

	return nil
}

// Patch updates a dentist by ID.
func (s *service) Patch(dentist domain.Dentist, id int64) (domain.Dentist, error) {
	dentist, err := s.repository.Patch(dentist, id)
	if err != nil {
		log.Println("[ServiceDentist][Patch] error patching dentist with ID: ", id, "\n", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}
