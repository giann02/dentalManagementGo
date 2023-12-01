package domain

//import "time"

// Paciente es una estructura que define ...
type Paciente struct {
	Id          int       `json:"id"`
	Nombre    	string    `json:"nombre"`
	Apellido    string    `json:"apellido"`
	Domicilio   string    `json:"domicilio"`
	Dni 		string	  `json:"DNI"`
	Fecha_alta	string    `json:"fecha_alta"`
}