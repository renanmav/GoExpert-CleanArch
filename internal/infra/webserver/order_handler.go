package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/renanmav/GoExpert-CleanArch/internal/usecase"
)

type OrderWebServerHandler struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderWebServerHandler(createOrderUseCase usecase.CreateOrderUseCase) *OrderWebServerHandler {
	return &OrderWebServerHandler{
		CreateOrderUseCase: createOrderUseCase,
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
