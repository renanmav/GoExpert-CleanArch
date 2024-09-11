package usecase

import (
	"github.com/renanmav/GoExpert-CleanArch/internal/entity"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
)

// No input for this use case
type ReadAllOrdersInput struct{}

type ReadAllOrdersOutput struct {
	Orders []*entity.Order `json:"orders"`
}

type ReadAllOrdersUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewReadAllOrdersUseCase(orderRepository repository.OrderRepositoryInterface) *ReadAllOrdersUseCase {
	return &ReadAllOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ReadAllOrdersUseCase) Execute() (*ReadAllOrdersOutput, error) {
	orders, err := u.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return &ReadAllOrdersOutput{Orders: orders}, nil
}
