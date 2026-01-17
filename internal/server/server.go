package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/okoye-dev/marchive/internal/api"
	"github.com/okoye-dev/marchive/internal/config"
)

type Server struct {
	cfg    *config.Config
	Router *chi.Mux
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		Router: api.NewRouter(),
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on port %s\n", s.cfg.Server.Port)
	return http.ListenAndServe(":"+s.cfg.Server.Port, s.Router)
}