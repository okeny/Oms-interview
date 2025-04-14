package apartment

import (
	ap "building_management/interfaces/api/apartment"
	"building_management/mocks/interfaces/api/apartment"

	"building_management/models"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetApartments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := apartment.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name           string
		mockSetup      func()
		expectedResult models.ApartmentSlice
		expectedErr    error
	}{
		{
			name: "Success - apartments retrieved",
			mockSetup: func() {
				mockRepo.EXPECT().GetApartments(ctx).Return(models.ApartmentSlice{
					&models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50},
				}, nil)
			},
			expectedResult: models.ApartmentSlice{
				&models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50},
			},
			expectedErr: nil,
		},
		{
			name: "Error - repository fails",
			mockSetup: func() {
				mockRepo.EXPECT().GetApartments(ctx).Return(nil, errors.New("db error"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.GetApartments(ctx)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GetApartmentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := apartment.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name           string
		id             int
		mockSetup      func()
		expectedResult *models.Apartment
		expectedErr    error
	}{
		{
			name: "Success - apartment found",
			id:   1,
			mockSetup: func() {
				mockRepo.EXPECT().GetApartmentByID(ctx, 1).Return(&models.Apartment{ID: 1, BuildingID: 1, Number: "A1"}, nil)
			},
			expectedResult: &models.Apartment{ID: 1, BuildingID: 1, Number: "A1"},
			expectedErr:    nil,
		},
		{
			name: "Error - apartment not found",
			id:   2,
			mockSetup: func() {
				mockRepo.EXPECT().GetApartmentByID(ctx, 2).Return(nil, errors.New("not found"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.GetApartmentByID(ctx, tt.id)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GetApartmentsByBuilding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := apartment.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name           string
		buildingID     int
		mockSetup      func()
		expectedResult models.ApartmentSlice
		expectedErr    error
	}{
		{
			name:       "Success - apartments found",
			buildingID: 1,
			mockSetup: func() {
				mockRepo.EXPECT().GetApartmentsByBuilding(ctx, 1).Return(models.ApartmentSlice{
					&models.Apartment{ID: 1, BuildingID: 1, Number: "A1"},
				}, nil)
			},
			expectedResult: models.ApartmentSlice{
				&models.Apartment{ID: 1, BuildingID: 1, Number: "A1"},
			},
			expectedErr: nil,
		},
		{
			name:       "Success - no apartments",
			buildingID: 2,
			mockSetup: func() {
				mockRepo.EXPECT().GetApartmentsByBuilding(ctx, 2).Return(models.ApartmentSlice{}, nil)
			},
			expectedResult: models.ApartmentSlice{},
			expectedErr:    nil,
		},
		{
			name:       "Error - repository fails",
			buildingID: 3,
			mockSetup: func() {
				mockRepo.EXPECT().GetApartmentsByBuilding(ctx, 3).Return(nil, errors.New("db error"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.GetApartmentsByBuilding(ctx, tt.buildingID)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_CreateOrUpdateApartment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := apartment.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name           string
		request        ap.Request
		mockSetup      func()
		expectedResult *models.Apartment
		expectedErr    error
	}{
		{
			name: "Success - apartment created",
			request: ap.Request{
				ID:         0,
				BuildingID: 1,
				Number:     "A1",
				Floor:      1,
				SQMeters:   50,
			},
			mockSetup: func() {
				apartmentModel := models.Apartment{
					ID:         0,
					BuildingID: 1,
					Number:     "A1",
					Floor:      1,
					SQMeters:   50,
				}
				mockRepo.EXPECT().CreateOrUpdateApartment(ctx, apartmentModel).
					Return(&models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50}, nil)
			},
			expectedResult: &models.Apartment{ID: 1, BuildingID: 1, Number: "A1", Floor: 1, SQMeters: 50},
			expectedErr:    nil,
		},
		{
			name: "Error - repository fails",
			request: ap.Request{
				ID:         0,
				BuildingID: 1,
				Number:     "A1",
			},
			mockSetup: func() {
				apartmentModel := models.Apartment{
					ID:         0,
					BuildingID: 1,
					Number:     "A1",
				}
				mockRepo.EXPECT().CreateOrUpdateApartment(ctx, apartmentModel).
					Return(nil, errors.New("upsert error"))
			},
			expectedResult: &models.Apartment{},
			expectedErr:    errors.New("upsert error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.CreateOrUpdateApartment(ctx, tt.request)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_DeleteApartment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := apartment.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name        string
		id          int
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "Success - apartment deleted",
			id:   1,
			mockSetup: func() {
				mockRepo.EXPECT().DeleteApartment(ctx, 1).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error - repository fails",
			id:   2,
			mockSetup: func() {
				mockRepo.EXPECT().DeleteApartment(ctx, 2).Return(errors.New("delete error"))
			},
			expectedErr: errors.New("delete error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.DeleteApartment(ctx, tt.id)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
