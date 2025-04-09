package router

import (
	"building_management/api/building"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for the building controller
func BuildingInitRoute(appv1 fiber.Router, controller building.Controller) {
	// Define the routes for the building controller
	buildings := appv1.Group("/buildings")
	buildings.Get("/", controller.GetBuildings)
	buildings.Get("/:id", controller.GetBuildingByID)
	buildings.Post("/", controller.CreateOrUpdateBuilding)
	buildings.Delete("/:id", controller.DeleteBuilding)
}
