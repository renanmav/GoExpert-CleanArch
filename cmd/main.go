package main

import (
	"github.com/renanmav/GoExpert-CleanArch/config"
	"github.com/renanmav/GoExpert-CleanArch/internal/events"
	"github.com/renanmav/GoExpert-CleanArch/internal/events/handlers"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/graphql"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/grpc"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/grpc/proto"
	"github.com/renanmav/GoExpert-CleanArch/internal/infra/webserver"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

func main() {
	cfg := config.LoadConfig("../")
	defer cfg.Close()

	orderCreatedEvent := events.NewOrderCreated()
	orderCreatedHandler := handlers.NewOrderCreatedHandler(
		cfg.RabbitMQChannel,
		cfg.RabbitMQExchange,
		cfg.RabbitMQRoutingKey,
	)
	cfg.EventDispatcher.Register("OrderCreated", orderCreatedHandler)

	orderRepository := repository.NewOrderRepository(cfg.DB)
	createOrderUseCase := usecase.NewCreateOrderUseCase(
		orderRepository,
		cfg.EventDispatcher,
		orderCreatedEvent,
	)
	readAllOrdersUseCase := usecase.NewReadAllOrdersUseCase(orderRepository)

	webServer := webserver.NewWebServer(cfg.WebServerPort)
	orderWebServerHandler := webserver.NewOrderWebServerHandler(*createOrderUseCase, *readAllOrdersUseCase)
	webServer.AddHandler("/order", orderWebServerHandler.HandleCreateOrder)
	webServer.AddHandler("/orders", orderWebServerHandler.HandleReadAllOrders)
	go webServer.Start()

	grpcServer := grpc.NewGrpcServer(cfg.GrpcServerPort)
	orderGrpcService := grpc.NewOrderService(*createOrderUseCase, *readAllOrdersUseCase)
	proto.RegisterOrderServiceServer(grpcServer.Server, orderGrpcService)
	go grpcServer.Start()

	graphqlServer := graphql.NewGraphQLServer(cfg.GraphqlServerPort, *createOrderUseCase, *readAllOrdersUseCase)
	graphqlServer.Start()
}
