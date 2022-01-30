package server

import (
	"context"
	"errors"
	"github.com/chincharovpc/goarch/config"
	"github.com/chincharovpc/goarch/controller"
	"github.com/chincharovpc/goarch/repository"
	"github.com/chincharovpc/goarch/router"
	"github.com/chincharovpc/goarch/service"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.HttpConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes << 20,
		},
	}
}

func Run() {
	db, err := config.GetConnection()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	httpConfig := config.GetHttpConfig()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories := repository.NewRepositories(db, logger)
	services := service.NewServices(repositories)
	controllers := controller.NewControllers(services)
	routers := router.NewRouters(controllers)

	// HTTP Server
	httpServer := NewServer(httpConfig, routers.Init())

	go func() {
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("%s%s", "произошла ошибка при запуске http сервера: %s\n", err.Error())
		}
	}()

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	defer func(server *Server, ctx context.Context) {
		err := server.Stop(ctx)
		if err != nil {
			log.Fatalf("произошла ошибка при остановке сервера: %s\n", err.Error())
		}
	}(httpServer, ctx)

	log.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
