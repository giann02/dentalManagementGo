-- Crear la base de datos
CREATE DATABASE IF NOT EXISTS bddOdontologia;
USE bddOdontologia;

-- Crear la tabla Dentista
CREATE TABLE IF NOT EXISTS Dentista (
    id INT AUTO_INCREMENT PRIMARY KEY,
    apellido VARCHAR(255) NOT NULL,
    nombre VARCHAR(255) NOT NULL,
    matricula VARCHAR(20) NOT NULL
);

-- Crear la tabla Paciente
CREATE TABLE IF NOT EXISTS Paciente (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    apellido VARCHAR(255) NOT NULL,
    domicilio VARCHAR(255),
    DNI VARCHAR(15) NOT NULL,
    fecha_alta DATE
);

-- Crear la tabla Turno
CREATE TABLE IF NOT EXISTS Turno (
    id INT AUTO_INCREMENT PRIMARY KEY,
    idPaciente INT,
    idDentista INT,
    fecha_hora DATETIME NOT NULL,
    descripcion VARCHAR(255),
    FOREIGN KEY (idPaciente) REFERENCES Paciente(id),
    FOREIGN KEY (idDentista) REFERENCES Dentista(id)
);
