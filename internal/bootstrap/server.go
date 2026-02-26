package bootstrap

import (
	"crypto/tls"
	"github/idbeholdv18/expense-tracker/internal/config"
	"log"
	"net/http"
	"os"
)

type Server struct {
	httpServer *http.Server
	certFile   string
	keyFile    string
}

func registerServer(cfg *config.Config, handler *http.ServeMux) *Server {
	cer, err := tls.LoadX509KeyPair(cfg.TLSFile, cfg.TLSKey)

	if err != nil {
		log.Fatal("tls config error:", err)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:      cfg.Host + ":" + cfg.Port,
		Handler:   handler,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cer}},
	}

	return &Server{
		httpServer: srv,
	}
}
