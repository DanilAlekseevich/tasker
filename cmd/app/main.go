package main

import (
	"context"
	"log/slog"
	"os"
	"tasker/internal/launcher"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yaml"
	}

	launcher := launcher.New()

	if err := launcher.Initialize(configPath); err != nil {
		slog.Error("failed to initialize application", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()
	if err := launcher.Run(ctx); err != nil {
		slog.Error("application run error", "error", err)
		os.Exit(1)
	}
}
