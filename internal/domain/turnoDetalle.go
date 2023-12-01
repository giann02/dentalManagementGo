package domain

import "time"

type TurnoDetalle struct {
	Id 					int 		`json:"id"`
	IdPaciente			int			`json:"idPaciente"`
	IdDentista    		int    		`json:"idDentista"`
	PacienteNombre     	string    	`json:"pacienteNombre"`
	PacienteApellido   	string    	`json:"pacienteApellido"`
	FechaHora          	time.Time 	 	`json:"fecha_hora"`
	Descripcion        	string    	`json:"descripcion"`
}