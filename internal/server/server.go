package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/okoye-dev/marchive/internal/api"
	"github.com/okoye-dev/marchive/internal/config"
	"github.com/okoye-dev/marchive/internal/files"
	"github.com/okoye-dev/marchive/internal/storage"
)

type Server struct {
	cfg    *config.Config
	Router *chi.Mux
}

func NewServer(cfg *config.Config) *Server {
	localStorage := storage.NewLocalClient(cfg.Storage.Root)
	fileService := files.NewFileService(localStorage)
	handlers := &api.Handlers{
		Files:         fileService,
		DefaultBucket: cfg.Storage.DefaultBucket,
	}

	return &Server{
		cfg:    cfg,
		Router: api.NewRouter(handlers),
	}
}

func (s *Server) Start() error {
	srv := &http.Server{
		Addr:         ":" + s.cfg.Server.Port,
		Handler:      s.Router,
		ReadTimeout:  time.Duration(s.cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.cfg.Server.IdleTimeout) * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("Starting server on port %s\n", s.cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(stop)

	select {
	case err := <-errCh:
		return err
	case <-stop:
		log.Println("Shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.Server.ShutdownTimeout)*time.Second)
		defer cancel()
		return srv.Shutdown(ctx)
	}
}
