package db

import (
	"github.com/Kudzeri/go-shoe-store-api/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
)

func AutoMigrate() {
	db, err := config.DB.DB()
	if err != nil {
		log.Fatalf("Could not get DB: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
