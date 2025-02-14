package main

import (
	"log/slog"
	"os"

	"github.com/NarthurN/url-shortener/internal/config"
	"github.com/NarthurN/url-shortener/internal/lib/logger/sl"
	"github.com/NarthurN/url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
	// config / cleanenv
	cfg := config.MustLoad()

	// loger / slog
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// storage / sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	id, err := storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save URL", sl.Err(err))
		os.Exit(1)
	}
	
	log.Info("save URL", slog.Int64("id", id))

	id, err = storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save URL", sl.Err(err))
		os.Exit(1)
	}
	_ = id
	_ = storage
	// router / chi, chi_render

	// server / go
}

func setupLogger(env string) *slog.Logger{
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}