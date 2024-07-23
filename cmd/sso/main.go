package main

import (
	"github.com/adametsderschopfer/gRPC-auth-service/internal/config"
	"github.com/adametsderschopfer/gRPC-auth-service/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("Starting gRPC Auth Service", slog.String("env", cfg.Env))
	log.Debug("Debug messages are: enabled")

	/*
		TODO:
			- Инициализировать приложение (app)
			- Запустить grpc приложение
	*/

	log.Error("Server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = setupPrettySlog()
	case config.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
