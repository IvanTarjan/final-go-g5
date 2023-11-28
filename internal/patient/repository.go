package patient

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
	ErrNotFound         = errors.New("patient not found")
)

type repositorypatientssql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewPatientsSqlRepository(db *sql.DB) RepositoryPatients {
	return &repositorypatientssql{db: db}
}

// Create creates a new patient.
func (r *repositorypatientssql) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryInsertPatient)
	if err != nil {
		return domain.Patient{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Name,
		patient.Surname,
		patient.Address,
		patient.Dni,
		patient.DischargeDate,
	)

	if err != nil {
		return domain.Patient{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Patient{}, ErrLastInsertedId
	}

	patient.Id = int64(lastId)

	return patient, nil

}

// GetAll returns all patient.
func (r *repositorypatientssql) GetAll(ctx context.Context) ([]domain.Patient, error) {
	rows, err := r.db.Query(QueryGetAllPatient)
	if err != nil {
		return []domain.Patient{}, err
	}

	defer rows.Close()

	var patients []domain.Patient

	for rows.Next() {
		var patient domain.Patient
		err := rows.Scan(
			&patient.Id,
			&patient.Name,
			&patient.Surname,
			&patient.Address,
			&patient.Dni,
			&patient.DischargeDate,
		)
		if err != nil {
			return []domain.Patient{}, err
		}

		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return []domain.Patient{}, err
	}

	return patients, nil
}

// GetByID returns a patient by ID.
func (r *repositorypatientssql) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	row := r.db.QueryRow(QueryGetPatientById, id)

	var patient domain.Patient

	err := row.Scan(
		&patient.Id,
		&patient.Name,
		&patient.Surname,
		&patient.Address,
		&patient.Dni,
		&patient.DischargeDate,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update updates a Patient by ID.
func (r *repositorypatientssql) Update(
	ctx context.Context,
	patient domain.Patient,
	id int) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryUpdatePatient)
	if err != nil {
		return domain.Patient{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Name,
		patient.Surname,
		patient.Address,
		patient.Dni,
		patient.DischargeDate,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	patient.Id = int64(id)

	return patient, nil

}

// Delete deletes a Patient by ID.
func (r *repositorypatientssql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeletePatient, id)
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

// Patch is a method that updates a patient by ID.
func (r *repositorypatientssql) Patch(
	ctx context.Context,
	patient domain.Patient,
	id int) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryUpdatePatient)
	if err != nil {
		return domain.Patient{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Name,
		patient.Surname,
		patient.Address,
		patient.Dni,
		patient.DischargeDate,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}
