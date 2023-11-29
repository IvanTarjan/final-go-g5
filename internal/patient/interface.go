package patient

import (
	"context"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type RepositoryPatients interface {
	Create(ctx context.Context, patient domain.Patient) (domain.Patient, error)
	GetAll(ctx context.Context) ([]domain.Patient, error)
	GetByID(ctx context.Context, id int64) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int64) (domain.Patient, error)
	Delete(ctx context.Context, id int64) error
	Patch(ctx context.Context, attributes map[string]string, id int64) (domain.Patient, error)
}
