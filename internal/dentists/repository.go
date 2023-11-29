package dentists

import (
	"database/sql"
	"errors"
	"strings"

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

// Create creates a new dentist.
func (r *repositorydentistsql) Create(dentist domain.Dentist) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryInsertDentists)
	if err != nil {
		return domain.Dentist{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.License,
	)

	if err != nil {
		return domain.Dentist{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Dentist{}, ErrLastInsertedId
	}

	dentist.Id = int64(lastId)

	return dentist, nil

}

// GetAll returns all dentists.
func (r *repositorydentistsql) GetAll() ([]domain.Dentist, error) {
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
			&dentist.License,
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

// GetByID returns a dentist by ID.
func (r *repositorydentistsql) GetByID(id int64) (domain.Dentist, error) {
	row := r.db.QueryRow(QueryGetDentistById, id)

	var dentist domain.Dentist
	err := row.Scan(
		&dentist.Id,
		&dentist.Name,
		&dentist.LastName,
		&dentist.License,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update  updates a Dentist by ID.
func (r *repositorydentistsql) Update(dentist domain.Dentist, id int64) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.License,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	dentist.Id = int64(id)

	return dentist, nil

}

// Delete deletes a Dentist by ID.
func (r *repositorydentistsql) Delete(id int64) error {
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

// Patch updates a Dentist by any field.
func (r *repositorydentistsql) Patch(
	dentist domain.Dentist,
	id int64) (domain.Dentist, error) {

	var fieldsToUpdate []string
	var args []interface{}

	if dentist.Name != "" {
		fieldsToUpdate = append(fieldsToUpdate, "name = ?")
		args = append(args, dentist.Name)
	}
	if dentist.LastName != "" {
		fieldsToUpdate = append(fieldsToUpdate, "last_name = ?")
		args = append(args, dentist.LastName)
	}
	if dentist.License != "" {
		fieldsToUpdate = append(fieldsToUpdate, "license = ?")
		args = append(args, dentist.License)
	}

	if len(fieldsToUpdate) == 0 {
		return domain.Dentist{}, ErrEmpty
	}

	queryString := "UPDATE dentists SET " + strings.Join(fieldsToUpdate, ", ") + " WHERE id = ?"
	args = append(args, id)

	statement, err := r.db.Prepare(queryString)
	if err != nil {
		return domain.Dentist{}, err
	}
	defer statement.Close()

	_, err = statement.Exec(args...)
	if err != nil {
		return domain.Dentist{}, err
	}

	return r.GetByID(int64(id))
}
