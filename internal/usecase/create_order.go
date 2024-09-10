package usecase

import (
	"github.com/renanmav/GoExpert-CleanArch/internal/entity"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
)

type CreateOrderInput struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewCreateOrderUseCase(orderRepository repository.OrderRepositoryInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{OrderRepository: orderRepository}
}

func (c *CreateOrderUseCase) Execute(input CreateOrderInput) (*CreateOrderOutput, error) {
	order, err := entity.NewOrder(input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	output := &CreateOrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	return output, nil
}
