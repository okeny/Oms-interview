package apartment

import (
	"building_management/models"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

// Get all apartments
func (repo Repository) GetApartments(ctx context.Context) (models.ApartmentSlice, error) {
	apartments, err := models.Apartments().All(ctx, repo.DB)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

// Get apartment by ID
func (repo Repository) GetApartmentByID(ctx context.Context, id int) (*models.Apartment, error) {
	apartment, err := models.Apartments(models.ApartmentWhere.ID.EQ(id)).One(ctx, repo.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrApartmentNotFound
		}
		return &models.Apartment{}, err
	}
	return apartment, nil
}

// Get all apartments in a specific building
func (repo Repository) GetApartmentsByBuilding(ctx context.Context, buildingID int) (models.ApartmentSlice, error) {
	apartments, err := models.Apartments(models.ApartmentWhere.BuildingID.EQ(buildingID)).All(ctx, repo.DB)
	if err != nil {
		return nil, err
	}

	return apartments, nil
}

// Create or update apartment
func (repo Repository) CreateOrUpdateApartment(ctx context.Context, apartment models.Apartment) (*models.Apartment, error) {
	// Start transaction
	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not start transaction: %w", err)
	}

	// Track if we need to rollback
	var shouldRollback = true
	defer func() {
		if shouldRollback {
			_ = tx.Rollback()
		}
	}()

	// Perform the upsert
	if err := apartment.Upsert(
		ctx,
		tx,
		true,
		[]string{"id"},
		boil.Infer(),
		boil.Infer(),
	); err != nil {
		return nil, fmt.Errorf("failed to upsert apartment: %w", err)
	}

	// Fetch the updated apartment
	updatedApartment, err := models.Apartments(models.ApartmentWhere.ID.EQ(apartment.ID)).One(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch upserted apartment: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Mark that we committed successfully
	shouldRollback = false

	return updatedApartment, nil
}

// Delete apartment by ID
func (repo Repository) DeleteApartment(ctx context.Context, ID int) error {
	apartment := models.Apartment{ID: ID}
	n, err := apartment.Delete(ctx, repo.DB)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrApartmentNotFound
	}
	return nil
}
