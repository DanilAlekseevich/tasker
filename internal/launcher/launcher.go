package launcher

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
	"tasker/internal/bootstrap"
	"tasker/internal/container"
	"tasker/internal/server"
	"time"
)

type Launcher struct {
	bootstrap *bootstrap.Bootstrap
	container *container.Container
	server    *fiber.App
}

func New() *Launcher {
	return &Launcher{
		bootstrap: bootstrap.New(),
	}
}

func (l *Launcher) Initialize(configPath string) error {

	if err := l.bootstrap.LoadConfig(configPath); err != nil {
		return fmt.Errorf("config initialization failed: %w", err)
	}

	if err := l.bootstrap.InitLogger(); err != nil {
		return fmt.Errorf("logger initialization failed: %w", err)
	}

	if err := l.bootstrap.InitDB(); err != nil {
		return fmt.Errorf("database initialization failed: %w", err)
	}

	l.container = container.New(l.bootstrap.Config, l.bootstrap.Logger, l.bootstrap.DB)
	if err := l.container.InitAll(); err != nil {
		return fmt.Errorf("container initialization failed: %w", err)
	}

	serverFactory := server.New(l.bootstrap.Config, l.bootstrap.Logger)
	l.server = serverFactory.CreateServer(l.container.Controllers)

	return nil
}

func (l *Launcher) Run(ctx context.Context) error {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%d", 8080)
		l.container.Logger.Info("starting server", "address", addr)
		if err := l.server.Listen(addr); err != nil {
			l.container.Logger.Error("server error", "error", err)
		}
	}()

	<-quit
	l.container.Logger.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := l.server.ShutdownWithContext(shutdownCtx); err != nil {
		l.container.Logger.Error("server shutdown error", "error", err)
		return err
	}

	if err := l.container.DB.Close(); err != nil {
		l.container.Logger.Error("database close error", "error", err)
		return err
	}

	l.container.Logger.Info("server stopped gracefully")
	return nil
}
