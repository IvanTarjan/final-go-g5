package turn

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
	ErrNotFound         = errors.New("turn not found")
)

type repositoryTurnssql struct {
	db *sql.DB
}

func NewTurnSqlRepository(db *sql.DB) RepositoryTurn {
	return &repositoryTurnssql{db: db}
}

// Create implements RepositoryTurn.
func (r *repositoryTurnssql) Create(turn domain.Turn) (domain.Turn, error) {
	result, err := r.db.Exec(QueryInsertTurn, turn.DentistId, turn.PatientId, turn.DateTime, turn.Details)
	if err != nil {
		return domain.Turn{}, ErrExecStatement
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Turn{}, ErrLastInsertedId
	}
	turn.Id = int64(lastId)
	return turn, nil
}

// GetAll implements RepositoryTurn.
func (r *repositoryTurnssql) GetAll() ([]domain.Turn, error) {
	rows, err := r.db.Query(QueryGetAllTurn)
	if err != nil {
		return []domain.Turn{}, err
	}
	defer rows.Close()

	var turns []domain.Turn

	for rows.Next() {
		var turn domain.Turn
		err := rows.Scan(
			&turn.Id,
			&turn.DentistId,
			&turn.PatientId,
			&turn.DateTime,
			&turn.Details,
		)

		if err != nil {
			return nil, err
		}
		turns = append(turns, turn)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return turns, nil
}

// GetByID implements RepositoryTurn.
func (r *repositoryTurnssql) GetByID(id int64) (domain.Turn, error) {
	row := r.db.QueryRow(QueryGetTurnById, id)

	var turn domain.Turn

	err := row.Scan(
		&turn.Id,
		&turn.DentistId,
		&turn.PatientId,
		&turn.DateTime,
		&turn.Details,
	)

	if err != nil {
		return domain.Turn{}, err
	}

	return turn, nil
}

func (r *repositoryTurnssql) GetByPatientDni(patientDni string) ([]domain.TurnDetails, error) {
	rows, err := r.db.Query(QueryGetByPatientDni, patientDni)
	if err != nil {
		return []domain.TurnDetails{}, err
	}

	defer rows.Close()

    var turns []domain.TurnDetails

	for rows.Next(){

		var turn domain.TurnDetails
		err = rows.Scan(
			&turn.Id,
			&turn.Dentist.Id,
			&turn.Dentist.Name,
			&turn.Dentist.LastName,
			&turn.Dentist.License,
			&turn.Patient.Id,
			&turn.Patient.Name,
			&turn.Patient.LastName,
			&turn.Patient.Address,
			&turn.Patient.Dni,
			&turn.Patient.DischargeDate,
			&turn.DateTime,
			&turn.Details,
		)
		if err != nil {
			return []domain.TurnDetails{}, err
		}
		turns = append(turns, turn)
	}
	return turns, nil
}

// Update implements RepositoryTurn.
func (r *repositoryTurnssql) Update(turn domain.Turn, id int64) (domain.Turn, error) {
	result, err := r.db.Exec(QueryUpdateTurn,
		turn.DentistId,
		turn.PatientId,
		turn.DateTime,
		turn.Details,
		id,
	)

	if err != nil {
		return domain.Turn{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Turn{}, err
	}

	turn.Id = int64(id)

	return turn, nil
}

// Patch implements RepositoryTurn.
func (r *repositoryTurnssql) Patch(turn domain.Turn, id int64) (domain.Turn, error) {
	var fieldsToUpdate []string
	var args []interface{}

	if turn.DentistId != 0 {
		fieldsToUpdate = append(fieldsToUpdate, "name = ?")
		args = append(args, turn.DentistId)
	}
	if turn.PatientId != 0 {
		fieldsToUpdate = append(fieldsToUpdate, "last_name = ?")
		args = append(args, turn.PatientId)
	}
	if !turn.DateTime.IsZero() {
		fieldsToUpdate = append(fieldsToUpdate, "address = ?")
		args = append(args, turn.DateTime)
	}
	if turn.Details != "" {
		fieldsToUpdate = append(fieldsToUpdate, "details = ?")
		args = append(args, turn.Details)
	}

	if len(fieldsToUpdate) == 0 {
		return domain.Turn{}, ErrEmpty
	}

	queryString := "UPDATE turn SET " + strings.Join(fieldsToUpdate, ", ") + " WHERE turn_id = ?"
	args = append(args, id)

	result, err := r.db.Exec(queryString, args...)
	if err != nil {
		return domain.Turn{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Turn{}, err
	}

	return r.GetByID(int64(id))
}

// Delete implements RepositoryTurn.
func (r *repositoryTurnssql) Delete(id int64) error {
	result, err := r.db.Exec(QueryDeleteTurn, id)
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
