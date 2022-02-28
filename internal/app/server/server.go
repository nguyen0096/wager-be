package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"

	"wager-be/internal/service"
	"wager-be/pkg/config"
)

type Server interface {
	Run(ctx context.Context)
}

type server struct {
	logger     logr.Logger
	httpServer *http.Server

	// services
	wagerService service.WagerService
}

func New(
	logger logr.Logger,
	cfg config.ServerConfig,
	wagerService service.WagerService,
) Server {
	s := &server{
		logger:       logger,
		wagerService: wagerService,
	}

	s.httpServer = &http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	// setup routes
	r := gin.Default()

	r.GET("/healthz", s.handleHealthcheck)
	r.GET("/livez", s.handleLivecheck)

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/wagers", s.handleCreateWager)
		v1.GET("/wagers", s.handleListWager)
		v1.POST("/buy/:wager_id", s.handleBuy)
	}
	s.httpServer.Handler = r

	return s
}

func (s *server) Run(ctx context.Context) {
	exitCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)

	signal.Notify(exitCh, os.Interrupt)

	s.logger.Info(fmt.Sprintf("listening on %s", s.httpServer.Addr))
	go func() {
		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	err := awaitTermination(ctx, errCh, exitCh)
	if err != nil {
		s.logger.Error(err, "failed to listen to server")
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		s.logger.Error(err, "failed to shut down server")
	}
}

func awaitTermination(ctx context.Context, errCh chan error, exitCh chan os.Signal) error {
	select {
	case err := <-errCh:
		return err
	case <-exitCh:
		return nil
	case <-ctx.Done():
		return nil
	}
}

func (s *server) handleHealthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}

func (s *server) handleLivecheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
