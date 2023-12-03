package turn

var (
	QueryInsertTurn = `INSERT INTO turn(patient_id, dentist_id, date, details)
    VALUES(?,?,?,?)`
	QueryGetAllTurn  = `SELECT * FROM turn`
	QueryDeleteTurn  = `DELETE FROM turn WHERE turn_id =?`
	QueryGetTurnById = `SELECT turn_id, patient_id, dentist_id, date_time, description FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id WHERE turn_id =?`
	QueryUpdateTurn  = `UPDATE turn SET patient_id =?, dentist_id = ?, date_time =?, details =?
    WHERE turn_id =?`
	QueryGetByPatientDni     = `SELECT turn_id, patient_id, dentist_id, date_time, description FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id where dni = ?`
	QueryGetByDentistLicence = `SELECT turn_id, patient_id, dentist_id, date_time, description FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id where licence = ?`
)
