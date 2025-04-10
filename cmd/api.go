package cmd

import (
	"building_management/config"
	"building_management/database"
	"building_management/router"
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func API() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Start BMS in API mode",
		Run: func(cmd *cobra.Command, args []string) {
			runAPI()
		},
	}
}

func runAPI() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.WithError(err).Fatal("Failed to load environment configuration")
	}

	port, err := config.Get(Port)
	if err != nil {
		log.WithError(err).Fatal("Failed to get port from config")
	}
	host := fmt.Sprintf(":%s", port)

	// Initialize DB
	db, err := database.NewClient()
	if err != nil {
		log.WithError(err).Fatal("Failed to initialize database client")
	}
	defer db.Close()

	// Initialize the app/router
	app, err := router.Init(db)
	if err != nil {
		log.WithError(err).Fatal("Failed to initialize router")
	}

	// Setup signal catching for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start the server in a goroutine
	go func() {
		log.Infof("Server starting on %s", host)
		if err := app.Listen(host); err != nil {
			log.WithError(err).Error("Server stopped unexpectedly")
			stop() // trigger graceful shutdown
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	log.Info("Shutdown signal received. Cleaning up...")

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.WithError(err).Fatal("Error during server shutdown")
	}

	log.Info("Server shutdown complete.")
}
