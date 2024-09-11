package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

type OrderWebServerHandler struct {
	CreateOrderUseCase   usecase.CreateOrderUseCase
	ReadAllOrdersUseCase usecase.ReadAllOrdersUseCase
}

func NewOrderWebServerHandler(
	createOrderUseCase usecase.CreateOrderUseCase,
	readAllOrdersUseCase usecase.ReadAllOrdersUseCase,
) *OrderWebServerHandler {
	return &OrderWebServerHandler{
		CreateOrderUseCase:   createOrderUseCase,
		ReadAllOrdersUseCase: readAllOrdersUseCase,
	}
}

func (o *OrderWebServerHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateOrderInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := o.CreateOrderUseCase.Execute(input)
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

func (o *OrderWebServerHandler) HandleReadAllOrders(w http.ResponseWriter, r *http.Request) {
	output, err := o.ReadAllOrdersUseCase.Execute()
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
