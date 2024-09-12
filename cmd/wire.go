//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
	"github.com/renanmav/GoExpert-Events/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	repository.NewOrderRepository,
	wire.Bind(new(repository.OrderRepositoryInterface), new(*repository.OrderRepository)),
)

func NewCreateOrderUseCase(
	db *sql.DB,
	eventDispatcher events.EventDispatcherInterface,
	event events.EventInterface,
) *usecase.CreateOrderUseCase {
	wire.Build(
		usecase.NewCreateOrderUseCase,
		repository.NewOrderRepository,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewReadAllOrdersUseCase(
	db *sql.DB,
) *usecase.ReadAllOrdersUseCase {
	wire.Build(
		usecase.NewReadAllOrdersUseCase,
		repository.NewOrderRepository,
	)
	return &usecase.ReadAllOrdersUseCase{}
}

func NewReadOrderByIdUseCase(
	db *sql.DB,
) *usecase.ReadOrderByIdUseCase {
	wire.Build(
		usecase.NewReadOrderByIdUseCase,
		repository.NewOrderRepository,
	)
	return &usecase.ReadOrderByIdUseCase{}
}
