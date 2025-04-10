package building

import (
	"building_management/interfaces/api/building"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	handler building.HandlerInterface
	service building.ServiceInterface
}

func NewController(handler building.HandlerInterface, s building.ServiceInterface) Controller {
	return Controller{handler: handler, service: s}
}

func (ctl Controller) GetBuildings(c *fiber.Ctx) error {
	buildings, err := ctl.service.GetBuildings(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(buildings) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "buildings not found"})
	}

	return c.JSON(buildings)
}

func (ctl Controller) GetBuildingByID(c *fiber.Ctx) error {
	// Use the handler to get the building ID from the request context
	id, err := ctl.handler.GetID(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid building ID"})
	}
	// call the service to get the building by ID
	building, err := ctl.service.GetBuildingByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, ErrBuildingNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "building not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(building)
}

func (ctl Controller) CreateOrUpdateBuilding(c *fiber.Ctx) error {
	// Use the handler to get the building request from the request context
	request, err := ctl.handler.GetCreateOrUpdateRequest(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// call the service to create or update the building
	building, err := ctl.service.CreateOrUpdateBuilding(c.Context(), request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(building)
}

func (ctl Controller) DeleteBuilding(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ctl.service.DeleteBuilding(c.Context(), id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "building not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Building deleted",
	})
}
