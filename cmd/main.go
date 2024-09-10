package main

import (
	"database/sql"

	"github.com/renanmav/GoExpert-CleanArch/config"
	"github.com/renanmav/GoExpert-CleanArch/internal/delivery/webserver"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"

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

	webServer := webserver.NewWebServer(cfg.WebServerPort)
	orderWebServerHandler := webserver.NewOrderWebServerHandler(orderRepository)
	webServer.AddHandler("/order", orderWebServerHandler.HandleCreateOrder)
	webServer.Start()
}
