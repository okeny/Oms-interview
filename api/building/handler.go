package building

import (
	"building_management/interfaces/api/building"

	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) GetID(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (h Handler) GetCreateOrUpdateRequest(c *fiber.Ctx) (building.Request, error) {
	var request building.Request
	// Parse request body into apartment struct
	if err := c.BodyParser(&request); err != nil {
		return building.Request{}, err
	}
	return request, nil
}
