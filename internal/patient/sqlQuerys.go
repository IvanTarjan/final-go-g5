package patient

var (
	QueryInsertPatient = `INSERT INTO patients(name, last_name, address, dni, ingress_date)
	VALUES(?,?,?,?,?)`
	QueryGetAllPatient = `SELECT id, name, last_name, address, dni, ingress_date 
	FROM patients`
	QueryDeletePatient  = `DELETE FROM patients WHERE id = ?`
	QueryGetPatientById = `SELECT id, name, last_name, address, dni, ingress_date
	FROM patients WHERE id = ?`
	QueryUpdatePatient = `UPDATE patients SET name = ?, last_name = ?, address = ?, dni = ?, ingress_date = ?
	WHERE id = ?`
)
