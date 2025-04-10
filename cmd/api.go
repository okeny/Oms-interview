package cmd

import (
	"building_management/config"
	"building_management/database"
	"building_management/router"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	if err := config.LoadEnv(); err != nil {
		log.WithError(err).Fatal("Failed to load config")
	}

	port, err := config.Get(Port)
	if err != nil {
		log.WithError(err).Fatal("Failed to get port")
	}

	// Initialize DB and inject into services if needed
	db, err := database.NewClient()
	if err != nil {
		log.WithError(err).Fatal("error creating database client")
	}
	defer db.Close()

	app, err := router.Init(db)
	if err != nil {
		log.WithError(err).Fatal("err initializing router")
	}
	// Start the server
	host := fmt.Sprintf(":%s", port)
	go func() {
		if err := app.Listen(host); err != nil {
			log.WithError(err).Fatal("Http server closed unexpectedly")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	shutdown := make(chan struct{})
	go func() {
		<-quit
		log.Println("Received termination signal. Shutting down...")
		close(shutdown) // Trigger shutdown for other services
	}()

	select {
	case <-shutdown:
		// Handle shutdown logic (e.g., close DB, stop workers)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Gracefully shutdown the server with a timeout
		if err := app.ShutdownWithContext(ctx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}
		log.Println("Server shutdown complete.")
	}
}
