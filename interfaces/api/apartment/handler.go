package apartment

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type HandlerInterface interface {
	GetID(c *fiber.Ctx) (int, error)
	GetCreateOrUpdateRequest(c *fiber.Ctx) (ApartmentRequest, error)
	CheckBuilding(ctx context.Context, buildingID int) (bool, error)
}
