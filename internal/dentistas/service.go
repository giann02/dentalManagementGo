package dentistas

import (
	"context"
	"log"

	"github.com/giann02/finalBackEndGo/internal/domain"
)

type Service interface {
	Create(ctx context.Context, dentista domain.Dentista) (domain.Dentista, error)
	GetAll(ctx context.Context) ([]domain.Dentista, error)
	GetByID(ctx context.Context, id int) (domain.Dentista, error)
	Update(ctx context.Context, dentista domain.Dentista, id int) (domain.Dentista, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, dentista domain.Dentista, id int) (domain.Dentista, error)
}

type service struct {
	repository Repository
}

func NewServiceDentista(repository Repository) Service {
	return &service{repository: repository}
}
// Create ....
func (s *service) Create(ctx context.Context, dentista domain.Dentista) (domain.Dentista, error) {
	dentista, err := s.repository.Create(ctx, dentista)
	if err != nil {
		log.Println("[DentistaService][Create] error creating dentista", err)
		return domain.Dentista{}, err
	}

	return dentista, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Dentista, error){
	listDentista, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("error searching all dentistas", err)
		return []domain.Dentista{}, err
	}
	return listDentista, nil
}

// GetByID is a method that returns a dentista by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Dentista, error) {
	dentista, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistaService][GetByID] error getting dentista by ID", err)
		return domain.Dentista{}, err
	}

	return dentista, nil
}

// Update is a method that updates a dentista by ID.
func (s *service) Update(ctx context.Context, dentista domain.Dentista, id int) (domain.Dentista, error) {
	dentista, err := s.repository.Update(ctx, dentista, id)
	if err != nil {
		log.Println("[DentistaService][Update] error updating dentista by ID", err)
		return domain.Dentista{}, err
	}

	return dentista, nil
}

// Delete is a method that deletes a dentista by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[DentistaService][Delete] error deleting dentista by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a dentista by ID.
func (s *service) Patch(ctx context.Context, dentista domain.Dentista, id int) (domain.Dentista, error) {
	dentistaStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[dentistaService][Patch] error getting dentista by ID", err)
		return domain.Dentista{}, err
	}

	dentistaPatch, err := s.validatePatch(dentistaStore, dentista)
	if err != nil {
		log.Println("[DentistaService][Patch] error validating dentista", err)
		return domain.Dentista{}, err
	}

	dentista, err = s.repository.Patch(ctx, dentistaPatch, id)
	if err != nil {
		log.Println("[DentistaService][Patch] error patching dentista by ID", err)
		return domain.Dentista{}, err
	}

	return dentista, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(dentistaStore, dentista domain.Dentista) (domain.Dentista, error) {

	if dentista.Apellido != "" {
		dentistaStore.Apellido = dentista.Apellido
	}

	if dentista.Nombre != "" {
		dentistaStore.Nombre = dentista.Nombre
	}

	if dentista.Matricula != "" {
		dentistaStore.Matricula = dentista.Matricula
	}

	return dentistaStore, nil

}