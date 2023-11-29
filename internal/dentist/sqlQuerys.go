package dentist

var (
	QueryInsertDentists = `INSERT INTO dentists(name, last_name, license)
	VALUES(?,?,?)`
	QueryGetAllDentists = `SELECT *	FROM dentists`
	QueryDeleteDentist  = `DELETE FROM dentists WHERE id = ?`
	QueryGetDentistById = `SELECT *	FROM dentists WHERE id = ?`
	QueryUpdateDentist  = `UPDATE dentists SET name = ?, last_name = ?, license = ?
	WHERE id = ?`
)
