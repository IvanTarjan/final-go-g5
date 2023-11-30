package turn

var (
	QueryInsertTurn = `INSERT INTO turns(patient_id, dentist_id, date, details)
    VALUES(?,?,?,?)`
    QueryGetAllTurn  = `SELECT * FROM turn`
    QueryDeleteTurn  = `DELETE FROM turns WHERE turno_id =?`
    QueryGetTurnById = `SELECT turn_id, patient_id, dentist_id, date_time, description FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id WHERE turno_id =?`
    QueryUpdateTurn  = `UPDATE turns SET patient_id =?, dentist_id = ?, date_time =?, details =?
    WHERE turno_id =?`
	QueryGetByPatientDni = `SELECT turn_id, patient_id, dentist_id, date_time, description FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id where dni = ?`
	QueryGetByDentistLicence = `SELECT turn_id, patient_id, dentist_id, date_time, description FROM turn INNER JOIN patient ON turn.patient_id = patient.patient_id INNER JOIN dentist ON turn.dentist_id = dentist.dentist_id where licence = ?`
)