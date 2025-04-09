package building

import (
	"building_management/mocks/interfaces/api/building"
	 b "building_management/interfaces/api/building"

	"building_management/models"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)
// TestService_GetBuildings tests the GetBuildings method of the Service struct
func TestService_GetBuildings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := building.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name           string
		mockSetup      func()
		expectedResult models.BuildingSlice
		expectedErr    error
	}{
		{
			name: "Success - buildings retrieved",
			mockSetup: func() {
				mockRepo.EXPECT().GetBuildings(ctx).Return(models.BuildingSlice{
					&models.Building{ID: 1, Name: "Building A", Address: "123 Main St"},
				}, nil)
			},
			expectedResult: models.BuildingSlice{
				&models.Building{ID: 1, Name: "Building A", Address: "123 Main St"},
			},
			expectedErr: nil,
		},
		{
			name: "Error - repository fails",
			mockSetup: func() {
				mockRepo.EXPECT().GetBuildings(ctx).Return(nil, errors.New("db error"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.GetBuildings(ctx)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GetBuildingByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := building.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name           string
		id             int
		mockSetup      func()
		expectedResult *models.Building
		expectedErr    error
	}{
		{
			name: "Success - building found",
			id:   1,
			mockSetup: func() {
				mockRepo.EXPECT().GetBuildingByID(ctx, 1).Return(&models.Building{ID: 1, Name: "Building A"}, nil)
			},
			expectedResult: &models.Building{ID: 1, Name: "Building A"},
			expectedErr:    nil,
		},
		{
			name: "Error - building not found",
			id:   2,
			mockSetup: func() {
				mockRepo.EXPECT().GetBuildingByID(ctx, 2).Return(nil, errors.New("not found"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.GetBuildingByID(ctx, tt.id)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_CreateOrUpdateBuilding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := building.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name           string
		request        b.BuildingRequest
		mockSetup      func()
		expectedResult *models.Building
		expectedErr    error
	}{
		{
			name: "Success - building created",
			request: b.BuildingRequest{
				ID:      0,
				Name:    "Building B",
				Address: "456 Oak St",
			},
			mockSetup: func() {
				buildingModel := models.Building{
					ID:      0,
					Name:    "Building B",
					Address: "456 Oak St",
				}
				mockRepo.EXPECT().CreateOrUpdateBuilding(ctx, buildingModel).
					Return(&models.Building{ID: 1, Name: "Building B", Address: "456 Oak St"}, nil)
			},
			expectedResult: &models.Building{ID: 1, Name: "Building B", Address: "456 Oak St"},
			expectedErr:    nil,
		},
		{
			name: "Error - repository fails",
			request: b.BuildingRequest{
				ID:      0,
				Name:    "Building C",
				Address: "789 Pine St",
			},
			mockSetup: func() {
				buildingModel := models.Building{
					ID:      0,
					Name:    "Building C",
					Address: "789 Pine St",
				}
				mockRepo.EXPECT().CreateOrUpdateBuilding(ctx, buildingModel).
					Return(nil, errors.New("upsert error"))
			},
			expectedResult: &models.Building{},
			expectedErr:    errors.New("upsert error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := service.CreateOrUpdateBuilding(ctx, tt.request)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_DeleteBuilding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := building.NewMockRepositoryInterface(ctrl)
	service := NewService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name        string
		id          int
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "Success - building deleted",
			id:   1,
			mockSetup: func() {
				mockRepo.EXPECT().DeleteBuilding(ctx, 1).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error - repository fails",
			id:   2,
			mockSetup: func() {
				mockRepo.EXPECT().DeleteBuilding(ctx, 2).Return(errors.New("delete error"))
			},
			expectedErr: errors.New("delete error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.DeleteBuilding(ctx, tt.id)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}