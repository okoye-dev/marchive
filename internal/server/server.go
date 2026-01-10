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

func NewServer() *Server {
	return &Server{
		cfg: config.DefaultConfig(),
		Router: api.NewRouter(),
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on port %s\n", s.cfg.Port)
	return http.ListenAndServe(":"+s.cfg.Port, s.Router)
}