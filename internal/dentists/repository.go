package dentists

import (
	"database/sql"
	"errors"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Post(dentist domain.Dentist) (domain.Dentist, error) {
	statement, err := r.db.Prepare(``)
	if err != nil {
		return domain.Dentist{}, errors.New("Error Prepare Statement")
	}
	result, err := statement.Exec()
	if err != nil {
		return domain.Dentist{}, errors.New("Error Exec Statement")
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Dentist{}, errors.New("Error Last Inserted Id")
	}
	result.RowsAffected()
	dentist.Id = lastId

	return dentist, nil
}

func (r *repository) GetByID(id int64) (domain.Dentist, error) {
	panic("ª")
}

func (r *repository) Put(dentist domain.Dentist, id int64) (domain.Dentist, error) {
	panic("ª")
}

func (r *repository) Patch(id int64, properties map[string]interface{}) (domain.Dentist, error) {
	panic("ª")
}

func (r *repository) Delete(id int64) error {
	panic("ª")
}
