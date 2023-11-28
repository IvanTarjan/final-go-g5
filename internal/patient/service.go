package patient

import (
	"context"
	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"log"
)

type ServicePatients interface {
	Create(ctx context.Context, patient domain.Patient) (domain.Patient, error)
	GetAll(ctx context.Context) ([]domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
}

type service struct {
	repository RepositoryPatients
}

func NewServicePatient(repository RepositoryPatients) ServicePatients {
	return &service{repository: repository}
}

// Create creates a new patient.
func (s *service) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	patient, err := s.repository.Create(ctx, patient)
	if err != nil {
		log.Println("[ServicePatients][Create] error creating patient", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// GetAll returns all patient.
func (s *service) GetAll(ctx context.Context) ([]domain.Patient, error) {
	listPatients, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[ServicePatients][GetAll] error getting all patients", err)
		return []domain.Patient{}, err
	}

	return listPatients, nil
}

// GetByID returns a Patient by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	patient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ServicePatients][GetByID] error getting patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update updates a patient by ID.
func (s *service) Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patient, err := s.repository.Update(ctx, patient, id)
	if err != nil {
		log.Println("[ServicePatients][Update] error updating patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Delete deletes a patient by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[ServicePatients][Delete] error deleting patient by ID", err)
		return err
	}

	return nil
}

// Patch updates a patient by ID.
func (s *service) Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patientStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ServicePatients][Patch] error getting patient by ID", err)
		return domain.Patient{}, err
	}

	patientPatch, err := s.validatePatch(patientStore, patient)
	if err != nil {
		log.Println("[ServicePatients][Patch] error validating patient", err)
		return domain.Patient{}, err
	}

	patient, err = s.repository.Patch(ctx, patientPatch, id)
	if err != nil {
		log.Println("[ServicePatients][Patch] error patching patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(patientStore, patient domain.Patient) (domain.Patient, error) {

	if patient.Name != "" {
		patientStore.Name = patient.Name
	}

	if patient.Surname != "" {
		patientStore.Name = patient.Surname
	}

	if patient.Address != "" {
		patientStore.Address = patient.Address
	}

	if patient.Dni != "" {
		patientStore.Dni = patient.Dni
	}

	if !patient.DischargeDate.IsZero() {
		patientStore.DischargeDate = patient.DischargeDate
	}

	return patientStore, nil
}
