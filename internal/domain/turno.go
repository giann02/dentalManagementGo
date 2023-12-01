package domain

import "time"

// Turno es una estructura que define ...
type Turno struct {
	Id          int     	`json:"id"`
	IdPaciente  int    		`json:"idPaciente"`
	IdDentista  int    		`json:"idDentista"`
	Fecha_hora  time.Time  	`json:"fecha_hora"`
	Descripcion	string  	`json:"descripcion"`
}