package bootstrap

import (
	"github/idbeholdv18/expense-tracker/internal/config"
	"github/idbeholdv18/expense-tracker/internal/database"
	"log"
	"net/http"
)

type App struct {
	config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{config: cfg}
}

func (a *App) Run() {
	db := database.New(a.config.DatabaseURL)

	repositories := registerRepositories(db)

	services := registerServices(a.config, repositories)

	handlers := registerHandlers(services)
	middlewares := registerMiddlewares(a.config, services)

	registerErrors()
	registerRoutes(middlewares, handlers)

	log.Fatal(http.ListenAndServeTLS(":"+a.config.Port, "cert/cert.pem", "cert/key.pem", nil))
}
