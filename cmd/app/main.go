package main

import (
	"context"
	"errors"
	"github.com/chincharovpc/goarch/internal/config"
	"github.com/chincharovpc/goarch/internal/repository"
	"github.com/chincharovpc/goarch/internal/service"
	"github.com/chincharovpc/goarch/internal/transport/rest"
	restHandler "github.com/chincharovpc/goarch/internal/transport/rest/handler"
	"github.com/chincharovpc/goarch/pkg/postgres"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	postgresConfigs := config.GetPostgresConfig()
	db, err := postgres.Connect(postgres.ConnectInput{
		Host:     postgresConfigs.Host,
		Port:     postgresConfigs.Port,
		User:     postgresConfigs.User,
		Password: postgresConfigs.Password,
		Database: postgresConfigs.Database,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories := repository.NewRepositories(db, logger)
	services := service.NewServices(repositories)

	// Rest Server
	restHandlers := restHandler.NewHandlers(services)
	restServer := rest.NewServer(config.GetHttpConfig(), restHandlers.Init())

	go func() {
		log.Println("Starting Rest Server")
		if err := restServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("%s%s", "rest ListenAndServe err: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := restServer.Stop(ctx); err != nil {
		logrus.Errorf("failed to stop server: %v", err)
	}

	log.Print("Shutting down server")
}
