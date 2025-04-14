package building

import (
	"building_management/interfaces/api/building"
	"building_management/models"
)

func mapBuildingRequestToModel(req building.Request) models.Building {
	return models.Building{
		ID:      req.ID,
		Name:    req.Name,
		Address: req.Address,
	}
}
