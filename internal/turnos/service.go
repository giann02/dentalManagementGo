package turnos

import (
	"context"
	"log"
	"strconv"
	"github.com/giann02/finalBackEndGo/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Turno, error)
	GetByID(ctx context.Context, id int) (domain.Turno, error)
	Update(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
	GetTurnoByDNIDetalle(ctx context.Context, dniPaciente string) (domain.TurnoDetalle, error)
	CreateTurnoByDNIMatricula(ctx context.Context, turno domain.Turno, dniPaciente string, matriculaDentista string) (int, error)
}

type service struct {
	repository Repository
}

func NewServiceTurno(repository Repository) Service {
	return &service{repository: repository}
}


func (s *service) GetAll(ctx context.Context) ([]domain.Turno, error){
	listTurno, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("error searching all turnos", err)
		return []domain.Turno{}, err
	}
	return listTurno, nil
}

// GetByID is a method that returns a turno by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Turno, error) {
	turno, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[TurnoService][GetByID] error getting turno by ID", err)
		return domain.Turno{}, err
	}

	return turno, nil
}

// Update is a method that updates a turno by ID.
func (s *service) Update(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error) {
	turno, err := s.repository.Update(ctx, turno, id)
	if err != nil {
		log.Println("[TurnoService][Update] error updating turno by ID", err)
		return domain.Turno{}, err
	}

	return turno, nil
}

// Delete is a method that deletes a turno by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[TurnoService][Delete] error deleting turno by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a turno by ID.
func (s *service) Patch(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error) {
	turnoStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[turnoService][Patch] error getting turno by ID", err)
		return domain.Turno{}, err
	}

	turnoPatch, err := s.validatePatch(turnoStore, turno)
	if err != nil {
		log.Println("[TurnoService][Patch] error validating turno", err)
		return domain.Turno{}, err
	}

	turno, err = s.repository.Patch(ctx, turnoPatch, id)
	if err != nil {
		log.Println("[TurnoService][Patch] error patching turno by ID", err)
		return domain.Turno{}, err
	}

	return turno, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(turnoStore, turno domain.Turno) (domain.Turno, error) {

	if strconv.Itoa(turno.IdPaciente) != "" {
		turnoStore.IdPaciente = turno.IdPaciente
	}

	if strconv.Itoa(turno.IdDentista) != "" {
		turnoStore.IdDentista = turno.IdDentista
	}

	if turno.Descripcion != "" {
		turnoStore.Descripcion = turno.Descripcion
	}

	return turnoStore, nil

}

// GetTurnoByDNIDetalle obtiene el detalle de un turno por el DNI del paciente.
func (s *service) GetTurnoByDNIDetalle(ctx context.Context, dniPaciente string) (domain.TurnoDetalle, error) {
	// Llama al método correspondiente en el repositorio para obtener el detalle del turno.
	turnoDetalle, err := s.repository.GetTurnoByDNIDetalle(ctx, dniPaciente)
	if err != nil {
		log.Println("[TurnoService][GetTurnoByDNIDetalle] error getting turno by ID", err)
		return domain.TurnoDetalle{}, err
	}
	return turnoDetalle, nil
}

// CreateTurnoByDNIMatricula crea un turno asociado al paciente con el DNI y al dentista con la matrícula.
func (s *service) CreateTurnoByDNIMatricula(ctx context.Context, turno domain.Turno, dniPaciente string, matriculaDentista string) (int, error) {

	turnoID, err := s.repository.CreateTurnoByDNIMatricula(ctx, turno, dniPaciente, matriculaDentista)
	if err != nil {
		log.Println("[TurnoService][CreateTurnoByDNIMatricula] error creating turno by Dni y Matricula", err)
		return 0, err
	}

	return turnoID, nil
}

