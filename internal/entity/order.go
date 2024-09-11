package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidID    = errors.New("invalid id")
	ErrInvalidPrice = errors.New("invalid price")
	ErrInvalidTax   = errors.New("invalid tax")
)

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func NewOrder(price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    uuid.New().String(),
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return ErrInvalidID
	}
	if o.Price <= 0 {
		return ErrInvalidPrice
	}
	if o.Tax <= 0 {
		return ErrInvalidTax
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
