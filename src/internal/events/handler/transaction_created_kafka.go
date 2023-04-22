package handler

import (
	"fmt"
	"sync"

	"github.com/deividroger/ms-wallet/src/pkg/events"
	"github.com/deividroger/ms-wallet/src/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.kafka.Publish(message, nil, "transactions")
	fmt.Println("TransactionCreatedKafkaHandler", message.GetPayLoad())
}
