package bootstrap

import (
	"context"
	"fmt"
	"github/idbeholdv18/expense-tracker/internal/config"
	"github/idbeholdv18/expense-tracker/internal/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{config: cfg}
}

func (a *App) Run() {
	db := database.New(a.config.DatabaseURL)
	defer db.Close()

	repositories := RegisterRepositories(db)
	services := RegisterServices(a.config, repositories)
	handlers := RegisterHandlers(services)
	middlewares := RegisterMiddlewares(a.config, services)

	RegisterErrors()

	mux := http.NewServeMux()
	RegisterRoutes(mux, middlewares, handlers)

	server := registerServer(a.config, mux)

	log.Println("server is running...")
	go func() {
		err := server.httpServer.ListenAndServeTLS("", "")
		if err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	fmt.Fprint(os.Stdout, "shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.httpServer.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error during shutting down server: %v\n", err)
	}
}
