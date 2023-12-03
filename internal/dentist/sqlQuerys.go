package dentist

var (
	QueryInsertDentists = `INSERT INTO dentist(name, last_name, license)
	VALUES(?,?,?)`
	QueryGetAllDentists = `SELECT *	FROM dentist`
	QueryDeleteDentist  = `DELETE FROM dentist WHERE dentist_id = ?`
	QueryGetDentistById = `SELECT *	FROM dentist WHERE dentist_id = ?`
	QueryUpdateDentist  = `UPDATE dentist SET name = ?, last_name = ?, license = ?
	WHERE dentist_id = ?` 
)
