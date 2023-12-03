package turn

var (
	QueryInsertTurn = `INSERT INTO turn(dentist_id, patient_id, date_time, details)
    VALUES(?,?,?,?)`
	QueryGetAllTurn  = `SELECT * FROM turn`
	QueryDeleteTurn  = `DELETE FROM turn WHERE turn_id =?`
	QueryGetTurnById = `SELECT turn.turn_id, turn.dentist_id, turn.patient_id, turn.date_time, turn.details FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id WHERE turn_id =?`
	QueryUpdateTurn  = `UPDATE turn SET patient_id =?, dentist_id = ?, date_time =?, details =?
    WHERE turn_id =?`
	QueryGetByPatientDni     = `SELECT turn_id, turn.dentist_id, turn.patient_id, date_time, details FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id where dni = ?`
	QueryGetByDentistLicence = `SELECT turn_id, turn.dentist_id, turn.patient_id, date_time, details FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id where licence = ?`
)
