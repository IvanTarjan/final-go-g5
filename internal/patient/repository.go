package patient

import (
	"database/sql"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type mySqlRepository struct {
	db *sql.DB
}

func NewmySqlRepository(db *sql.DB) RepositoryInterface {
    return &mySqlRepository{db}
}

func (r *mySqlRepository) Post(patient domain.Patient) (domain.Patient, error)

func (r *mySqlRepository) GetById(id int64) (domain.Patient, error)

func (r *mySqlRepository) Put(patient domain.Patient) (domain.Patient, error)

func (r *mySqlRepository) Patch(attributes map[string]string) (domain.Patient, error)

func (r *mySqlRepository) DeleteById(id int64) error