package turnos

var (
	QueryGetAllTurnos = `SELECT id,idPaciente, idDentista, fecha_hora, descripcion
	FROM turno`
	QueryDeleteTurno  = `DELETE FROM turno WHERE id = ?`
	QueryGetTurnoById = `SELECT id,idPaciente, idDentista, fecha_hora, descripcion
	FROM turno WHERE id = ?`
	QueryUpdateTurno = `UPDATE turno SET idPaciente = ?, idDentista = ?, fecha_hora = ?, descripcion = ?
	WHERE id = ?`
	QueryGetTurnoByPacienteDni = `SELECT turno.id,idPaciente, idDentista, p.nombre,p.apellido, fecha_hora, descripcion FROM turno  
	INNER JOIN paciente p ON p.id = turno.idPaciente
	WHERE p.DNI = ?`
	QueryGetDentistaIDByMatricula = `SELECT id FROM dentista WHERE matricula = ?`
	QueryGetPacienteIDByDNI = `SELECT id FROM paciente WHERE DNI = ?`
	QueryInsertTurno = `INSERT INTO turno (idPaciente, idDentista, fecha_hora, descripcion) VALUES (?, ?, ?, ?)`
)