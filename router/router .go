package router

import (
	"building_management/api/apartment"
	"building_management/api/building"
	"building_management/config"
	"building_management/middleware"
	"database/sql"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init sets up and returns the Fiber app
func Init(db *sql.DB) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName: "Building Management API",
	})
	// Add logger middleware
	app.Use(logger.New())
	// Add CORS middleware
	app.Use(middleware.HandleCorsMiddleware())

	prometheusPort, err := config.Get("PROMETHEUS_PORT")
	if err != nil {
		log.WithError(err).Fatal("Failed to get port from config")
	}

	// Add Prometheus custom middleware
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("Prometheus metrics available at :%s/metrics", prometheusPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", prometheusPort), nil); err != nil {
			log.Fatal("Metrics server failed:", err)
		}
	}()

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
