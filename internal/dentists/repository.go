package dentists

import (
	"context"
	"database/sql"
	"errors"
	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrEmpty            = errors.New("empty list")
	ErrNotFound         = errors.New("dentist not found")
)

type repositorydentistsql struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) RepositoryDentists {
	return &repositorydentistsql{db: db}
}

// Create
func (r *repositorydentistsql) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryInsertDentists)
	if err != nil {
		return domain.Dentist{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.Code,
	)

	if err != nil {
		return domain.Dentist{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Dentist{}, ErrLastInsertedId
	}

	dentist.Id = int(lastId)

	return dentist, nil

}

// GetAll
func (r *repositorydentistsql) GetAll(ctx context.Context) ([]domain.Dentist, error) {
	rows, err := r.db.Query(QueryGetAllDentists)
	if err != nil {
		return []domain.Dentist{}, err
	}

	defer rows.Close()

	var dentists []domain.Dentist

	for rows.Next() {
		var dentist domain.Dentist
		err := rows.Scan(
			&dentist.Id,
			&dentist.Name,
			&dentist.LastName,
			&dentist.Code,
		)
		if err != nil {
			return []domain.Dentist{}, err
		}

		dentists = append(dentists, dentist)
	}

	if err := rows.Err(); err != nil {
		return []domain.Dentist{}, err
	}

	return dentists, nil
}

// GetByID
func (r *repositorydentistsql) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	row := r.db.QueryRow(QueryGetDentistById, id)

	var dentist domain.Dentist
	err := row.Scan(
		&dentist.Id,
		&dentist.Name,
		&dentist.LastName,
		&dentist.Code,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update
func (r *repositorydentistsql) Update(
	ctx context.Context,
	dentist domain.Dentist,
	id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.Code,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	dentist.Id = id

	return dentist, nil

}

// Delete
func (r *repositorydentistsql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteDentist, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// Patch Updates a dentist by ID.
func (r *repositorydentistsql) Patch(
	ctx context.Context,
	dentist domain.Dentist,
	id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.Code,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}
