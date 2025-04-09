package building

import (
	"building_management/models"
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	cleanup := func() {
		db.Close()
	}
	//for test purposes only
	boil.DebugMode = true

	return db, mock, cleanup
}

func TestRepository_GetBuildings(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	t.Run("Success - buildings retrieved", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "address"}).
			AddRow(1, "Building A", "123 Main St")

		mock.ExpectQuery(`SELECT "building"\.\* FROM "building"`).
			WillReturnRows(rows)

		ctx := context.Background()
		buildings, err := repo.GetBuildings(ctx)

		assert.NoError(t, err)
		assert.Len(t, buildings, 1)
		assert.Equal(t, 1, buildings[0].ID)
		assert.Equal(t, "Building A", buildings[0].Name)
		assert.Equal(t, "123 Main St", buildings[0].Address)
	})

	t.Run("Error - database query fails", func(t *testing.T) {
		mock.ExpectQuery(`SELECT "building"\.\* FROM "building"`).
			WillReturnError(errors.New("db error"))

		ctx := context.Background()
		buildings, err := repo.GetBuildings(ctx)

		assert.Error(t, err)
		assert.Nil(t, buildings)
		assert.Contains(t, err.Error(), "db error")
	})
}

func TestRepository_GetBuildingByID(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRepository(db)
	ctx := context.Background()

	tests := []struct {
		name           string
		id             int
		mockSetup      func(sqlmock.Sqlmock)
		expectedResult *models.Building
		expectedErr    string
	}{
		{
			name: "Success - building found",
			id:   1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "address"}).
					AddRow(1, "Building A", "123 Main St")
				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id" = \$1\) LIMIT 1;?`).
					WithArgs(1).
					WillReturnRows(rows)
			},
			expectedResult: &models.Building{ID: 1, Name: "Building A", Address: "123 Main St"},
			expectedErr:    "",
		},
		{
			name: "Error - building not found",
			id:   2,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id" = \$1\) LIMIT 1;?`).
					WithArgs(2).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedErr:    "no rows",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			result, err := repo.GetBuildingByID(ctx, tt.id)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// checkError is a helper to compare errors
func checkError(t *testing.T, err error, expectedMsg string) {
	if expectedMsg == "" {
		assert.NoError(t, err, "expected no error, got %v", err)
	} else {
		assert.Error(t, err, "expected an error, got nil")
		assert.Contains(t, err.Error(), expectedMsg, "expected error message to contain %q, got %v", expectedMsg, err)
	}
}

func TestRepository_CreateOrUpdateBuilding(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := Repository{DB: db} // Assuming NewRepository returns Repository{DB: db}
	ctx := context.Background()

	tests := []struct {
		name           string
		building       models.Building
		mockSetup      func(sqlmock.Sqlmock)
		expectedResult *models.Building
		expectedErrMsg string
	}{
		{
			name: "Success - insert new building",
			building: models.Building{
				ID:      0, // New building, ID assigned by DB
				Name:    "Test Building",
				Address: "123 Main St",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				// Insert without id for new building
				mock.ExpectQuery(`INSERT INTO "building" \("name",\s*"address",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4\) ON CONFLICT \("id"\) DO UPDATE SET "name"\s*=\s*EXCLUDED\."name",\s*"address"\s*=\s*EXCLUDED\."address",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id",\s*"created_at"`).
					WithArgs("Test Building", "123 Main St", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
						AddRow(11, time.Now())) // Next ID after sample data
				// Fetch with returned ID
				rows := sqlmock.NewRows([]string{"id", "name", "address", "created_at", "updated_at"}).
					AddRow(11, "Test Building", "123 Main St", time.Now(), time.Now())
				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id"\s*=\s*\$1\) LIMIT 1`).
					WithArgs(11).
					WillReturnRows(rows)
			},
			expectedResult: &models.Building{
				ID:      11,
				Name:    "Test Building",
				Address: "123 Main St",
			},
			expectedErrMsg: "",
		},
		{
			name: "Success - update existing building",
			building: models.Building{
				ID:      1,
				Name:    "Updated Towers",
				Address: "456 New Ave",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				// Update with id
				mock.ExpectQuery(`INSERT INTO "building" \("id",\s*"name",\s*"address",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5\) ON CONFLICT \("id"\) DO UPDATE SET "name"\s*=\s*EXCLUDED\."name",\s*"address"\s*=\s*EXCLUDED\."address",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id",\s*"created_at"`).
					WithArgs(1, "Updated Towers", "456 New Ave", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
						AddRow(1, time.Now()))
				rows := sqlmock.NewRows([]string{"id", "name", "address", "created_at", "updated_at"}).
					AddRow(1, "Updated Towers", "456 New Ave", time.Now(), time.Now())
				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id"\s*=\s*\$1\) LIMIT 1`).
					WithArgs(1).
					WillReturnRows(rows)
			},
			expectedResult: &models.Building{
				ID:      1,
				Name:    "Updated Towers",
				Address: "456 New Ave",
			},
			expectedErrMsg: "",
		},
		{
			name: "Error - upsert fails",
			building: models.Building{
				ID:      1,
				Name:    "Test Building",
				Address: "123 Main St",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO "building" \("id",\s*"name",\s*"address",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5\) ON CONFLICT \("id"\) DO UPDATE SET "name"\s*=\s*EXCLUDED\."name",\s*"address"\s*=\s*EXCLUDED\."address",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id",\s*"created_at"`).
					WithArgs(1, "Test Building", "123 Main St", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("upsert error"))
			},
			expectedResult: nil,
			expectedErrMsg: "failed to upsert building",
		},
		{
			name: "Error - fetch upserted building fails",
			building: models.Building{
				ID:      1,
				Name:    "Test Building",
				Address: "123 Main St",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO "building" \("id",\s*"name",\s*"address",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5\) ON CONFLICT \("id"\) DO UPDATE SET "name"\s*=\s*EXCLUDED\."name",\s*"address"\s*=\s*EXCLUDED\."address",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id",\s*"created_at"`).
					WithArgs(1, "Test Building", "123 Main St", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
						AddRow(1, time.Now()))
				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id"\s*=\s*\$1\) LIMIT 1`).
					WithArgs(1).
					WillReturnError(errors.New("fetch error"))
			},
			expectedResult: nil,
			expectedErrMsg: "failed to fetch upserted building",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			result, err := repo.CreateOrUpdateBuilding(ctx, tt.building)

			// Compare results
			if tt.expectedResult != nil && result != nil {
				assert.Equal(t, tt.expectedResult.ID, result.ID)
				assert.Equal(t, tt.expectedResult.Name, result.Name)
				assert.Equal(t, tt.expectedResult.Address, result.Address)
				// Note: created_at and updated_at are set by DB, so we skip exact comparison
			} else {
				assert.Nil(t, result, "expected nil result, got %+v", result)
			}

			// Compare errors using checkError helper
			checkError(t, err, tt.expectedErrMsg)

			// Check for unmet expectations
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unmet expectations: %v", err)
			}
		})
	}
}
func TestRepository_DeleteBuilding(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRepository(db)
	ctx := context.Background()

	tests := []struct {
		name        string
		id          int
		mockSetup   func(sqlmock.Sqlmock)
		expectedErr string
	}{
		{
			name: "Success - building deleted",
			id:   1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "building" WHERE "id"\s*=\s*\$1`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected
			},
			expectedErr: "",
		},
		{
			name: "Error - no rows deleted",
			id:   2,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "building" WHERE "id"\s*=\s*\$1`).
					WithArgs(2).
					WillReturnResult(sqlmock.NewResult(0, 0)) // 0 rows affected
			},
			expectedErr: "building not found", // Updated to match actual error
		},
		{
			name: "Error - database failure",
			id:   3,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "building" WHERE "id"\s*=\s*\$1`).
					WithArgs(3).
					WillReturnError(errors.New("delete error"))
			},
			expectedErr: "delete error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			err := repo.DeleteBuilding(ctx, tt.id)

			// Compare errors using checkError helper
			checkError(t, err, tt.expectedErr)

			// Check for unmet expectations
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unmet expectations: %v", err)
			}
		})
	}
}