package dentists

var (
	QueryInsertDentists = `INSERT INTO dentists(name, last_name, code)
	VALUES(?,?,?)`
	QueryGetAllDentists = `SELECT id, name, last_name, code 
	FROM dentists`
	QueryDeleteDentist  = `DELETE FROM dentists WHERE id = ?`
	QueryGetDentistById = `SELECT id, name, last_name, code
	FROM dentists WHERE id = ?`
	QueryUpdateDentist = `UPDATE dentists SET name = ?, last_name = ?, code = ?
	WHERE id = ?`
)
