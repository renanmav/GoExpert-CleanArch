package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	Server *grpc.Server
	Port   string
}

func NewGrpcServer(grpcPort string) *GrpcServer {
	return &GrpcServer{
		Server: grpc.NewServer(),
		Port:   grpcPort,
	}
}

func (g *GrpcServer) RegisterService(service grpc.ServiceDesc, impl interface{}) {
	g.Server.RegisterService(&service, impl)
}

func (g *GrpcServer) Start() {
	fmt.Println("Starting gRPC server on port:", g.Port)
	reflection.Register(g.Server)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", g.Port))
	if err != nil {
		panic(err)
	}
	g.Server.Serve(listener)
}
