package building

import (
	"building_management/models"
	"context"
)

type RepositoryInterface interface {
	GetBuildings(ctx context.Context) (models.BuildingSlice, error)
	GetBuildingByID(ctx context.Context, ID int) (*models.Building, error)
	CreateOrUpdateBuilding(ctx context.Context, building models.Building) (*models.Building, error)
	DeleteBuilding(ctx context.Context, id int) error
}