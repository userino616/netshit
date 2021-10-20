package main

import (
	"netflix-movies/internal/config"
	"netflix-movies/internal/controller"
	"netflix-movies/internal/repository"
	"netflix-movies/internal/server"
	"netflix-movies/internal/services"
	"netflix-movies/pkg/logger"
	"netflix-movies/pkg/postgres"

	"github.com/userino616/netflix-grpc/movieservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	grpcConn, err := grpc.Dial(cfg.Server.Addr, grpc.WithInsecure())
	defer grpcConn.Close()
	if err != nil {
		l.Fatal(err)
	}

	r := repository.New(postgres.GetDB())
	s := services.New(r)
	ctrl := controller.New(s)
	srv := server.New(cfg.Server.Addr)

	movieservice.RegisterMovieServiceServer(srv, ctrl.Movie)
	reflection.Register(srv)

	if err := srv.Start(); err != nil {
		l.Error(err)
	}
}
