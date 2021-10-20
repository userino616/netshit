package main

import (
	"context"
	"netflix-auth/internal/config"
	"netflix-auth/internal/handlers"
	"netflix-auth/internal/repository"
	"netflix-auth/internal/server"
	"netflix-auth/internal/services"
	"netflix-auth/pkg/logger"
	"netflix-auth/pkg/postgres"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	logger.Init()
	l := logger.GetLogger()
	defer logger.Close()
	l.Info("logger initialized")

	cfg := config.GetConfig()
	l.Info("config initialized")
	l.Debugf("config data: %v", cfg)

	postgres.Load(cfg)
	db := postgres.GetDB()
	defer db.Close()
	l.Info("db initialized")

	grpcConn, err := grpc.Dial(cfg.Server.GRPCAddr, grpc.WithInsecure())
	defer grpcConn.Close()
	if err != nil {
		l.Fatal(err)
	}

	r := repository.New(db)
	s := services.New(r, grpcConn, cfg)
	h := handlers.New(s)

	srv := server.New(cfg, h.InitRoutes())

	go func() {
		if err := srv.Start(); err != nil {
			l.Fatal(err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown

	l.Info("shutting down server")
	if err := srv.Shutdown(context.Background()); err != nil {
		l.Error(err)
	}
}
