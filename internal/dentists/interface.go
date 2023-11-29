package dentists

import (
	"context"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type RepositoryDentists interface {
	Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error)
	GetAll(ctx context.Context) ([]domain.Dentist, error)
	GetByID(ctx context.Context, id int64) (domain.Dentist, error)
	Update(ctx context.Context, dentist domain.Dentist, id int64) (domain.Dentist, error)
	Delete(ctx context.Context, id int64) error
	Patch(ctx context.Context, attributes map[string]string, id int64) (domain.Dentist, error)
}
