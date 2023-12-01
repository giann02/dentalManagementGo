package turnos

import (
	"context"
	"errors"
	//"log"
	"database/sql"
	"github.com/giann02/finalBackEndGo/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrNotFound  		= errors.New("error not found")
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Turno, error)
	GetByID(ctx context.Context, id int) (domain.Turno, error)
	Update(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
	GetTurnoByDNIDetalle(ctx context.Context, dniPaciente string) (domain.TurnoDetalle, error)
	GetPacienteIDByDNI(ctx context.Context, dniPaciente string) (int, error)
	GetDentistaIDByMatricula(ctx context.Context, matriculaDentista string) (int, error)
	CreateTurnoByDNIMatricula(ctx context.Context, turno domain.Turno, dniPaciente string, matriculaDentista string) (int, error)
}


type repository struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repository{db: db}
}



// GetAll is a method that returns all turnos.
func (r *repository) GetAll(ctx context.Context) ([]domain.Turno, error) {
	rows, err := r.db.Query(QueryGetAllTurnos)
	if err != nil {
		return []domain.Turno{}, err
	}

	defer rows.Close()

	var turnos []domain.Turno

	for rows.Next() {
		var turno domain.Turno
		err := rows.Scan(
			&turno.Id,
			&turno.IdPaciente,
			&turno.IdDentista,
			&turno.Fecha_hora,
			&turno.Descripcion,
		)
		if err != nil {
			return []domain.Turno{}, err
		}

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return []domain.Turno{}, err
	}

	return turnos, nil
}

// GetByID is a method that returns a turno by ID.
func (r *repository) GetByID(ctx context.Context, id int) (domain.Turno, error) {
	row := r.db.QueryRow(QueryGetTurnoById, id)

	var turno domain.Turno
	err := row.Scan(
		&turno.Id,
		&turno.IdPaciente,
		&turno.IdDentista,
		&turno.Fecha_hora,
		&turno.Descripcion,
	)

	if err != nil {
		return domain.Turno{}, err
	}

	return turno, nil
}

// Update is a method that updates a turno by ID.
func (r *repository) Update(
	ctx context.Context,
	turno domain.Turno,
	id int) (domain.Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateTurno)
	if err != nil {
		return domain.Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.IdPaciente,
		turno.IdDentista,
		turno.Fecha_hora,
		turno.Descripcion,
		turno.Id,
	)

	if err != nil {
		return domain.Turno{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Turno{}, err
	}

	turno.Id = id

	return turno, nil

}

// Delete is a method that deletes a turno by ID.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteTurno, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// Patch is a method that updates a turno by ID.
func (r *repository) Patch(
	ctx context.Context,
	turno domain.Turno,
	id int) (domain.Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateTurno)
	if err != nil {
		return domain.Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.IdPaciente,
		turno.IdDentista,
		turno.Fecha_hora,
		turno.Descripcion,
		turno.Id,
	)

	if err != nil {
		return domain.Turno{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Turno{}, err
	}

	return turno, nil
}

// GetTurnoByDNIDetalle obtiene el detalle de un turno por el DNI del paciente.
func (r *repository) GetTurnoByDNIDetalle(ctx context.Context, dniPaciente string) (domain.TurnoDetalle, error) {
	row := r.db.QueryRow(QueryGetTurnoByPacienteDni, dniPaciente)

	var turnoDetalle domain.TurnoDetalle
	err := row.Scan(
		&turnoDetalle.Id,
		&turnoDetalle.IdPaciente,
		&turnoDetalle.IdDentista,
		&turnoDetalle.PacienteNombre,
		&turnoDetalle.PacienteApellido,
		&turnoDetalle.FechaHora,
		&turnoDetalle.Descripcion,
	)
	if err != nil {
		return domain.TurnoDetalle{}, err
	}

	return turnoDetalle, nil
}

// CreateTurnoByDNIMatricula crea un turno asociado al paciente con el DNI y al dentista con la matrícula.
func (r *repository) CreateTurnoByDNIMatricula(ctx context.Context, turno domain.Turno, dniPaciente string, matriculaDentista string) (int, error) {

	idPaciente, err := r.GetPacienteIDByDNI(ctx, dniPaciente)
	if err != nil {
		return 0, err
	}

	idDentista, err := r.GetDentistaIDByMatricula(ctx, matriculaDentista)
	if err != nil {
		return 0, err
	}


	turno.IdPaciente = idPaciente
	turno.IdDentista = idDentista



	result, err := r.db.ExecContext(
		ctx,
		QueryInsertTurno,
		turno.IdPaciente,
		turno.IdDentista,
		turno.Fecha_hora,
		turno.Descripcion,
	)
	if err != nil {
		return 0, err
	}

	turnoID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(turnoID), nil
}

func (r *repository) GetPacienteIDByDNI(ctx context.Context, dniPaciente string) (int, error) {
	var idPaciente int
	err := r.db.QueryRowContext(ctx, QueryGetPacienteIDByDNI, dniPaciente).Scan(&idPaciente)
	if err != nil {
		return 0, err
	}
	return idPaciente, nil
}

func (r *repository) GetDentistaIDByMatricula(ctx context.Context, matriculaDentista string) (int, error) {
	var idDentista int
	err := r.db.QueryRowContext(ctx, QueryGetDentistaIDByMatricula, matriculaDentista).Scan(&idDentista)
	if err != nil {
		return 0, err
	}
	return idDentista, nil
}