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

	repositories := RegisterRepositories(db)

	services := RegisterServices(a.config, repositories)

	handlers := RegisterHandlers(services)
	middlewares := RegisterMiddlewares(a.config, services)

	RegisterErrors()
	RegisterRoutes(middlewares, handlers)

	log.Fatal(http.ListenAndServeTLS(":"+a.config.Port, "cert/cert.pem", "cert/key.pem", nil))
}
