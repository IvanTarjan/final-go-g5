package patient

var (
	QueryInsertPatient = `INSERT INTO patients(name, last_name, address, dni, discharge_date)
	VALUES(?,?,?,?,?)`
	QueryGetAllPatient  = `SELECT * FROM patients`
	QueryDeletePatient  = `DELETE FROM patients WHERE id = ?`
	QueryGetPatientById = `SELECT *	FROM patients WHERE id = ?`
	QueryUpdatePatient  = `UPDATE patients SET name = ?, last_name = ?, address = ?, dni = ?, discharge_date = ?
	WHERE id = ?`
)
