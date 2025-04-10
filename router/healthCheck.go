package router

import "github.com/gofiber/fiber/v2"

func HealthCheckInitRoute(appv1 fiber.Router) {

	appv1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}
