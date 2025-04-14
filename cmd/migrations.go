package cmd

import (
	"building_management/database"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	// Import PostgreSQL driver for database/sql
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrationDir string

func Migrations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrations",
		Short: "start BMS in migrations mode",
	}
	cmd.AddCommand(
		runUpCmd(),
		runDownCmd(),
		runCreateCmd(),
		runStatusCmd(),
	)
	return cmd
}

func runUpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Apply all up migrations",
		Run: func(_ *cobra.Command, _ []string) {

			dbClient, err := database.NewClient()
			if err != nil {
				log.Fatalln(err)
			}
			defer dbClient.Close()

			n, err := migrate.Exec(dbClient, "postgres", &migrate.FileMigrationSource{
				Dir: "migrations",
			}, migrate.Up)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Applied %d migration(s)!\n", n)

		},
	}
}

func runDownCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "run migrations down",
		Run: func(_ *cobra.Command, _ []string) {
			// Connect to the database
			dbClient, err := database.NewClient()
			if err != nil {
				log.Fatalln(err)
			}
			defer dbClient.Close()

			n, err := migrate.ExecMax(dbClient, "postgres", &migrate.FileMigrationSource{
				Dir: "migrations",
			}, migrate.Down, 1)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Rolled back %d migration(s)\n", n)
		},
	}
}

func runCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new SQL migration file",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			name := args[0]
			timestamp := time.Now().Format("20060102150405")
			filename := fmt.Sprintf("%s_%s.sql", timestamp, name)
			path := filepath.Join(migrationDir, filename)

			content := `-- +migrate Up

-- +migrate Down
`
			if err := os.MkdirAll(migrationDir, 0o755); err != nil {
				return fmt.Errorf("failed to create migration directory: %w", err)
			}

			if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
				return fmt.Errorf("failed to write migration file: %w", err)
			}

			fmt.Println("✅ Created migration:", path)
			return nil
		},
	}

	cmd.Flags().StringVarP(&migrationDir, "dir", "d", "migrations", "Path to migrations directory")
	return cmd
}

func runStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "create a new migration using migrate CLI",
		Run: func(_ *cobra.Command, args []string) {
			dbClient, err := database.NewClient()
			if err != nil {
				log.Fatalln(err)
			}
			defer dbClient.Close()

			records, err := migrate.GetMigrationRecords(dbClient, "postgres")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Applied migrations (%d):\n", len(records))
			for _, m := range records {
				fmt.Println("-", m.Id)
			}
		},
	}
}
