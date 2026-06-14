package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"embed"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Database struct {
	DB *gorm.DB
}

type DatabaseConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
}

func CreateConnection(config *DatabaseConfig) *sql.DB {
	db, err := sql.Open("pgx", config.DSN)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)

	return db
}

func NewDatabase(cfg *DatabaseConfig) (*Database, func(), error) {
	sqlDB := CreateConnection(cfg)
	applyMigrations(sqlDB)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")

	cleanup := func() {
		log.Println("Closing database connection...")
		sqlDB.Close()
	}

	return &Database{DB: db}, cleanup, nil
}

func applyMigrations(db *sql.DB) {
	// Настройка Goose для работы с embed файлами
	goose.SetBaseFS(embedMigrations)

	// Применение миграций
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
