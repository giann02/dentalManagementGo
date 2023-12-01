package pacientes

var (
	QueryInsertPaciente = `INSERT INTO paciente(nombre,apellido,domicilio,DNI,fecha_alta)
	VALUES(?,?,?,?,?)`
	QueryGetAllPacientes = `SELECT id,nombre, apellido, domicilio, DNI, fecha_alta
	FROM paciente`
	QueryDeletePaciente  = `DELETE FROM paciente WHERE id = ?`
	QueryGetPacienteById = `SELECT id,nombre, apellido, domicilio, DNI, fecha_alta
	FROM paciente WHERE id = ?`
	QueryUpdatePaciente = `UPDATE paciente SET nombre = ?, apellido = ?, domicilio = ?, DNI = ?, fecha_alta = ?
	WHERE id = ?`
)