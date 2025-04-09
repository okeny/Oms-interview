package apartment

import (
	"building_management/models"
	"context"
)

type RepositoryInterface interface {
	GetApartments(ctx context.Context) (models.ApartmentSlice, error)
	GetApartmentByID(ctx context.Context, ID int) (*models.Apartment, error)
	GetApartmentsByBuilding(ctx context.Context, buildingID int) (models.ApartmentSlice, error)
	CreateOrUpdateApartment(ctx context.Context, apartment models.Apartment) (*models.Apartment, error)
	DeleteApartment(ctx context.Context, id int) error
}