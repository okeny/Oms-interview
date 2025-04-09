package apartment

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"building_management/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// setupTestDB initializes the mock database
func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	return db, mock, func() { db.Close() }
}

// TestGetApartments tests retrieving all apartments
func TestGetApartments(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRepository(db)
	ctx := context.Background()

	tests := []struct {
		name           string
		mockSetup      func(sqlmock.Sqlmock)
		expectedResult models.ApartmentSlice
		expectedErrMsg string
	}{
		{
			name: "Success - multiple apartments",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "building_id", "number", "floor", "sq_meters"}).
					AddRow(1, 1, "A1", 1, 50).
					AddRow(2, 1, "A2", 2, 60)
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment"`).
					WillReturnRows(rows)
			},
			expectedResult: models.ApartmentSlice{
				&models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50},
				&models.Apartment{ID: 2, BuildingID: 1, Number: "A2", Floor: 2, SQMeters: 60},
			},
			expectedErrMsg: "", // No error expected
		},
		{
			name: "Error - database failure",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment"`).
					WillReturnError(errors.New("db error"))
			},
			expectedResult: nil,
			expectedErrMsg: "db error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			result, err := repo.GetApartments(ctx)

			// Compare results
			assert.Equal(t, tt.expectedResult, result)

			// Compare errors using checkError helper
			checkError(t, err, tt.expectedErrMsg)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// TestGetApartmentByID tests retrieving an apartment by ID
func TestGetApartmentByID(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRepository(db)
	ctx := context.Background()

	tests := []struct {
		name           string
		id             int
		mockSetup      func(sqlmock.Sqlmock)
		expectedResult *models.Apartment
		expectedErrMsg string
	}{
		{
			name: "Success - apartment found",
			id:   1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "building_id", "number", "floor", "sq_meters"}).
					AddRow(1, 1, "A1", 1, 50)
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment" WHERE \("apartment"\."id" = \$1\) LIMIT 1`).
					WithArgs(1).
					WillReturnRows(rows)
			},
			expectedResult: &models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50},
			expectedErrMsg: "",
		},
		{
			name: "Error - apartment not found",
			id:   2,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment" WHERE \("apartment"\."id" = \$1\) LIMIT 1`).
					WithArgs(2).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedErrMsg: "apartment not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			result, err := repo.GetApartmentByID(ctx, tt.id)

			// Compare results
			assert.Equal(t, tt.expectedResult, result)

			// Compare errors using checkError helper
			checkError(t, err, tt.expectedErrMsg)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// TestGetApartmentsByBuilding tests retrieving apartments by building ID
func TestGetApartmentsByBuilding(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRepository(db)
	ctx := context.Background()

	tests := []struct {
		name           string
		buildingID     int
		mockSetup      func(sqlmock.Sqlmock)
		expectedResult models.ApartmentSlice
		expectedErrMsg string
	}{
		{
			name:       "Success - apartments found",
			buildingID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "building_id", "number", "floor", "sq_meters"}).
					AddRow(1, 1, "A1", 1, 50)
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment" WHERE \("apartment"\."building_id" = \$1\)`).
					WithArgs(1).
					WillReturnRows(rows)
			},
			expectedResult: models.ApartmentSlice{
				&models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50},
			},
			expectedErrMsg: "",
		},
		{
			name:       "Success - no apartments",
			buildingID: 2,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "building_id", "number", "floor", "sq_meters"})
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment" WHERE \("apartment"\."building_id" = \$1\)`).
					WithArgs(2).
					WillReturnRows(rows)
			},
			expectedResult: nil, // Empty slice, not nil
			expectedErrMsg: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			result, err := repo.GetApartmentsByBuilding(ctx, tt.buildingID)

			// Compare results
			assert.Equal(t, tt.expectedResult, result)

			// Compare errors using checkError helper
			checkError(t, err, tt.expectedErrMsg)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// TestCreateOrUpdateApartment tests creating or updating an apartment
