package patient

import (
	"log"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type ServicePatients interface {
	Create(patient domain.Patient) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	GetByID(id int64) (domain.Patient, error)
	Update(patient domain.Patient, id int64) (domain.Patient, error)
	Delete(id int64) error
	Patch(patient domain.Patient, id int64) (domain.Patient, error)
}

type service struct {
	repository RepositoryPatients
}

func NewServicePatient(repository RepositoryPatients) ServicePatients {
	return &service{repository: repository}
}

// Create creates a new patient.
func (s *service) Create(patient domain.Patient) (domain.Patient, error) {
	patient, err := s.repository.Create(patient)
	if err != nil {
		log.Println("[ServicePatients][Create] error creating patient", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// GetAll returns all patients.
func (s *service) GetAll() ([]domain.Patient, error) {
	listPatients, err := s.repository.GetAll()
	if err != nil {
		log.Println("[ServicePatients][GetAll] error getting all patients", err)
		return []domain.Patient{}, err
	}

	return listPatients, nil
}

// GetByID returns a Patient by ID.
func (s *service) GetByID(id int64) (domain.Patient, error) {
	patient, err := s.repository.GetByID(id)
	if err != nil {
		log.Println("[ServicePatients][GetByID] error getting patient with ID: ", id, "\n", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update updates a patient by ID.
func (s *service) Update(patient domain.Patient, id int64) (domain.Patient, error) {
	patient, err := s.repository.Update(patient, id)
	if err != nil {
		log.Println("[ServicePatients][Update] error updating patient with ID: ", id, "\n", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Delete deletes a patient by ID.
func (s *service) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		log.Println("[ServicePatients][Delete] error deleting patient with ID: ", id, "\n", err)
		return err
	}

	return nil
}

// Patch updates a patient by ID.
func (s *service) Patch(patient domain.Patient, id int64) (domain.Patient, error) {
	patient, err := s.repository.Patch(patient, id)
	if err != nil {
		log.Println("[ServicePatients][Patch] error patching patient with ID: ", id, "\n", err)
		return domain.Patient{}, err
	}

	return patient, nil
}
