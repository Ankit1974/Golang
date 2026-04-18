package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/imp012-deep/buildingApi/internal/config"
	"github.com/imp012-deep/buildingApi/internal/repository"
	"github.com/imp012-deep/buildingApi/internal/routes"
)

// Entry point of your application

/*
   Why cmd/?

   Allows multiple binaries in future

   Example:
   cmd/api        → REST API
   cmd/worker     → background jobs
   cmd/migrate    → DB migrations
*/

/*
  you can treat this main.go as a template. You can copy-paste this file to start a new API, but you would just
  change the imports and the specific repository/handler initialization lines.
*/

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("Starting application",
		slog.String("name", cfg.App.Name),
		slog.String("version", cfg.App.Version),
		slog.String("environment", cfg.App.Environment),
		slog.String("port", cfg.Server.Port),
	)

	// Initialize repository
	repo := repository.NewInMemoryCourseRepository()

	// Seed with sample data
	if err := repo.Seed(); err != nil {
		logger.Error("Failed to seed data", slog.Any("error", err))
		os.Exit(1)
	}
	logger.Info("Repository initialized", slog.Int("courses", repo.Count()))

	// Setup routes with middleware
	router := routes.SetupRouter(repo, logger)

	// Configure HTTP server with timeouts
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server listening", slog.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal
	<-quit
	logger.Info("Shutdown signal received, starting graceful shutdown...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", slog.Any("error", err))
		os.Exit(1)
	}

	logger.Info("Server exited gracefully")
	fmt.Println("✅ Server shutdown complete")
}
