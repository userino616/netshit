package server

import (
	"net"
	"netflix-movies/pkg/logger"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
	*grpc.Server
}

func New(addr string) *GRPCServer {
	return &GRPCServer{
		addr:   addr,
		Server: grpc.NewServer(),
	}
}

func (s *GRPCServer) Start() error {
	l := logger.GetLogger()
	l.Info("starting grpc server")

	con, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}

	return s.Serve(con)
}
