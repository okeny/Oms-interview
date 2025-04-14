package router

import (
	"building_management/api/apartment"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for the apartment controller
func ApartmentInitRoute(appv1 fiber.Router, controller apartment.Controller) {
	// Define the routes for the building controller
	apartments := appv1.Group("/apartments")
	// Define the routes for the apartment controller
	apartments.Get("/", controller.GetApartments)
	apartments.Get("/:id", controller.GetApartmentByID)
	apartments.Get("/building/:id", controller.GetApartmentsByBuilding)
	apartments.Post("/", controller.CreateOrUpdateApartment)
	apartments.Delete("/:id", controller.DeleteApartment)
}
