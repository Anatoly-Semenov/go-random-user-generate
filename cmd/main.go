package cmd

import (
	"golang.org/x/exp/slog"
	"os"
	"user-service/internal/config"
	"user-service/pkg/logger/handlers/slogpretty"
)

func Main() {
	// Init config
	cfg := config.GetConfig()

	// Init logger
	log := setupLogger(cfg.Env)

	// init storage | gorm TODO

	// Init router TODO
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
