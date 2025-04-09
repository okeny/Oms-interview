package apartment

import (
	"building_management/models"
	"context"
)

type ServiceInterface interface {
	GetApartments(ctx context.Context) (models.ApartmentSlice, error)
	GetApartmentByID(ctx context.Context, id int) (*models.Apartment, error)
	GetApartmentsByBuilding(ctx context.Context, buildingID int) (models.ApartmentSlice, error)
	CreateOrUpdateApartment(ctx context.Context, apartment ApartmentRequest) (*models.Apartment, error)
	DeleteApartment(ctx context.Context, id int) error
}
