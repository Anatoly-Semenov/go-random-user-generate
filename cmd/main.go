package main

import (
	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"user-service/internal/config"
	"user-service/pkg/logger/handlers/slogpretty"
)

func main() {
	// Init config
	cfg := config.GetConfig()

	// Init logger
	log := setupLogger(cfg.Env)

	// init storage | gorm TODO

	// Init router
	router := httprouter.New()

	start(cfg, log, router)
}

func start(cfg *config.Config, log *slog.Logger, router *httprouter.Router) {
	// Run server
	log.Info("Starting server", slog.String("Address", cfg.Address))

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HttpsSerer.Timeout,
		WriteTimeout: cfg.HttpsSerer.Timeout,
		IdleTimeout:  cfg.HttpsSerer.Iddle_timeout,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("Failed to start server")
	}

	log.Error("Server stopped")
}

func setupLogger(env config.Env) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.LOCAL:
		log = setupPrettySlog()
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
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
