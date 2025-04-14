package apartment

import (
	"building_management/interfaces/api/apartment"
	"errors"
	"building_management/metrics"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	handler apartment.HandlerInterface
	service apartment.ServiceInterface
}

func NewController(h apartment.HandlerInterface, s apartment.ServiceInterface) Controller {
	return Controller{handler: h, service: s}
}

// Get all apartments
func (ctl Controller) GetApartments(c *fiber.Ctx) error {
	apartments, err := ctl.service.GetApartments(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(apartments) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "apartments not found"})
	}
	return c.JSON(apartments)
}

// Get apartment by ID
func (ctl Controller) GetApartmentByID(c *fiber.Ctx) error {
	// Use the handler to get the apartment ID from the request context
	id, err := ctl.handler.GetID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid apartment ID"})
	}

	apartment, err := ctl.service.GetApartmentByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, ErrApartmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "apartment not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(apartment)
}

// Get all apartments in a specific building
func (ctl Controller) GetApartmentsByBuilding(c *fiber.Ctx) error {
	buildingId, err := ctl.handler.GetID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid building ID"})
	}

	apartments, err := ctl.service.GetApartmentsByBuilding(c.Context(), buildingId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(apartments) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "apartments not found"})
	}

	return c.JSON(apartments)
}

// Create or update apartment
func (ctl Controller) CreateOrUpdateApartment(c *fiber.Ctx) error {
	// Use the handler to get the apartment request from the request context
	apartmentRequest, err := ctl.handler.GetCreateOrUpdateRequest(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	apartment, err := ctl.service.CreateOrUpdateApartment(c.Context(), apartmentRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// increment the appartment counter
	if apartmentRequest.ID != 0 {
		metrics.ApartmentCreatedCounter.Inc()
	}

	// Wrap the response in a slice
	return c.Status(fiber.StatusCreated).JSON(apartment)
}

// Delete apartment by ID
func (ctl Controller) DeleteApartment(c *fiber.Ctx) error {
	id, err := ctl.handler.GetID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid apartment ID"})
	}

	if err := ctl.service.DeleteApartment(c.Context(), id); err != nil {
		if errors.Is(err, ErrApartmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "apartment not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
