package turn

import "github.com/IvanTarjan/final-go-g5/internal/domain"

type ServiceTurn interface {
	Create(turn domain.Turn) (domain.Turn, error)
	GetAll() ([]domain.Turn, error)
	GetByID(id int64) (domain.Turn, error)
	Update(turn domain.Turn, id int64) (domain.Turn, error)
	Delete(id int64) error
	Patch(turn domain.Turn, id int64) (domain.Turn, error)
}

type service struct {
	repository RepositoryTurn
}

// Create implements ServiceTurn.
func (*service) Create(turn domain.Turn) (domain.Turn, error) {
	panic("unimplemented")
}

// Delete implements ServiceTurn.
func (*service) Delete(id int64) error {
	panic("unimplemented")
}

// GetAll implements ServiceTurn.
func (*service) GetAll() ([]domain.Turn, error) {
	panic("unimplemented")
}

// GetByID implements ServiceTurn.
func (*service) GetByID(id int64) (domain.Turn, error) {
	panic("unimplemented")
}

// Patch implements ServiceTurn.
func (*service) Patch(turn domain.Turn, id int64) (domain.Turn, error) {
	panic("unimplemented")
}

// Update implements ServiceTurn.
func (*service) Update(turn domain.Turn, id int64) (domain.Turn, error) {
	panic("unimplemented")
}

func NewServiceTurn(repository RepositoryTurn) ServiceTurn {
	return &service{repository: repository}
}
