package router

import (
	"building_management/api/apartment"
	"building_management/api/building"
	"building_management/config"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Init sets up and returns the Fiber app
func Init(db *sql.DB) (*fiber.App, error) {
	app := fiber.New()
	app.Use(logger.New()) // Add logger middleware

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))

	// Get API version from config
	version, err := config.Get("API_VERSION")
	if err != nil || version == "" {
		log.Printf("API_VERSION not found, falling back to 'v1': %v", err)
		version = "v1"
	}
	apiVersion := app.Group("/api/" + version)

	// Initialize Building
	buildingRepo := building.NewRepository(db)
	buildingService := building.NewService(buildingRepo)
	buildingHandler := building.NewHandler()
	buildingController := building.NewController(buildingHandler, buildingService)

	// Initialize Apartment
	apartmentRepo := apartment.NewRepository(db)
	apartmentService := apartment.NewService(apartmentRepo)
	apartmentHandler := apartment.NewHandler(db)
	apartmentController := apartment.NewController(apartmentHandler, apartmentService)

	// Register routes
	BuildingInitRoute(apiVersion, buildingController)
	ApartmentInitRoute(apiVersion, apartmentController)
    HealthCheckInitRoute(apiVersion)
	return app, nil
}
