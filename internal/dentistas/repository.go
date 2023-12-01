package dentistas

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
	Create(ctx context.Context, dentista domain.Dentista) (domain.Dentista, error)
	GetAll(ctx context.Context) ([]domain.Dentista, error)
	GetByID(ctx context.Context, id int) (domain.Dentista, error)
	Update(ctx context.Context, dentista domain.Dentista, id int) (domain.Dentista, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, dentista domain.Dentista, id int) (domain.Dentista, error)
}


type repository struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create is a method that creates a new dentista.
func (r *repository) Create(ctx context.Context, dentista domain.Dentista) (domain.Dentista, error) {
	statement, err := r.db.Prepare(QueryInsertDentista)
	if err != nil {
		return domain.Dentista{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentista.Apellido,
		dentista.Nombre,
		dentista.Matricula,
	)

	if err != nil {
		return domain.Dentista{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Dentista{}, ErrLastInsertedId
	}

	dentista.Id = int(lastId)

	return dentista, nil

}

// GetAll is a method that returns all dentistas.
func (r *repository) GetAll(ctx context.Context) ([]domain.Dentista, error) {
	rows, err := r.db.Query(QueryGetAllDentistas)
	if err != nil {
		return []domain.Dentista{}, err
	}

	defer rows.Close()

	var dentistas []domain.Dentista

	for rows.Next() {
		var dentista domain.Dentista
		err := rows.Scan(
			&dentista.Id,
			&dentista.Apellido,
			&dentista.Nombre,
			&dentista.Matricula,
		)
		if err != nil {
			return []domain.Dentista{}, err
		}

		dentistas = append(dentistas, dentista)
	}

	if err := rows.Err(); err != nil {
		return []domain.Dentista{}, err
	}

	return dentistas, nil
}

// GetByID is a method that returns a dentista by ID.
func (r *repository) GetByID(ctx context.Context, id int) (domain.Dentista, error) {
	row := r.db.QueryRow(QueryGetDentistaById, id)

	var dentista domain.Dentista
	err := row.Scan(
		&dentista.Id,
		&dentista.Apellido,
		&dentista.Nombre,
		&dentista.Matricula,
	)

	if err != nil {
		return domain.Dentista{}, err
	}

	return dentista, nil
}

// Update is a method that updates a dentista by ID.
func (r *repository) Update(
	ctx context.Context,
	dentista domain.Dentista,
	id int) (domain.Dentista, error) {
	statement, err := r.db.Prepare(QueryUpdateDentista)
	if err != nil {
		return domain.Dentista{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentista.Nombre,
		dentista.Apellido,
		dentista.Matricula,
		dentista.Id,
	)

	if err != nil {
		return domain.Dentista{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentista{}, err
	}

	dentista.Id = id

	return dentista, nil

}

// Delete is a method that deletes a dentista by ID.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteDentista, id)
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

// Patch is a method that updates a dentista by ID.
func (r *repository) Patch(
	ctx context.Context,
	dentista domain.Dentista,
	id int) (domain.Dentista, error) {
	statement, err := r.db.Prepare(QueryUpdateDentista)
	if err != nil {
		return domain.Dentista{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentista.Nombre,
		dentista.Apellido,
		dentista.Matricula,
		dentista.Id,
	)

	if err != nil {
		return domain.Dentista{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentista{}, err
	}

	return dentista, nil
}