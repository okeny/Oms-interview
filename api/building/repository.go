package building

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

// Get all buildings
func (repo Repository) GetBuildings(ctx context.Context) (models.BuildingSlice, error) {
	buildings, err := models.Buildings().All(ctx, repo.DB)
	if err != nil {
		return nil, err
	}
	return buildings, nil
}

// Get building by ID
func (repo Repository) GetBuildingByID(ctx context.Context, id int) (*models.Building, error) {
	building, err := models.Buildings(models.BuildingWhere.ID.EQ(id)).One(ctx, repo.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("building with id %d not found: %w", id, sql.ErrNoRows)
		}
		return nil, err
	}

	return building, nil
}

func (repo Repository) CreateOrUpdateBuilding(ctx context.Context, building models.Building) (*models.Building, error) {

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

	// Create or update apartment
	if err := building.Upsert(
		ctx,
		tx,
		true,
		[]string{"id"},
		boil.Infer(),
		boil.Infer(),
	); err != nil {
		return nil, fmt.Errorf("failed to upsert building: %w", err)
	}

	updatedBuilding, err := models.Buildings(models.BuildingWhere.ID.EQ(building.ID)).One(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch upserted building: %w", err)
	}
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Mark that we committed successfully
	shouldRollback = false

	return updatedBuilding, nil
}

// Delete building by ID
func (repo Repository) DeleteBuilding(ctx context.Context, id int) error {
	building := models.Building{ID: id}
	rowsAffected, err := building.Delete(ctx, repo.DB)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrBuildingNotFound
	}
	return nil
}
