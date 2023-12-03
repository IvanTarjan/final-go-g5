package patient

var (
	QueryInsertPatient = `INSERT INTO patient(name, last_name, address, dni, discharge_date)
	VALUES(?,?,?,?,?)`
	QueryGetAllPatient  = `SELECT * FROM patient`
	QueryDeletePatient  = `DELETE FROM patient WHERE patient_id = ?`
	QueryGetPatientById = `SELECT *	FROM patient WHERE patient_id = ?`
	QueryUpdatePatient  = `UPDATE patient SET name = ?, last_name = ?, address = ?, dni = ?, discharge_date = ?
	WHERE patient_id = ?`
)
