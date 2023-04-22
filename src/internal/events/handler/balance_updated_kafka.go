package handler

import (
	"fmt"
	"sync"

	"github.com/deividroger/ms-wallet/src/pkg/events"
	"github.com/deividroger/ms-wallet/src/pkg/kafka"
)

type UpdatedBalanceKafkaHandler struct {
	kafka *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *UpdatedBalanceKafkaHandler {
	return &UpdatedBalanceKafkaHandler{
		kafka: kafka,
	}
}

func (h *UpdatedBalanceKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.kafka.Publish(message, nil, "balances")
	fmt.Println("UpdatedBalanceKafkaHandler", message.GetPayLoad())
}
