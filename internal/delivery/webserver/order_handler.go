package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/renanmav/GoExpert-CleanArch/internal/repository"
	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

type OrderWebServerHandler struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewOrderWebServerHandler(orderRepository repository.OrderRepositoryInterface) *OrderWebServerHandler {
	return &OrderWebServerHandler{
		OrderRepository: orderRepository,
	}
}

func (o *OrderWebServerHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateOrderInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrderUseCase := usecase.NewCreateOrderUseCase(o.OrderRepository)

	output, err := createOrderUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
