package events

import (
	"time"

	"github.com/renanmav/GoExpert-Events/pkg/events"
)

type OrderCreated struct {
	Name    events.EventName
	Payload interface{}
}

func NewOrderCreated() events.EventInterface {
	return &OrderCreated{
		Name: "OrderCreated",
	}
}

func (e *OrderCreated) GetName() events.EventName {
	return e.Name
}

func (e *OrderCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderCreated) GetDateTime() time.Time {
	return time.Now()
}
