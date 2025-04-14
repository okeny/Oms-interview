package building

import (
	"building_management/models"
	"context"
)

type ServiceInterface interface {
	GetBuildings(ctx context.Context) (models.BuildingSlice, error)
	GetBuildingByID(ctx context.Context, id int) (*models.Building, error)
	CreateOrUpdateBuilding(ctx context.Context, building Request) (*models.Building, error)
	DeleteBuilding(ctx context.Context, id int) error
}
