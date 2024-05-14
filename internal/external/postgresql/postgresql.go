package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // postgres driver
)

func New(config *config.Config) *sql.DB {
	connectionString := config.DBConfig.ConnectionString

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
	}

	err = runMigrations(db)
	if err != nil {
		log.Println(err)
	}

	return db
}

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("ERROR creating driver instance: %w", err)
	}

	migrationsPath := "migrations"

	fmt.Println("db: creating migrations")

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("ERROR creating migrate instance: %w", err)
	}

	if err = m.Up(); err != nil {
		fmt.Printf("ERROR running up migrations: %v\n", err)
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}

	return nil
}