func TestCreateOrUpdateApartment(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRepository(db)
	ctx := context.Background()

	tests := []struct {
		name           string
		apartment      models.Apartment
		mockSetup      func(sqlmock.Sqlmock)
		expectedResult *models.Apartment
		expectedErrMsg string
	}{
		{
			name: "Success - insert new apartment",
			apartment: models.Apartment{
				ID:         0, // New apartment
				BuildingID: 1,
				Number:     "A2",
				Floor:      2,
				SQMeters:   60,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				// Upsert for insert, matching actual sqlboiler SQL
				mock.ExpectQuery(`INSERT INTO "apartment" \("building_id",\s*"number",\s*"floor",\s*"sq_meters",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5,\s*\$6\) ON CONFLICT \("id"\) DO UPDATE SET "building_id"\s*=\s*EXCLUDED\."building_id",\s*"number"\s*=\s*EXCLUDED\."number",\s*"floor"\s*=\s*EXCLUDED\."floor",\s*"sq_meters"\s*=\s*EXCLUDED\."sq_meters",\s*"created_at"\s*=\s*EXCLUDED\."created_at",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id"`).
					WithArgs(1, "A2", 2, 60, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
				// Fetch the full record
				rows := sqlmock.NewRows([]string{"id", "building_id", "number", "floor", "sq_meters", "created_at", "updated_at"}).
					AddRow(2, 1, "A2", 2, 60, time.Now(), time.Now())
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment" WHERE \("apartment"\."id"\s*=\s*\$1\) LIMIT 1`).
					WithArgs(2).
					WillReturnRows(rows)
			},
			expectedResult: &models.Apartment{
				ID:         2,
				BuildingID: 1,
				Number:     "A2",
				Floor:      2,
				SQMeters:   60,
			},
			expectedErrMsg: "",
		},
		{
			name: "Success - update existing apartment",
			apartment: models.Apartment{
				ID:         1,
				BuildingID: 1,
				Number:     "A1",
				Floor:      1,
				SQMeters:   55,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				// Upsert for update, matching actual sqlboiler SQL
				mock.ExpectQuery(`INSERT INTO "apartment" \("id",\s*"building_id",\s*"number",\s*"floor",\s*"sq_meters",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5,\s*\$6,\s*\$7\) ON CONFLICT \("id"\) DO UPDATE SET "building_id"\s*=\s*EXCLUDED\."building_id",\s*"number"\s*=\s*EXCLUDED\."number",\s*"floor"\s*=\s*EXCLUDED\."floor",\s*"sq_meters"\s*=\s*EXCLUDED\."sq_meters",\s*"created_at"\s*=\s*EXCLUDED\."created_at",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id"`).
					WithArgs(1, 1, "A1", 1, 55, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				// Fetch the full record
				rows := sqlmock.NewRows([]string{"id", "building_id", "number", "floor", "sq_meters", "created_at", "updated_at"}).
					AddRow(1, 1, "A1", 1, 55, time.Now(), time.Now())
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment" WHERE \("apartment"\."id"\s*=\s*\$1\) LIMIT 1`).
					WithArgs(1).
					WillReturnRows(rows)
			},
			expectedResult: &models.Apartment{
				ID:         1,
				BuildingID: 1,
				Number:     "A1",
				Floor:      1,
				SQMeters:   55,
			},
			expectedErrMsg: "",
		},
		{
			name: "Error - upsert fails",
			apartment: models.Apartment{
				ID:         1,
				BuildingID: 1,
				Number:     "A1",
				Floor:      1,
				SQMeters:   55,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO "apartment" \("id",\s*"building_id",\s*"number",\s*"floor",\s*"sq_meters",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5,\s*\$6,\s*\$7\) ON CONFLICT \("id"\) DO UPDATE SET "building_id"\s*=\s*EXCLUDED\."building_id",\s*"number"\s*=\s*EXCLUDED\."number",\s*"floor"\s*=\s*EXCLUDED\."floor",\s*"sq_meters"\s*=\s*EXCLUDED\."sq_meters",\s*"created_at"\s*=\s*EXCLUDED\."created_at",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id"`).
					WithArgs(1, 1, "A1", 1, 55, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("upsert error"))
			},
			expectedResult: nil,
			expectedErrMsg: "failed to upsert apartment",
		},
		{
			name: "Error - fetch fails after upsert",
			apartment: models.Apartment{
				ID:         0,
				BuildingID: 1,
				Number:     "A3",
				Floor:      3,
				SQMeters:   70,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				// Upsert for insert
				mock.ExpectQuery(`INSERT INTO "apartment" \("building_id",\s*"number",\s*"floor",\s*"sq_meters",\s*"created_at",\s*"updated_at"\) VALUES \(\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5,\s*\$6\) ON CONFLICT \("id"\) DO UPDATE SET "building_id"\s*=\s*EXCLUDED\."building_id",\s*"number"\s*=\s*EXCLUDED\."number",\s*"floor"\s*=\s*EXCLUDED\."floor",\s*"sq_meters"\s*=\s*EXCLUDED\."sq_meters",\s*"created_at"\s*=\s*EXCLUDED\."created_at",\s*"updated_at"\s*=\s*EXCLUDED\."updated_at" RETURNING "id"`).
					WithArgs(1, "A3", 3, 70, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))
				// Fetch fails
				mock.ExpectQuery(`SELECT "apartment"\.\* FROM "apartment" WHERE \("apartment"\."id"\s*=\s*\$1\) LIMIT 1`).
					WithArgs(3).
					WillReturnError(errors.New("fetch error"))
			},
			expectedResult: nil,
			expectedErrMsg: "failed to fetch upserted apartment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			result, err := repo.CreateOrUpdateApartment(ctx, tt.apartment)

			// Compare results
			if tt.expectedResult != nil && result != nil {
				assert.Equal(t, tt.expectedResult.ID, result.ID)
				assert.Equal(t, tt.expectedResult.BuildingID, result.BuildingID)
				assert.Equal(t, tt.expectedResult.Number, result.Number)
				assert.Equal(t, tt.expectedResult.Floor, result.Floor)
				assert.Equal(t, tt.expectedResult.SQMeters, result.SQMeters)
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
// TestDeleteApartment tests deleting an apartment by ID
func TestDeleteApartment(t *testing.T) {
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
			name: "Success - apartment deleted",
			id:   1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "apartment" WHERE "id"=\$1`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: "",
		},
		{
			name: "Error - no rows affected",
			id:   1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "apartment" WHERE "id"=\$1`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: "apartment not found",
		},
		{
			name: "Error - database failure",
			id:   1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "apartment" WHERE "id"=\$1`).
					WithArgs(1).
					WillReturnError(errors.New("delete error"))
			},
			expectedErr: "delete error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			err := repo.DeleteApartment(ctx, tt.id)

			// Compare errors using checkError helper
			checkError(t, err, tt.expectedErr)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// checkError is a helper to validate error assertions
func checkError(t *testing.T, err error, expectedErrMsg string) {
	if expectedErrMsg != "" {
		assert.Error(t, err, "expected an error but got none")
		assert.Contains(t, err.Error(), expectedErrMsg,
			"expected error to contain %q, got %v", expectedErrMsg, err)
	} else {
		assert.NoError(t, err, "expected no error but got %v", err)
	}
}
