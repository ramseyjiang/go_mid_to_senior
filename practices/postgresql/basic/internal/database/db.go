package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	dbHost         string
	dbPort         string
	dbUser         string
	dbPassword     string
	dbName         string
	migrationsPath string
)

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting current directory")
	}
	parentDir := filepath.Dir(currentDir)
	envPath := filepath.Join(parentDir, ".env")
	if err = godotenv.Load(envPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to load .env file")
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	migrationsPath = os.Getenv("APP_MIGRATIONS_PATH")
}

func NewPostgresDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Info().Msg("Successfully connected to PostgreSQL!")
	return db, nil
}

func RunMigrations(db *sql.DB) error {
	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName),
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	log.Info().Msg("Migrations applied successfully!")
	return nil
}
