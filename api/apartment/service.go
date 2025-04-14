package apartment

import (
	"building_management/interfaces/api/apartment"
	"building_management/models"
	"context"
)

type Service struct {
	repo apartment.RepositoryInterface
}

func NewService(repo apartment.RepositoryInterface) Service {
	return Service{
		repo: repo,
	}
}

// Get all apartments
func (s Service) GetApartments(ctx context.Context) (models.ApartmentSlice, error) {
	apartments, err := s.repo.GetApartments(ctx)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

// Get apartment by ID
func (s Service) GetApartmentByID(ctx context.Context, ID int) (*models.Apartment, error) {
	apartment, err := s.repo.GetApartmentByID(ctx, ID)
	if err != nil {
		return apartment, err
	}
	return apartment, nil
}

// Get all apartments in a specific building
func (s Service) GetApartmentsByBuilding(ctx context.Context,
	buildingID int) (models.ApartmentSlice, error) {

	apartments, err := s.repo.GetApartmentsByBuilding(ctx, buildingID)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

// Create or update apartment
func (s Service) CreateOrUpdateApartment(ctx context.Context,
	request apartment.Request) (*models.Apartment, error) {

	apartmentModel := mapApartmentRequestToModel(request)
	apart, err := s.repo.CreateOrUpdateApartment(ctx, apartmentModel)
	if err != nil {
		return &models.Apartment{}, err
	}
	return apart, nil
}

// Delete apartment by ID
func (s Service) DeleteApartment(ctx context.Context, id int) error {
	if err := s.repo.DeleteApartment(ctx, id); err != nil {
		return err
	}
	return nil
}
