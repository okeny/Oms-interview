package apartment

import (
	"building_management/interfaces/api/apartment"
	"building_management/models"
)

func mapApartmentRequestToModel(request apartment.Request) models.Apartment {
	return models.Apartment{
		ID:         request.ID,
		BuildingID: request.BuildingID,
		Number:     request.Number,
		Floor:      request.Floor,
		SQMeters:   request.SQMeters,
	}
}
