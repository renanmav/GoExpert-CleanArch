package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

type OrderWebServerHandler struct {
	CreateOrderUseCase   usecase.CreateOrderUseCase
	ReadAllOrdersUseCase usecase.ReadAllOrdersUseCase
	ReadOrderByIdUseCase usecase.ReadOrderByIdUseCase
}

func NewOrderWebServerHandler(
	createOrderUseCase usecase.CreateOrderUseCase,
	readAllOrdersUseCase usecase.ReadAllOrdersUseCase,
	readOrderByIdUseCase usecase.ReadOrderByIdUseCase,
) *OrderWebServerHandler {
	return &OrderWebServerHandler{
		CreateOrderUseCase:   createOrderUseCase,
		ReadAllOrdersUseCase: readAllOrdersUseCase,
		ReadOrderByIdUseCase: readOrderByIdUseCase,
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

func (o *OrderWebServerHandler) HandleReadOrderById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	input := usecase.ReadOrderByIdInput{
		ID: id,
	}

	output, err := o.ReadOrderByIdUseCase.Execute(input)
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
