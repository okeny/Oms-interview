package cmd

import (
	"building_management/config"
	"building_management/database"
	"building_management/router"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
		log.Fatalf("Failed to load config: %s", err)
	}

	port, err := config.Get(Port)
	if err != nil {
		log.Fatalf("Port not found: %s", err)
	}

	// Initialize DB and inject into services if needed
	db, err := database.NewClient()
	if err != nil {
		log.Fatalf("error creating database client: %s", err)
	}
	defer db.Close()

	app, err := router.Init(db)
	if err != nil {
		log.Fatalf("err initializing router: %s", err)
	}
	// Start the server
	host := fmt.Sprintf(":%s", port)
	go func() {
		if err := app.Listen(host); err != nil {
			log.Fatalf("Http server closed unexpectedly, error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Gracefully shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

}
