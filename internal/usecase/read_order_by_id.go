package usecase

import "github.com/renanmav/GoExpert-CleanArch/internal/repository"

type ReadOrderByIdInput struct {
	ID string
}

type ReadOrderByIdOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ReadOrderByIdUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewReadOrderByIdUseCase(orderRepository repository.OrderRepositoryInterface) *ReadOrderByIdUseCase {
	return &ReadOrderByIdUseCase{OrderRepository: orderRepository}
}

func (u *ReadOrderByIdUseCase) Execute(input ReadOrderByIdInput) (*ReadOrderByIdOutput, error) {
	order, err := u.OrderRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}

	output := &ReadOrderByIdOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	return output, nil
}
