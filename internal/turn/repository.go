package turn

import (
	"database/sql"
	"errors"
	"time"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrEmpty            = errors.New("empty list")
	ErrNotFound         = errors.New("dentist not found")
)

type repositoryturnssql struct {
	db *sql.DB
}

// Create implements RepositoryTurn.
func (r *repositoryturnssql) Create(turn domain.Turn) (domain.Turn, error) {
	statement, err := r.db.Prepare(QueryInsertTurn)
	if err != nil {
		return domain.Turn{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(turn.PatientId, turn.DentistId, turn.DateTime.Format(time.DateTime), turn.Description)
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
func (r *repositoryturnssql) GetAll() ([]domain.Turn, error) {
	rows, err := r.db.Query(QueryGetAllTurn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var turns []domain.Turn

	for rows.Next() {
		var turn domain.Turn
		var dateTimeString string
		err := rows.Scan(
			&turn.Id,
			&turn.PatientId,
			&turn.DentistId,
			&dateTimeString,
			&turn.Description,
		)

		if err != nil {
			return nil, err
		}
		turn.DateTime.Time, err = time.Parse(time.DateTime, dateTimeString)
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

// Delete implements RepositoryTurn.
func (r *repositoryturnssql) Delete(id int64) error {
	panic("unimplemented")
}

// GetByID implements RepositoryTurn.
func (r *repositoryturnssql) GetByID(id int64) (domain.Turn, error) {
	panic("unimplemented")
}

// Patch implements RepositoryTurn.
func (r *repositoryturnssql) Patch(turn domain.Turn, id int64) (domain.Turn, error) {
	panic("unimplemented")
}

// Update implements RepositoryTurn.
func (r *repositoryturnssql) Update(turn domain.Turn, id int64) (domain.Turn, error) {
	panic("unimplemented")
}

func NewTurnSqlRepository(db *sql.DB) RepositoryTurn {
	return &repositoryturnssql{db: db}
}
