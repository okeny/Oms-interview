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

func TestCreateOrUpdateBuilding(t *testing.T) {
	db, mock, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewRepository(db)
	ctx := context.Background()

	tests := []struct {
		name           string
		building       models.Building
		mockSetup      func(mock sqlmock.Sqlmock)
		expectedResult *models.Building
		expectedErrMsg string
	}{
		{
			name: "Success - insert new building",
			building: models.Building{
				ID:      0,
				Name:    "New Tower",
				Address: "123 Sky Ave",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "building" .* ON CONFLICT \("id"\) DO UPDATE SET .* RETURNING "id"`).
					WithArgs("New Tower", "123 Sky Ave", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(10))

				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id" = \$1\) LIMIT 1`).
					WithArgs(10).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "name", "address", "created_at", "updated_at",
					}).AddRow(10, "New Tower", "123 Sky Ave", time.Now(), time.Now()))

				mock.ExpectCommit()
			},
			expectedResult: &models.Building{ID: 10, Name: "New Tower", Address: "123 Sky Ave"},
		},
		{
			name: "Success - update existing building",
			building: models.Building{
				ID:      1,
				Name:    "Main Block",
				Address: "1 Main St",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "building" .* ON CONFLICT \("id"\) DO UPDATE SET .* RETURNING "id"`).
					WithArgs(1,"Main Block", "1 Main St", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id" = \$1\) LIMIT 1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "name", "address", "created_at", "updated_at",
					}).AddRow(1, "Main Block", "1 Main St", time.Now(), time.Now()))

				mock.ExpectCommit()
			},
			expectedResult: &models.Building{ID: 1, Name: "Main Block", Address: "1 Main St"},
		},
		{
			name: "Error - upsert fails",
			building: models.Building{
				ID:      2,
				Name:    "Fail Block",
				Address: "404 Error Rd",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "building" .* ON CONFLICT \("id"\) DO UPDATE SET .* RETURNING "id"`).
					WithArgs(2,"Fail Block", "404 Error Rd", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("upsert failure"))

				mock.ExpectRollback()
			},
			expectedResult: nil,
			expectedErrMsg: "failed to upsert building",
		},
		{
			name: "Error - fetch fails after upsert",
			building: models.Building{
				ID:      3,
				Name:    "Post Tower",
				Address: "500 Server Ln",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "building" .* ON CONFLICT \("id"\) DO UPDATE SET .* RETURNING "id"`).
					WithArgs(3,"Post Tower", "500 Server Ln", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))

				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id" = \$1\) LIMIT 1`).
					WithArgs(3).
					WillReturnError(errors.New("fetch error"))

				mock.ExpectRollback()
			},
			expectedResult: nil,
			expectedErrMsg: "failed to fetch upserted building",
		},
		{
			name: "Error - commit fails",
			building: models.Building{
				ID:      4,
				Name:    "Commit Fail",
				Address: "999 Commit Way",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "building" .* ON CONFLICT \("id"\) DO UPDATE SET .* RETURNING "id"`).
					WithArgs(4,"Commit Fail", "999 Commit Way", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(4))

				mock.ExpectQuery(`SELECT "building"\.\* FROM "building" WHERE \("building"\."id" = \$1\) LIMIT 1`).
					WithArgs(4).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "name", "address", "created_at", "updated_at",
					}).AddRow(4, "Commit Fail", "999 Commit Way", time.Now(), time.Now()))

				mock.ExpectCommit().WillReturnError(errors.New("commit error"))
			},
			expectedResult: nil,
			expectedErrMsg: "failed to commit transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mock)
			result, err := repo.CreateOrUpdateBuilding(ctx, tt.building)

			if tt.expectedResult != nil && result != nil {
				assert.Equal(t, tt.expectedResult.ID, result.ID)
				assert.Equal(t, tt.expectedResult.Name, result.Name)
				assert.Equal(t, tt.expectedResult.Address, result.Address)
			} else {
				assert.Nil(t, result)
			}

			checkError(t, err, tt.expectedErrMsg)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unmet expectations: %v", err)
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