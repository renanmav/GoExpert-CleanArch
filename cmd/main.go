package main

import (
	"database/sql"

	"github.com/renanmav/GoExpert-CleanArch/config"
	"github.com/renanmav/GoExpert-CleanArch/internal/delivery/grpc"
	"github.com/renanmav/GoExpert-CleanArch/internal/delivery/grpc/proto"
	"github.com/renanmav/GoExpert-CleanArch/internal/delivery/webserver"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(cfg.DBDriver, cfg.DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := repository.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)

	webServer := webserver.NewWebServer(cfg.WebServerPort)
	orderWebServerHandler := webserver.NewOrderWebServerHandler(*createOrderUseCase)
	webServer.AddHandler("/order", orderWebServerHandler.HandleCreateOrder)
	go webServer.Start()

	grpcServer := grpc.NewGrpcServer(cfg.GrpcServerPort)
	orderGrpcService := grpc.NewOrderService(*createOrderUseCase)
	proto.RegisterOrderServiceServer(grpcServer.Server, orderGrpcService)
	grpcServer.Start()
}
