package gateway

import "github.com/deividroger/ms-wallet/src/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
