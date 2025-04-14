package database

import (
	"building_management/config"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewClient() (*sql.DB, error) {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	// Build the DSN dynamically
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)

	// Connect to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Set connection pool settings to avoid connection timeouts
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to database!")
	return db, nil
}
