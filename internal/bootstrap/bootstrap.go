package bootstrap

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"tasker/internal/config"
	"time"
)

type Bootstrap struct {
	Config *config.Config
	Logger *slog.Logger
	DB     *sql.DB
}

func New() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) LoadConfig(configPath string) error {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	b.Config = cfg
	return nil
}

func (b *Bootstrap) InitLogger() error {
	b.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return nil
}

func (b *Bootstrap) InitDB() error {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		b.Config.DB.Host, b.Config.DB.Port, b.Config.DB.User, b.Config.DB.Password, b.Config.DB.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	b.DB = db
	return nil
}
