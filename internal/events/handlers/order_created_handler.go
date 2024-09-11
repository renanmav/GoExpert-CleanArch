package handlers

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/renanmav/GoExpert-Events/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel    *amqp.Channel
	RabbitMQExchange   string
	RabbitMQRoutingKey string
}

func NewOrderCreatedHandler(
	rabbitMQChannel *amqp.Channel,
	rabbitMQExchange string,
	rabbitMQRoutingKey string,
) events.EventHandlerInterface {
	return &OrderCreatedHandler{
		RabbitMQChannel:    rabbitMQChannel,
		RabbitMQExchange:   rabbitMQExchange,
		RabbitMQRoutingKey: rabbitMQRoutingKey,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	jsonOutput, err := json.Marshal(event.GetPayload())
	if err != nil {
		panic(err)
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		h.RabbitMQExchange,   // exchange
		h.RabbitMQRoutingKey, // routing key
		false,                // mandatory
		false,                // immediate
		msg,                  // message to publish
	)

	fmt.Printf("Order created: %v\n", event.GetPayload())
}
