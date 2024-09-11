package grpc

import (
	"context"

	"github.com/renanmav/GoExpert-CleanArch/internal/infra/grpc/proto"
	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

type OrderService struct {
	proto.UnimplementedOrderServiceServer
	CreateOrderUseCase   usecase.CreateOrderUseCase
	ReadAllOrdersUseCase usecase.ReadAllOrdersUseCase
	ReadOrderByIdUseCase usecase.ReadOrderByIdUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	readAllOrdersUseCase usecase.ReadAllOrdersUseCase,
	readOrderByIdUseCase usecase.ReadOrderByIdUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase:   createOrderUseCase,
		ReadAllOrdersUseCase: readAllOrdersUseCase,
		ReadOrderByIdUseCase: readOrderByIdUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	input := usecase.CreateOrderInput{
		Price: float64(req.Price),
		Tax:   float64(req.Tax),
	}

	output, err := s.CreateOrderUseCase.Execute(input)
	if err != nil {
		return nil, err
	}

	return &proto.CreateOrderResponse{
		Order: &proto.Order{
			Id:         output.ID,
			Price:      float32(output.Price),
			Tax:        float32(output.Tax),
			FinalPrice: float32(output.FinalPrice),
		},
	}, nil
}

func (s *OrderService) RealAllOrders(ctx context.Context, req *proto.RealAllOrdersRequest) (*proto.RealAllOrdersResponse, error) {
	output, err := s.ReadAllOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := make([]*proto.Order, len(output.Orders))
	for i, order := range output.Orders {
		orders[i] = &proto.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
	}

	return &proto.RealAllOrdersResponse{
		Orders: orders,
	}, nil
}

func (s *OrderService) ReadOrderById(ctx context.Context, req *proto.ReadOrderByIdRequest) (*proto.ReadOrderByIdResponse, error) {
	input := usecase.ReadOrderByIdInput{
		ID: req.Id,
	}

	output, err := s.ReadOrderByIdUseCase.Execute(input)
	if err != nil {
		return nil, err
	}

	return &proto.ReadOrderByIdResponse{
		Order: &proto.Order{
			Id:         output.ID,
			Price:      float32(output.Price),
			Tax:        float32(output.Tax),
			FinalPrice: float32(output.FinalPrice),
		},
	}, nil
}
