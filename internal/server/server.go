package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"smart-city/config"
	logger "smart-city/pkg/logger"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	db     *gorm.DB
	logger logger.Logger
}

func NewServer(cfg *config.Config, db *gorm.DB, logger logger.Logger) *Server {
	return &Server{cfg: cfg, db: db, logger: logger, echo: echo.New()}
}
func (s *Server) Run() error {
	// create http server
	server := &http.Server{
		Addr:           ":" + s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	//start server
	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()
	// Map handlers
	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	//gracefully shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	shutdownCtx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(shutdownCtx); err != nil {
		s.logger.Errorf("http server shutdown: %w", err)
		return err
	}
	s.logger.Info("Server shutdown")
	return nil
}
