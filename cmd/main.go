package main

import (
	"github.com/renanmav/GoExpert-CleanArch/config"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/graphql"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/grpc"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/grpc/proto"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/webserver"
)

func main() {
	cfg := config.LoadConfig("../")
	defer cfg.Close()

	// Wired dependencies
	createOrderUseCase := NewCreateOrderUseCase(cfg.DB, cfg.EventDispatcher, cfg.OrderCreatedEvent)
	readAllOrdersUseCase := NewReadAllOrdersUseCase(cfg.DB)
	readOrderByIdUseCase := NewReadOrderByIdUseCase(cfg.DB)

	// Web server
	webServer := webserver.NewWebServer(cfg.WebServerPort)
	orderWebServerHandler := webserver.NewOrderWebServerHandler(*createOrderUseCase, *readAllOrdersUseCase, *readOrderByIdUseCase)
	webServer.AddHandler("POST", "/order", orderWebServerHandler.HandleCreateOrder)
	webServer.AddHandler("GET", "/orders", orderWebServerHandler.HandleReadAllOrders)
	webServer.AddHandler("GET", "/order", orderWebServerHandler.HandleReadOrderById)
	go webServer.Start()

	// GRPC server
	grpcServer := grpc.NewGrpcServer(cfg.GrpcServerPort)
	orderGrpcService := grpc.NewOrderService(*createOrderUseCase, *readAllOrdersUseCase, *readOrderByIdUseCase)
	proto.RegisterOrderServiceServer(grpcServer.Server, orderGrpcService)
	go grpcServer.Start()

	// GraphQL server
	graphqlServer := graphql.NewGraphQLServer(cfg.GraphqlServerPort, *createOrderUseCase, *readAllOrdersUseCase, *readOrderByIdUseCase)
	graphqlServer.Start()
}
