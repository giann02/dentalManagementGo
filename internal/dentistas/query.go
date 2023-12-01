package dentistas

var (
	QueryInsertDentista = `INSERT INTO dentista(apellido,nombre,matricula)
	VALUES(?,?,?)`
	QueryGetAllDentistas = `SELECT id, apellido, nombre, matricula 
	FROM dentista`
	QueryDeleteDentista  = `DELETE FROM dentista WHERE id = ?`
	QueryGetDentistaById = `SELECT id, apellido, nombre, matricula
	FROM dentista WHERE id = ?`
	QueryUpdateDentista = `UPDATE dentista SET apellido = ?, nombre = ?, matricula = ?
	WHERE id = ?`
)