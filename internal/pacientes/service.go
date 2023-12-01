package pacientes

import (
	"context"
	"log"

	"github.com/giann02/finalBackEndGo/internal/domain"
)

type Service interface {
	Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error)
	GetAll(ctx context.Context) ([]domain.Paciente, error)
	GetByID(ctx context.Context, id int) (domain.Paciente, error)
	Update(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error)
}

type service struct {
	repository Repository
}

func NewServicePaciente(repository Repository) Service {
	return &service{repository: repository}
}
// Create ....
func (s *service) Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error) {
	paciente, err := s.repository.Create(ctx, paciente)
	if err != nil {
		log.Println("[PacienteService][Create] error creating paciente", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Paciente, error){
	listPaciente, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("error searching all pacientes", err)
		return []domain.Paciente{}, err
	}
	return listPaciente, nil
}

// GetByID is a method that returns a paciente by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Paciente, error) {
	paciente, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PacienteService][GetByID] error getting paciente by ID", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Update is a method that updates a paciente by ID.
func (s *service) Update(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error) {
	paciente, err := s.repository.Update(ctx, paciente, id)
	if err != nil {
		log.Println("[PacienteService][Update] error updating paciente by ID", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Delete is a method that deletes a paciente by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[PacienteService][Delete] error deleting paciente by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a paciente by ID.
func (s *service) Patch(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error) {
	pacienteStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[pacienteService][Patch] error getting paciente by ID", err)
		return domain.Paciente{}, err
	}

	pacientePatch, err := s.validatePatch(pacienteStore, paciente)
	if err != nil {
		log.Println("[PacienteService][Patch] error validating paciente", err)
		return domain.Paciente{}, err
	}

	paciente, err = s.repository.Patch(ctx, pacientePatch, id)
	if err != nil {
		log.Println("[PacienteService][Patch] error patching paciente by ID", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(pacienteStore, paciente domain.Paciente) (domain.Paciente, error) {

	if paciente.Apellido != "" {
		pacienteStore.Apellido = paciente.Apellido
	}

	if paciente.Nombre != "" {
		pacienteStore.Nombre = paciente.Nombre
	}

	if paciente.Domicilio != "" {
		pacienteStore.Domicilio = paciente.Domicilio
	}

	if paciente.Dni != "" {
		pacienteStore.Dni = paciente.Dni
	}

	if paciente.Fecha_alta != "" {
		pacienteStore.Fecha_alta = paciente.Fecha_alta
	}

	return pacienteStore, nil

}