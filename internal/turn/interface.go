package turn

import "github.com/IvanTarjan/final-go-g5/internal/domain"

type RepositoryTurn interface {
	Create(turn domain.Turn) (domain.Turn, error)
	GetAll() ([]domain.Turn, error)
	GetByID(id int64) (domain.Turn, error)
	GetByPatientDni(id string) ([]domain.TurnDetails, error)
	Update(turn domain.Turn, id int64) (domain.Turn, error)
	Delete(id int64) error
	Patch(turn domain.Turn, id int64) (domain.Turn, error)
}
