package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenValidParams_WhenCallNewOrder_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10, 1)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	assert.Equal(t, 0.0, order.FinalPrice) // final price is calculated by the CalculateFinalPrice method
}

func TestGivenValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10, 1)
	assert.NoError(t, err)
	assert.NoError(t, order.CalculateFinalPrice())
	assert.Equal(t, 11.0, order.FinalPrice)
}
