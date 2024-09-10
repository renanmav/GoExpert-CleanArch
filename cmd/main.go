package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := repository.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository)

	input := usecase.CreateOrderInput{
		Price: 100,
		Tax:   10,
	}

	output, err := createOrderUseCase.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
