package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"user-service/internal/config"
	"user-service/internal/user"
	"user-service/pkg/logger/handlers/slogpretty"
)

func main() {
	// Init config
	cfg := config.GetConfig()

	// Init logger
	log := setupLogger(cfg.Env)

	// init storage | gorm
	db := connectionDB(cfg, log)

	// Init router
	router := httprouter.New()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userController := user.NewController(userService)
	userController.Register(router)

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

func connectionDB(cfg *config.Config, log *slog.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", cfg.Database.Host, cfg.Database.User, cfg.Database.Db, cfg.Database.Password, cfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		message := "Failed to connect database"

		log.Error(message)
		panic(message)
	} else {
		log.Info("Successful connect to database")
	}

	db.AutoMigrate(&user.Model{})

	setDefaultDataToDb(db)

	return db
}

func setDefaultDataToDb(db *gorm.DB) {
	fmt.Println("Users: ", user.Users)

	for id := range user.Users {
		db.Create(user.Users[id])
	}
}
