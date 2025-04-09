package building

import "github.com/gofiber/fiber/v2"

type HandlerInterface interface {
	GetID(c *fiber.Ctx) (int, error)
	GetCreateOrUpdateRequest(c *fiber.Ctx) (BuildingRequest, error)
}
