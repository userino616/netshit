package main

import (
	"context"
	"netflix-auth/pkg/redis"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"netflix-auth/internal/config"
	"netflix-auth/internal/handlers"
	"netflix-auth/internal/repository"
	"netflix-auth/internal/server"
	"netflix-auth/pkg/logger"
	"netflix-auth/pkg/postgres"
)

func main() {
	// норм ли оставлять чтение env файла, потому что на локалке его нужно читать, а в докере нет ?
	godotenv.Load()
	cfg := config.GetConfig()

	logger.Init(cfg.LogLvl)
	l := logger.GetLogger()
	l.Info("logger initialized")

	l.Debugf("config data: %v", cfg)

	postgres.Load(cfg)
	defer postgres.Close()
	l.Info("db initialized")

	redis.Load(cfg)
	defer redis.Close()
	l.Info("redis initialized")

	grpcConn, err := grpc.Dial(cfg.Server.GRPCAddr, grpc.WithInsecure())
	defer grpcConn.Close()
	if err != nil {
		l.Fatal(err)
	}

	db := postgres.GetDB()
	redisDB := redis.GetDB()
	r := repository.New(db, redisDB)
	h := handlers.New(r, grpcConn, cfg)

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
