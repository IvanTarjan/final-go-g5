package dentists

import (
	"context"
	"log"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type ServiceDentists interface {
	Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error)
	GetAll(ctx context.Context) ([]domain.Dentist, error)
	GetByID(ctx context.Context, id int64) (domain.Dentist, error)
	Update(ctx context.Context, dentist domain.Dentist, id int64) (domain.Dentist, error)
	Delete(ctx context.Context, id int64) error
	Patch(ctx context.Context, dentist domain.Dentist, id int64) (domain.Dentist, error)
}

type service struct {
	repository RepositoryDentists
}

func NewServiceDentist(repository RepositoryDentists) *service {
	return &service{repository: repository}
}

// Create New dentist.
func (s *service) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[ServiceDentist][Create] error creating dentist", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// GetAll All dentist.
func (s *service) GetAll(ctx context.Context) ([]domain.Dentist, error) {
	listDentists, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[ServiceDentist][GetAll] error getting all dentist", err)
		return []domain.Dentist{}, err
	}

	return listDentists, nil
}

// GetByID Dentist by ID.
func (s *service) GetByID(ctx context.Context, id int64) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ServiceDentist][GetByID] error getting dentist by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update updates a dentist by ID.
func (s *service) Update(ctx context.Context, dentist domain.Dentist, id int64) (domain.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[ServiceDentist][Update] error updating dentist by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Delete Deletes a dentist by ID.
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[ServiceDentist][Delete] error deleting dentist by ID", err)
		return err
	}

	return nil
}

// validatePatch validates the fields to be updated.
func (s *service) validatePatch(dentistStore, dentist domain.Dentist) (domain.Dentist, error) {

	if dentist.Name != "" {
		dentistStore.Name = dentist.Name
	}

	if dentist.LastName != "" {
		dentistStore.Name = dentist.LastName
	}

	if dentist.License != "" {
		dentistStore.License = dentist.License
	}

	return dentistStore, nil

}
