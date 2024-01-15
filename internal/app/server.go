package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fairusatoir/golang-gin-rest/cmd/api/config"
	"github.com/fairusatoir/golang-gin-rest/internal/controllers/handlers"
	"github.com/fairusatoir/golang-gin-rest/internal/controllers/routers"
	"github.com/fairusatoir/golang-gin-rest/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*http.Server
}

func NewServer(cfg *config.Config, ph *handlers.ProductHandler) *Server {
	router := setupApiServer(cfg)
	api := router.Group("/api")
	routers.NewProductRouter(ph, api).AddRoutes()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{server}
}

func setupApiServer(cfg *config.Config) *gin.Engine {
	var mode = gin.ReleaseMode
	if cfg.Debug {
		mode = gin.DebugMode
	}

	gin.ForceConsoleColor()
	gin.SetMode(mode)

	router := gin.New()
	router.Use(gin.Recovery())

	return router
}

func (s *Server) Run() (err error) {
	// Gracefull Shutdown
	go func() {
		log.InfoF("success to listen and serve on :%d", logrus.Fields{log.Category: log.Server}, 8080)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.FatalF("Failed to listen and serve:%+v", logrus.Fields{log.Category: log.Server}, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	log.Info("shutdown server ...", logrus.Fields{log.Category: log.Server})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Info("timeout of 5 seconds.", logrus.Fields{log.Category: log.Server})
	log.Info("server exiting", logrus.Fields{log.Category: log.Server})
	return
}
