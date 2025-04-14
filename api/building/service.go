package building

import (
	"building_management/interfaces/api/building"
	"building_management/models"
	"context"
)

type Service struct {
	repo building.RepositoryInterface
}

func NewService(repo building.RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

// Get all buildings
func (s Service) GetBuildings(ctx context.Context) (models.BuildingSlice, error) {
	buildings, err := s.repo.GetBuildings(ctx)
	if err != nil {
		return nil, err
	}
	return buildings, nil
}

// Get building by ID
func (s Service) GetBuildingByID(ctx context.Context, id int) (*models.Building, error) {
	building, err := s.repo.GetBuildingByID(ctx, id)
	if err != nil {
		return building, err
	}
	return building, nil
}

// Create or update building
func (s Service) CreateOrUpdateBuilding(ctx context.Context, request building.Request) (*models.Building, error) {
	buildingModel := mapBuildingRequestToModel(request)
	build, err := s.repo.CreateOrUpdateBuilding(ctx, buildingModel)
	if err != nil {
		return &models.Building{}, err
	}
	return build, nil
}

// Delete building by ID
func (s Service) DeleteBuilding(ctx context.Context, id int) error {
	err := s.repo.DeleteBuilding(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
