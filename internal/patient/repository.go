package patient

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
	ErrNotFound         = errors.New("patient not found")
)

type repositorypatientssql struct {
	db *sql.DB
}

func NewPatientsSqlRepository(db *sql.DB) RepositoryPatients {
	return &repositorypatientssql{db: db}
}

// Create creates a new patient.
func (r *repositorypatientssql) Create(patient domain.Patient) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryInsertPatient)
	if err != nil {
		return domain.Patient{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Name,
		patient.LastName,
		patient.Address,
		patient.Dni,
		patient.DischargeDate,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Patient{}, ErrLastInsertedId
	}

	patient.Id = int64(lastId)

	return patient, nil

}

// GetAll returns all patients.
func (r *repositorypatientssql) GetAll() ([]domain.Patient, error) {
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
			&patient.LastName,
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
func (r *repositorypatientssql) GetByID(id int64) (domain.Patient, error) {
	row := r.db.QueryRow(QueryGetPatientById, id)

	var patient domain.Patient

	err := row.Scan(
		&patient.Id,
		&patient.Name,
		&patient.LastName,
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
func (r *repositorypatientssql) Update(patient domain.Patient, id int64) (domain.Patient, error) {
	// statement, err := r.db.Prepare(QueryUpdatePatient)
	// if err != nil {
	// 	return domain.Patient{}, err
	// }

	// defer statement.Close()

	result, err := r.db.Exec(QueryUpdatePatient,
		patient.Name,
		patient.LastName,
		patient.Address,
		patient.Dni,
		patient.DischargeDate,
		id,
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
func (r *repositorypatientssql) Delete(id int64) error {
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

// Patch updates a Patient by any field.
func (r *repositorypatientssql) Patch(
	patient domain.Patient,
	id int64) (domain.Patient, error) {

	var fieldsToUpdate []string
	var args []interface{}

	if patient.Name != "" {
		fieldsToUpdate = append(fieldsToUpdate, "name = ?")
		args = append(args, patient.Name)
	}
	if patient.LastName != "" {
		fieldsToUpdate = append(fieldsToUpdate, "last_name = ?")
		args = append(args, patient.LastName)
	}
	if patient.Address != "" {
		fieldsToUpdate = append(fieldsToUpdate, "address = ?")
		args = append(args, patient.Address)
	}
	if patient.Dni != "" {
		fieldsToUpdate = append(fieldsToUpdate, "dni = ?")
		args = append(args, patient.Dni)
	}
	// if time.Time.IsZero(patient.DischargeDate.Time) {
	// 	fieldsToUpdate = append(fieldsToUpdate, "discharge_date = ?")
	// 	args = append(args, patient.DischargeDate)
	// }

	if !patient.DischargeDate.Time.IsZero() {
		fieldsToUpdate = append(fieldsToUpdate, "discharge_date = ?")
		args = append(args, patient.DischargeDate)
	}

	if len(fieldsToUpdate) == 0 {
		return domain.Patient{}, ErrEmpty
	}

	queryString := "UPDATE patient SET " + strings.Join(fieldsToUpdate, ", ") + " WHERE patient_id = ?"
	args = append(args, id)

	// statement, err := r.db.Prepare(queryString)
	// if err != nil {
	// 	return domain.Patient{}, err
	// }
	// defer statement.Close()

	result, err := r.db.Exec(queryString, args...)
	if err != nil {
		return domain.Patient{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	return r.GetByID(int64(id))
}
