package usecase

import (
	"github.com/renanmav/GoExpert-CleanArch/internal/entity"
	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
	"github.com/renanmav/GoExpert-Events/pkg/events"
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
	EventDispatcher events.EventDispatcherInterface
	OrderCreated    events.EventInterface
}

func NewCreateOrderUseCase(
	orderRepository repository.OrderRepositoryInterface,
	eventDispatcher events.EventDispatcherInterface,
	orderCreated events.EventInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
		EventDispatcher: eventDispatcher,
		OrderCreated:    orderCreated,
	}
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

	c.OrderCreated.SetPayload(output)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return output, nil
}
